# -*- fill-column: 100 -*-
.PHONY: all libs tests

all:


# Rebuild the compiled libraries.
#
# This is not totally safe since the files are built sequentially and a compiler bug will render the
# system inoperable.  Better would be to "install" the current version first and then run that.

.SUFFIXES: .sch .go
.sch.go:
	go run sint compile $<

TARGETS=runtime/booleans.go runtime/control.go runtime/equivalence.go \
	runtime/exceptions.go runtime/io.go runtime/numbers.go runtime/pairs.go \
	runtime/strings.go runtime/symbols.go runtime/system.go \
	runtime/generator.go runtime/sort.go

libs: $(TARGETS)
$(TARGETS): compiler/compiler.go compiler/emitter.go runtime/reader.go


# Run all test cases in the current development system.  Note tests/asserts.sch must be loaded
# first.
#
# This is probably a little brittle since all tests are loaded into the same system; there could be
# interference between the different tests.

tests:
	go run sint load tests/asserts.sch tests/booleans.sch tests/chars.sch tests/concurrency.sch \
		tests/control.sch tests/equivalence.sch tests/io.sch tests/numbers.sch tests/pairs.sch \
		tests/regexp.sch \
		tests/strings.sch tests/symbols.sch tests/system.sch tests/syntax.sch tests/generator.sch \
		tests/sort.sch 

