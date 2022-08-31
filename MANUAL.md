# sint User Manual

`sint` is an evolving variant of R7RS-small Scheme, embedded in Go.  Eventually
it will have a Go FFI, regular expressions, and many other facilities.

`sint` is not meant to be performant (yet), it is more important that
it is easy to evolve and play with.

`sint` is not meant to be all of R7RS-small, or to be completely compatible
with it ever, but is meant to approach R7RS-small over time, to the extent
possible.

## Deliberate restrictions to and incompatibilities with R7RS-small Scheme

Numbers are exact integers (math/big.Int) and inexact rationals (math/big.Float).

`call-with-current-continuation` produces continuations that are one-shot, upward-only, and only usable within the same goroutine.  That is, these are strictly for same-thread nonlocal jumps.  Many other uses of first-class continuations, eg generators and threads, can be implemented using goroutines.

Strings are Go strings, ie, they are immutable byte arrays holding UTF8-encoded Unicode text, and indices into strings are byte indices, not character indices.  Read more about this below.

## Strings

`sint` strings are Go strings, ie, they are immutable byte slices containing UTF-8 encoded Unicode code points and possibly other values.  They are indexed using byte indices, which means it's possible to index a string in the middle of a code point.  `sint` characters are however restricted to be valid code points always.  `sint` character accessors (actually most string operations, essentially anything that can't be an operation on byte slices) will throw if encountering an invalid encoding; non-character accessors `string=?`, `string>?`, `string>=?`, `string<?`, `string<=?`, `string-append` and `string-copy` in principle need not check that the encoding is valid and probably should not.

Since strings are immutable, all the string update functions -- `string-set!`, `string-fill!`, and `string-copy!` -- are missing.

### Whole-string operations

The following operate on whole strings and have the same semantics as the same-named functions in R7RS Scheme, except that those that need to access individual characters will throw if encountering invalid code points, as noted above:

* (string char ...)
* string?
* make-string
* string=?
* string>?
* string>=?
* string<?
* string<=?
* string-ci=?
* string-ci>?
* string-ci>=?
* string-ci<?
* string-ci<=?
* string-append
* string-map
* string-for-each
* string->vector
* vector->string
* string->list (whole-string form)
* string->copy (whole-string form)
* list->string
* string-upcase
* string-downcase
* string-foldcase

### Byte-index operations

The following operate on byte indices and lengths and perhaps should have new names

* string-length - returns byte length
* string-ref - returns a decoded code point starting at the given byte index AND the size in bytes of that code point's encoding
* substring - returns a string if the byte indices are proper for full characters
* string-copy (substring form) - as for substring
* string->list (substring form) - as for substring

### Mutators

The mutators operators are missing:

* string-set!
* string-copy!
* string-fill!

### New procedures

Not sure what we want here.  We may want a `string-ref` that can report a decoding error rather than failing.  We may want to surface some of the Go string operations (searching, replacing, slicing).

## Goroutines (evolving)

The form `(go (expr expr ...))` is syntax.  The exprs are evaluated on the current thread; the first expr must evaluate to a procedure of the appropriate arity; the procedure is invoked on the argument values on a new, concurrent thread.  If the procedure returns, its thread terminates and any return value is discarded.

The memory model is that of Go.  All Scheme values are word-sized and racy accesses have sensible outcomes with no torn values, provided the implementation does not detect the race and terminate the program.  Atomic operations and channels (see subsequent sections) can be used to synchronize concurrent access to avoid races.

TODO: It's possible we want some type of primitive to identify the goroutine we're in, like a thread ID?

## Channels (evolving)

There are channels that can transmit all scheme values.  `(make-input-channel capacity)`, `(make-output-channel capacity)`, and `(make-channel capacity)` create the channels.  The capacity is optional.

Channels are closed with the `(close-channel ch)` operator.

TODO: The GC must close a channel when the channel is reaped.  Don't know how to do that (yet), maybe it happens already.

Operations on channels: `(send ch val)`, `(receive ch)`.

The `select` operation is however tricky.  It can use an arbitrary number of channels but those must be written into the statement, there isn't a notion of higher-order select clauses.  We could probably have a hack where we have "up to n" but the combinations of sends and receives means the number of variants is exponential in n.  And it's not obvious to me yet that it's possible to have an array of channels?  It ought to be...

## Atomics (evolving)

TBD.
