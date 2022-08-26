# sint

Subset scheme implementation embedded in Go

## Mottos

"Vigor is better than rigor, unless you're already dead."  --Larry Wall

"LISP programmers know the value of everything, and the cost of nothing."  --Alan Perlis

"Learn more Go" --me

## Some details

R7RS-small Scheme embedded in Go with some features:

- arbitrary precision exact integers
- arbitrary precision inexact reals (well, maybe rationals, but ...)
- unicode
- most standard built-ins (eventually)
- some additional built-ins
- Go FFI (eventually)
- Goroutines (eventually)

A number of subtractions and weirdnesses:

- strings are Go strings, ie immutable byte arrays holding utf8.  This means a fair amount of string-oriented Scheme code will not work out of the box.  See below.
- no exact rationals or exact complexes - i never found these to be useful in practice
- call/cc is only one-shot and upwards within the same goroutine


## Near-term TODO

The immediate priority is to get this far enough along to be useful.  This means more types,
some basic ergonomics, and more primitives and library, esp for I/O

High priority

- strings - add enough primitives to make them useful
- apply - design done, only some implementation left
- multiple values - this will be disruptive
- basic error handling & recovery
- more testing and bugfixing
- a verb to load and run a file - eases testing, also delve
- a verb to evaluate some code - eases testing, also delve

Medium priority

- a verb to print help
- ports and I/O, including string ports
- a number of primitives
- many more library functions, compiled-in
- the sint/runtime package could provide a Processor abstraction that encapsulates boilerplate?

## Longer-term TODO

- goroutines and channels that can transmit scheme values
- some notion of what mutation means in the context of concurrency.  atm, all values are pointer-sized, which is pretty good, but what does go's memory model do with unsynchronized concurrent access?
- Go FFI.  Note plugins as a way of loading code dynamically, but nice also to be able to link in user code statically.
- apropos
- "doc" function (or form) on functions at least
- "source" function on functions
- lots of documentation: variable names, function names, function comments, doc strings, 

## Some notes on design

### Performance

Performance is not a concern.  Functionality is.

### Strings

Sint strings are Go strings, ie, they are immutable byte slices containing UTF-8 encoded
Unicode code points.

Go strings can contain invalid encodings, not sure if we want that here.
By restricting character values to valid Unicode and not allowing non-code point values
in string literals, and by also checking that on input, we guarantee that there are no
invalid code points.  Not sure yet if that's worth it - experimenting.  NOTE that substring
and the substring forms of string-copy and string->list are able to start in the middle
of an encoding and may thus produce garbage, or they must check that this does not happen. 
Ditto, string-ref may provide a byte index not at the start of a character.

Thus our strings are immutable and weirdly indexed; ie, they are quite incompatible with
standard Scheme strings.  So the (unresolved) question is whether to call the type 'string'
or use a new name, eg 'gstring' (bad), 'str' (why not?), 'gostring', and so on.  For now,
it is "string".  Another alternative is to give the procedures that have new semantics
new names.

The library:

These operate on characters and should be indistinguishable from the same-named functions
in normal Scheme:

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