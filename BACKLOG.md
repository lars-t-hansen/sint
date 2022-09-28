# Ad-hoc backlog

This is a grab-bag of TODO items that are not yet prominent enough to make it into the issue tracker.  Probably there are many more items that could go on this list.

All MVP backlog items are in the issue tracker.

Also see many TODO comments in the source code.

- regexes and string matching.  Syntax for literal regex could be #/.../ for example
- implement `select`, at least in a limited form
- vectors
- bytevectors
- everything to do with exceptions / conditions
- everything to do with environments, if we care
- locks, for safe concurrent access to variables
- maybe some notion of atomic operation, though in scheme this means atomic-set-car!,
  atomic-vector-set!, etc, and also data-structure specific cmpxchg operations, and
  there would have to be something for globals too - basically this is a fair mess.
  at least common lisp has setf to syntactically merge these things.
- clean up how we do floats.  The exponent range is vast and is not a
  problem, but the default mantissa is only 53 bits.  We should
  consider whether this is the best default (maybe 100 bits?) and
  perhaps also whether it should be configurable somehow.  It's hard
  to do this from Scheme, since values are created at
  hard-to-determine times.
- Go FFI.  Note plugins as a way of loading code dynamically, but nice also to be able to link in user code statically.
- "source" function on functions
- lots of documentation: variable names, function names, function comments, doc strings, ...
- inexact complexes
- compiled code can probably use init() to register an init callback that will be invoked by initCompiledCode, so that the latter need not be updated every time we add a new file

