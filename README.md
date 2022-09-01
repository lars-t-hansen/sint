# sint

Subset R7RS-small Scheme implementation embedded in Go, with many Go facilities

## Mottos

"Vigor is better than rigor, unless you're already dead."  --Larry Wall

"LISP programmers know the value of everything, and the cost of nothing."  --Alan Perlis

"Learn more Go" --me

## Some details

R7RS-small Scheme embedded in Go with some features (evolving):

- arbitrary precision exact integers (big.Int)
- arbitrary precision inexact rationals (big.Float)
- unicode
- Goroutines
- Channels (in progress)
- Go RegExes (soon)
- synchronization (eventually)
- most standard built-ins (eventually)
- Go FFI (eventually)

A number of subtractions and weirdnesses:

- strings are Go strings, ie immutable byte arrays holding utf8-encoded unicode.  This means a fair amount of string-oriented Scheme code will not work out of the box.  See MANUAL.md for more.
- no exact rationals or exact complexes - I never found these to be useful in practice
- call/cc (not actually implemented yet) is only one-shot and upwards within the same goroutine

Standards conformance is not a goal; but progression toward it is desirable.

Performance is not a concern.  Functionality and easy modifiability are.

## Installation

After cloning the repo, just `go install sint`

## Usage

Try `sint help`

## Near-term TODO

The immediate priority is to get this far enough along to be useful.  This means more types,
some basic ergonomics, and more primitives and library, esp for I/O

### High priority (for MVP)

- implement let-values, because multiple values are ubiquitous
- ports and some I/O, possibly including string ports
- regexes and string matching.  Syntax for literal regex could be #/.../ for example
- high-value number operations, see numbers.sch
- high-value control operations, see control.sch
- load procedure
- exit procedure and maybe emergency-exit
- features, for hack value
- basic error handling & recovery during execution; document it too
- a verb to load and run a file
- clean up repl.go

Maybe also:

- implement select, at least in a limited form

### Backlog (this is actually much longer)

- vectors
- bytevectors
- `dynamic-wind` and its abbreviation (in sint) `unwind-protect`
- everything to do with exceptions
- parameters - these are per-goroutine unwind-protected globals...
- everything to do with environments, if we care
- locks, for safe concurrent access to variables
- maybe some notion of atomic operation, though in scheme this means atomic-set-car!,
  vector-set-car!, etc, and also data-structure specific cmpxchg operations, and
  there would have to be something for globals too - basically this is a fair mess.
  at least common lisp has setf to syntactically merge these things.
- clean up how we do floats.  The exponent range is vast and is not a
  problem, but the default mantissa is only 53 bits.  We should
  consider whether this is the best default (maybe 100 bits?) and
  perhaps also whether it should be configurable somehow.  It's hard
  to do this from Scheme, since values are created at
  hard-to-determine times.
- more testing and bugfixing
- additional primitives
- additional library functions
- the sint/runtime package could provide a Processor abstraction that encapsulates boilerplate?
- Go FFI.  Note plugins as a way of loading code dynamically, but nice also to be able to link in user code statically.
- apropos
- "doc" function (or form) on functions at least
- "source" function on functions
- lots of documentation: variable names, function names, function comments, doc strings, ...
- inexact complexes
- all missing special forms
