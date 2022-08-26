# Rebuild the compiled libraries
.SUFFIXES: .sch .go
.PHONY: all
.sch.go:
	go run sint compile $<

TARGETS=runtime/booleans.go runtime/control.go runtime/equivalence.go runtime/symbols.go

all: $(TARGETS)
$(TARGETS): compiler/compiler.go compiler/emitter.go
