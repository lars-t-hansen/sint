# sint

Extended subset of R7RS-small Scheme, embedded in Go, with many Go facilities

The purpose here is not really to have another Scheme system, but rather to explore Go.

## Mottos

"Vigor is better than rigor, unless you're already dead."  --Larry Wall

"LISP programmers know the value of everything, and the cost of nothing."  --Alan Perlis

"You can't think about thinking without thinking about thinking about something." --Seymour Papert

## Status

Pre-MVP, but basically operational.  See issue tracker for all MVP issues, and BACKLOG.md for longer-term things.

## Some details

R7RS-small Scheme embedded in Go with some features (evolving):

- arbitrary precision exact integers (big.Int)
- arbitrary precision inexact rationals (big.Float)
- unicode
- goroutines
- channels
- Go regexes (soon)
- synchronization (eventually)
- most standard Scheme built-ins (eventually)
- Go FFI (eventually)

A number of subtractions and weirdnesses:

- strings are Go strings, ie immutable byte arrays holding utf8-encoded unicode.  This means a fair amount of string-oriented Scheme code will not work out of the box.  See MANUAL.md for more.
- no exact rationals or exact complexes - I never found these to be useful in practice
- call/cc is only one-shot and upwards within the same goroutine.  Many other uses for call/cc (generators, threads, coroutines) are better implemented as goroutines+channels, and call/cc+mutation is a nightmare anyway.
- no first-class environments, and no protected primitives.  Everything is defined in an open top-level scope, no primitives are inlined anywhere, and library functions use standard procedures freely.  You can redefine CAR - though you probably shouldn't!

R7RS conformance is not a goal; but progression toward it is desirable.

Performance is not a concern (modulo obvious stupidity).  Functionality and easy modifiability are.

## Installation

After cloning the repo, just `go install sint`

## Usage

Try `sint help`
