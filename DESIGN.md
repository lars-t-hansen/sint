# Design and implementation notes

## Unwinding, errors, and call/cc

(This is rough)

Currently we use `panic` to signal evaluation errors and there was a plan to use it 
also for `call/cc` unwinding, since our continuations are only one-shot same-thread 
scape functions (for everything else, use goroutines directly).

In addition to our uses, library functions not controlled by us may panic on error.

One possibility is therefore to store a flag on the context to signal `call/cc` and to
use recover to handle panics and `call/cc` alike.  This sort of works.

However, this does not allow the system to truly panic when it needs to, ie, when
some logic error is found in the system itself or a library function panics.  In this
case we don't want unwind handlers to run, we just want to abort.

No doubt some way out of that could be found, but a better alternative might be to
perform unwinding in a controlled way and let `panic` mean exactly that.

For example, controlled unwinding could return (val, -1) to signal a controlled unwind,
the value carried is the unwind reason, it could be some object with proper fields etc.
This means every evaluation step in the evaluator must check for unwinding, which is not
welcome but is OK I guess for our use.  Also, uses of Invoke() must check.

In the end, error handling will just be invoking the error continuation and all unwinds
will really be invoking a continuation, so it's all call/cc.  The "val" carried could be
a list of values to return from call/cc / apply the continuation to.  In this setup,
the Invoke call within sint:unwind-protect would intercept the throw, invoke the unwind
handler, and then restart the throw.


## Parameters

Scheme parameters are implemented on top of a simple thread-local map tied to the per-thread
context, using a global integer key.

A parameter object can be called with a single argument, which sets the parameter in the current
thread.  The `parameterize` syntax compiles down to `dynamic-wind` plus read and write calls to
the parameter object.

In turn, the thread-local storage is supported by new primitives `sint:new-tls-key`, 
`sint:read-tls-value` and `sint:write-tls-value`.

## Thread safety

### Ports

Ports are built on thread-unsafe data structures (bufio, among other things) and are
therefore protected by a mutex.  Read and write operations acquire exclusive access to
the pertinent streams in the port via protected accessors on the port.

Of course, it's the individual operations (write, display, newline, read, read-char, etc) that
are protected - not groups of these.  So i/o may still be jumbled at the scheme level but
it will be thread-safe at the Go level.

Of course, the lock may make it quite expensive to do things like byte-at-a-time reading.

### Strings

If s and t are strings, s+t is thread-safe (consider the case where one thread does s+t
while another is doing s+u and the fact that strings are simply byte slices; raw `append` on slices
would not be thread-safe).  Safety is
evident from the implementation in the Go run-time support - a truly new string is created.
The language spec is perhaps a little too
understated here: "String addition creates a new string by concatenating the operands."
However, the alternative would probably be too error-prone, since string addition is
used "casually" for error messages and many other things where worrying about thread
safety is simply unreasonable.  We do want effect-free semantics.

## Performance

Performance is awful, but this was expected.  The bulk of the time is spent in memory allocation and garbage collection (indeed in stack walking in the incremental collector); interpreter speed as such is not much of a factor at this time.

There are several (remaining) reasons why the system allocates so much:

- every value (`Val`) is heap-allocated, this being an interface type, and even small integers are represented as boxed values, probably incurring multiple allocations since we use *big.Int for everything, presumably this is a fixed-sized object plus a variable-length array, depending on how optimized it is
- every environment rib is heap-allocated, and since everything is a true function call (no primitives inlined) there is a rib allocated for almost every operation that is not a primitive invocation.  A rib is in turn a fixed-sized lexenv object with a slice of slots attached to it, requiring at least two allocations

The profiler is not yet helpful (to me) in pinpointing which of these causes the bulk of the pain (I'm guessing the ribs are much worse than the values), but they are both bad.

To fix this, several things are necessary.

### Simple values must not be heap-allocated

In principle, `Val` must become a composite immutable value that is not routinely heap-allocated.  For example,
```
type Val struct {
    imm   int64
    other any
}
```
where the smallint has at least a bitfield for type information and another one for small integer-sized values.  It is useful if Val is as small as possible, as it will be copied everywhere.

In the case when the value is a small integer, a rune, null, true, false, unspecified, undefined, or eof, `imm` carries a tag and the value and `other` is nil.  In other cases, `imm` carries the tag and `other` the value, which must be cast from any to a concrete type depending on `tag`.

However, making the value two words makes it very difficult to use concurrency safely.  With single-word values,
every write (set-car!, set-cdr!, vector-set!) is copy-atomic.  With two-word values, not so.  It is possible to use a bit in the immediate as a spinlock, but this might then have to be acquired for reads as well, and synchronization of
values becomes necessary to make sure the words are seen together by all cores.

### Ribs must not be heap-allocated routinely

Lexical ribs must not be heap-allocated except when necessary.  There are several aspects to this.  A more sophisticated compiler will be required to flatten ribs and perform assignment elimination / escape analysis (mutable variables that escape must be boxed).  A more sophisticated runtime must avoid creating ribs and must instead manage variables differently.

### Some results

The time to run fib(30) dropped by 50% by avoiding the creation of slices of values for calls to most primitive operations, and in other redundant cases.  (In practice, almost no primitives take more than two arguments in almost all cases.)
