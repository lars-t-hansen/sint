# sint

Subset R7RS-small Scheme implementation, embedded in Go

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

Standards conformance is not a goal; but progression toward it is desirable.

Performance is not a concern.  Functionality and easy modifiability are.

## Near-term TODO

The immediate priority is to get this far enough along to be useful.  This means more types,
some basic ergonomics, and more primitives and library, esp for I/O

High priority

- integer division operators and maybe other numerics
- clean up how we do floats.  The exponent range is vast and is not a
  problem, but the default mantissa is only 53 bits.  We should
  consider whether this is the best default (maybe 100 bits?) and
  perhaps also whether it should be configurable somehow.  It's hard
  to do this from Scheme, since values are created at
  hard-to-determine times.
- ports and I/O, including string ports
- basic error handling & recovery during execution

Medium priority

- more testing and bugfixing
- a verb to load and run a file - eases testing, also delve
- a verb to print help
- a number of primitives
- many more library functions, compiled-in
- the sint/runtime package could provide a Processor abstraction that encapsulates boilerplate?
- regexes and string matching.  Syntax for literal regex could be #/.../ for example

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


