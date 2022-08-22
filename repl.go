// Usage:
//
//  sint
//  sint repl
//    Enter the interactive repl
//
//  sint compile filename.sch
//    Compile filename.sch into filename.go with a function initFilename() that
//    takes a *Scheme and evaluates the expressions and definitions of filename.sch
//    in order in that runtime.  Punctuation in the filename is removed.

package main

import (
	"bufio"
	"os"
	"sint/compiler"
	"sint/core"
	"sint/runtime"
	"strings"
)

func main() {
	engine := core.NewScheme()
	comp := compiler.NewCompiler(engine)

	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "compile" {
			if len(args) == 2 {
				compileFile(engine, comp, args[1])
				return
			}
			panic("Bad 'compile' command")
		}
		if args[0] == "repl" {
			enterRepl(engine, comp)
			return
		}
		panic("Bad arguments")
	} else {
		enterRepl(engine, comp)
	}
}

func enterRepl(engine *core.Scheme, comp *compiler.Compiler) {
	runtime.InitPrimitives(engine)
	runtime.InitCompiled(engine)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for {
		writer.WriteString("> ")
		writer.Flush()
		v := runtime.Read(engine, reader)
		if v == engine.EofVal {
			break
		}
		// TODO: Recover from compilation error
		prog := comp.CompileToplevel(v)
		writer.WriteString(prog.String() + "\n")
		writer.Flush()
		// TODO: Recover from runtime error
		result := engine.EvalToplevel(prog)
		if result != engine.UnspecifiedVal {
			runtime.Write(result, writer)
			writer.WriteRune('\n')
			writer.Flush()
		}
	}
}

func compileFile(engine *core.Scheme, comp *compiler.Compiler, fn string) {
	if strings.LastIndex(fn, ".sch") != len(fn)-4 {
		panic("Input file for 'compile' must have type '.sch'")
	}
	tmpFn := fn[:len(fn)-4] + ".tmp"
	outFn := fn[:len(fn)-4] + ".go"
	os.Stdout.WriteString(tmpFn + "\n")
	os.Stdout.WriteString(outFn + "\n")
	input, inErr := os.Open(fn)
	if inErr != nil {
		panic(inErr)
	}
	tmp, tmpErr := os.Create(tmpFn)
	if tmpErr != nil {
		panic(tmpErr)
	}
	/*
		defer {
			if tmp != nil {
				tmp.Close()
				os.Remove(tmpFn)
			}
		}
	*/
	// Compute output file from input file
	// Open input file
	// Open temp output file, remove on exit if failure
	// Read expressions
	// Compile them
	// Emit them as Go code
	// Close files
	// Rename temp file as final output file
}
