# Rebuild the compiled libraries
.SUFFIXES: .sch .go
.PHONY: all libs tests
.sch.go:
	go run sint compile $<

TARGETS=runtime/booleans.go runtime/control.go runtime/equivalence.go \
	runtime/exceptions.go runtime/io.go runtime/numbers.go runtime/pairs.go \
	runtime/strings.go runtime/symbols.go runtime/system.go

all:

libs: $(TARGETS)
$(TARGETS): compiler/compiler.go compiler/emitter.go

tests:
	go run sint load tests/booleans.sch tests/chars.sch tests/io.sch 
