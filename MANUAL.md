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

## Goroutines

The form `(go (expr expr ...))` is syntax.  The exprs are evaluated on the current thread; the first expr must evaluate to a procedure of the appropriate arity; the procedure is invoked on the argument values on a new, concurrent thread.  If the procedure returns, its thread terminates and any return value is discarded.

The memory model is that of Go.  All Scheme values are word-sized and racy accesses have sensible outcomes with no torn values, provided the implementation does not detect the race and terminate the program.  Atomic operations and channels (see subsequent sections) can be used to synchronize concurrent access to avoid races.

* `(goroutine-id)` -> exact nonnegative integer, an ID for the current goroutine

## Channels and communication (evolving)

### Channel primitives

The channels can transmit all scheme values.  Operations on channels:

* `(make-input-channel [capacity])` -> bidirectional channel, default capacity is 0
* `(channel? obj)` -> boolean
* `(channel-send ch val)` -> unspecified; panics if the channel is closed
* `(channel-receive ch)` -> (value, status) as in Go
* `(channel-length ch)` -> number of items in the channel
* `(channel-capacity ch)` -> capacity of channel's buffer
* `(close-channel ch)` -> unspecified

In Go there is additionally a type system to constrain the view of a channel to the input side or the output side, but I've not implemented that here.  If we were to do it, it would be a `(make-input-channel channel)` type of operation,
and there would be additional predicates.

TODO: The GC must close a channel when the channel object is reaped.  Don't know how to do that (yet), maybe it happens already.

### Select (not yet implemented)

In principle there should be a `select` as in Go to allow communication to proceed on some channel that is ready:

```
(select ((send send-ch-expr send-value-expr) send-body-expr ...)
        (((recv-value-var recv-status-var) receive recv-ch-expr) receive-body-expr ...)
        (else else-expr ...))
```

In the case of receive, either variable can be the identifier _.  (Generally, a binding for name _ should be a shorthand for a gensym'd name.)

As in Go, this would evaluate the send-ch-expr, the send-value-expr, and the recv-ch-expr, and only when everything has been evaluated would it pick a channel, if any, to operate on, and then to evaluate the body expressions of.

Implementing this is a little tricky.  It can use an arbitrary number of channels but those must be written into the statement in Go, there isn't a notion of higher-order select clauses.  We could probably have a hack where we have "up to n" but the combinations of sends and receives means the number of variants is exponential in n, so this is limited to about five, in practice.

If we have a channel that always delivers a value then it may be possible to chain those fixed-size nests.  Closed channels deliver the null value, ie, nil, which is basically perfect for this.  This dummy channel will be selected with probability 1/n for n-1 cases.  Thus a second-level chain is selected with probability 1/n, and the cases in that will have probability 1/n*1/m where m is the size of the second level - ie, the probabilities are not going to be uniform.   To fix that, if we need more than one level then the leaves - the true channels - must all be at the same level, ie, say we have up to five for a baked in select stmt and we have seven channels to select from.  Then a first-level switch will have two dummy input channels, selecting one of two lower-level switches with equal priority, and these (one of size four, the other of size three, maybe) will then select on the true channels.  It is true that the priorities will still not be the same, but they will be close to each other.

## Synchronization and atomics (evolving)

TBD.

## Numbers

### Bitwise operations

Bitwise operations are interpreted as if on two's complement integers; the arguments must be exact integers.

* `(bitwise-and n ...)`
* `(bitwise-or n ...)`
* `(bitwise-xor n ...)`
* `(bitwise-and-not n1 n2)` computes `(bitwise-and n1 (bitwise-not n2))`
* `(bitwise-not n)`

## Lists

* `(list-sort! <? xs)` sorts list `xs` in-place according to binary predicate `<?`
* `(list-sorted? <? xs)` returns #t iff list `xs` is sorted according to binary predicate `<?`
* `(some? pred xs)` returns #t iff pred tests true for one of the elements of list `xs`
* `(every? pred xs)` returns #t iff pred tests true for all of the elements of list `xs`
* `(filter pred xs)` returns those elements of list `xs` for which `pred` tests true

## Generators

* `(make-generator p [end]) => thunk` takes a procedure `p` of one argument, `yield`, and optionally an `end` value.  `p` is invoked once and must call `yield` on values to generate them.  Calls to the returned thunk retrieve successively yielded values.  If `end` is present it is yielded once `p` returns.  After that, the generator yields `#!unspecified` indefinitely.  See samples/generator.sch.

## Introspection and reflection

* `(apropos pattern)` prints (on the current output port) the sorted result of `(filter-global-variables pattern)`, one name per line.
* `(filter-global-variables pattern)` where `pattern` is a string or a symbol returns an unsorted list of the names (symbols) of all the global variables whose names have `pattern` as a substring.
* `(procedure-name proc)` returns the name of the procedure as a string, derived from the source code
* `(procedure-arity proc)` returns the arity of the procedure as a number.  The number denotes the number of fixed arguments of the procedure; it is inexact iff the procedure accepts rest arguments.
* `(symbol-has-value? symbol)` returns true iff there is a global variable with the name `symbol`
* `(symbol-value symbol)` returns the value of the global variable with the name `symbol`.  It throws if there is no such variable.
* `(doc <object>)` returns documentation about the object, normally as a list `(<tag> ...)` where the `<tag>` is a symbol denoting the class of the object (`procedure`, and so on), and the rest of the values depend on the object and the documentation available about it.  In the case when nothing interesting is known it returns `(datum <object>)`.

## Input and output

* `(for-each-line proc input)` applies proc to each line of the text input port `input`
* `(filter-lines proc input)` applies proc to each line of the text input port `input`, returning a list of the lines for which `proc` returns a truthy answer
