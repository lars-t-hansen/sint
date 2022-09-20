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

## Select

```
(select ((send send-ch-expr send-value-expr) send-body-expr ...)
        ((receive (recv-value-var recv-status-var) recv-ch-expr) receive-body-expr ...)
        (else else-expr ...))
```

As in Go, this would evaluate the send-ch-expr, the send-value-expr, and the recv-ch-expr, and only when everything has been evaluated would it pick a channel, if any, to operate on, and then to evaluate the body expressions of the appropriate clause.

Implementing this requires the use of reflection in the general case: SelectCase, SelectDir, and Select.  However, in the common case there will be few cases and we can have ready-made select functions that just open-code the select statement appropriately.  For four cases + default we need 16 functions.  If we have fewer than four cases we can pad with dummy channels that are never ready.

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
