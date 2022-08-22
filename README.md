# sint
Subset scheme implementation in Go

To help prioritize:
- "Learn more go" -- emphasize things that do that, eg, I/O, efficiency

Near-term TODO
- more testing and bugfixing
- compile to Go code on disk, and figure out how to incorporate this
- some kind of loader from file
- strings
- characters
- a number of primitives
- eval & apply
- multiple values
  - let-values syntax
  - values primitive
  - call-with-values primitive
  - probably each function starts returning two values, the first value plus a flag indicating additional values
  - "values" stores additional values in secret slots
  - "call-with-values" interprets the flag and the slots
  - single-value continuations ignore the flag and the secret slots
- library functions, compiled-in
- some kind of repl
- probably want to represent ports somehow, including string ports
- basic error handling
- the sint/runtime package could provide a Processor abstraction that encapsulates boilerplate

Longer-term(?) TODO
- goroutines and channels
