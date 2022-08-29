# sint User Manual

`sint` is an evolving variant of Scheme, embedded in Go.  Eventually
it will have a Go FFI, regular expressions, and many other facilities.

`sint` is not meant to be performant (yet), it is more important that
it is easy to evolve and play with.

## Deliberate restrictions to and incompatibilities with R7RS-small Scheme

Numbers are only exact integers (math/big.Int) and inexact rationals
(math/big.Float).

`call-with-current-continuation` produces continuations that are
one-shot, upward-only, and (eventually) only usable within the same
goroutine.  That is, these are strictly for same-thread nonlocal
jumps.  Many other uses of first-class continuations, eg generators
and threads, can be implemented using goroutines.

Strings are Go strings, ie, they are immutable byte-indexed byte
arrays holding UTF8-encoded Unicode text.  Read more about this below.

## Strings

`sint` strings are Go strings, ie, they are immutable byte slices
containing UTF-8 encoded Unicode code points.

Go strings can contain invalid encodings, not sure if we want that
here.  By restricting character values to valid Unicode and not
allowing non-code point values in string literals, and by also
checking that on input, we guarantee that there are no invalid code
points.  Not sure yet if that's worth it - experimenting.  NOTE that
substring and the substring forms of string-copy and string->list are
able to start in the middle of an encoding and may thus produce
garbage, or they must check that this does not happen.  Ditto,
string-ref may provide a byte index not at the start of a character.

Thus our strings are immutable and weirdly indexed; ie, they are quite
incompatible with standard Scheme strings.  So the (unresolved)
question is whether to call the type 'string' or use a new name, eg
'gstring' (bad), 'str' (why not?), 'gostring', and so on.  For now, it
is "string".  Another alternative is to give the procedures that have
new semantics new names.

The library:

These operate on characters and should be indistinguishable from the
same-named functions in normal Scheme:

(string char ...)
string?
make-string  ;; not very useful
string=?
string>?
string>=?
string<?
string<=?
string-ci=?
string-ci>?
string-ci>=?
string-ci<?
string-ci<=?
string-append
string-map
string-for-each
string->vector
vector->string
string->list (whole-string form)
string->copy (whole-string form)
list->string
string-upcase
string-downcase
string-foldcase

These operates on byte indices and lengths and perhaps should have new names

string-length - returns byte length
string-ref - returns a decoded code point starting at the given byte index
substring - returns a string if the byte indices are proper for full characters
string-copy (substring form) - as for substring
string->list (substring form) - as for substring

Mutators are missing:

string-set!
string-copy!
string-fill!

It is likely we would want some new procedures, to compensate for immutability.  Splicing,
replacing, and searching are obvious.  Decoding an UTF8 char at byte index ditto.
