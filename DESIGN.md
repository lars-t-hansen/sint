# Design notes

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
