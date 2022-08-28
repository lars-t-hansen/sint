// Usage:
//
//  sint
//  sint repl
//    Enter the interactive repl
//
//  sint eval expr
//    Evaluate the expression, print its result, and exit
//
//  sint compile filename.sch
//    Compile filename.sch into filename.go with a function initFilename() that
//    takes a *Scheme and evaluates the expressions and definitions of filename.sch
//    in order in that runtime.  Punctuation in the filename is removed.

package main

import (
	"bufio"
	"fmt"
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
		// An idea is that "eval" is the default verb, or perhaps,
		// it is the default verb if the first letter of the verb
		// is left paren.  That way, `sint 'some expr'` will evaluate
		// and print it.
		if args[0] == "eval" {
			evalExpr(engine, comp, args[1])
			return
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
	writer := &runtime.StdoutWriter{}
	for {
		writer.WriteString("> ")
		v := runtime.Read(engine, reader)
		if v == engine.EofVal {
			writer.WriteRune('\n')
			break
		}
		// TODO: Recover from compilation error
		prog, progErr := comp.CompileToplevel(v)
		if progErr != nil {
			os.Stderr.WriteString(progErr.Error() + "\n")
			continue
		}
		writer.WriteString(prog.String() + "\n")
		result := engine.EvalToplevel(prog)
		if prog == nil {
			continue
		}
		if result != engine.UnspecifiedVal {
			runtime.Write(result, false, writer)
			writer.WriteRune('\n')
		}
	}
}

func evalExpr(engine *core.Scheme, comp *compiler.Compiler, expr string) {
	runtime.InitPrimitives(engine)
	runtime.InitCompiled(engine)
	reader := bufio.NewReader(strings.NewReader(expr))
	writer := bufio.NewWriter(os.Stdout)
	v := runtime.Read(engine, reader)
	prog, progErr := comp.CompileToplevel(v)
	if progErr != nil {
		os.Stderr.WriteString(progErr.Error() + "\n")
		os.Stderr.WriteString("Aborting\n")
		os.Exit(1)
	}
	result := engine.EvalToplevel(prog)
	if result != engine.UnspecifiedVal {
		runtime.Write(result, false, writer)
		writer.WriteRune('\n')
		writer.Flush()
	}
}

func compileFile(engine *core.Scheme, comp *compiler.Compiler, fn string) {
	if strings.LastIndex(fn, ".sch") != len(fn)-4 {
		panic("Input file for 'compile' must have type '.sch'")
	}
	withoutExt := fn[:len(fn)-4]
	ix := strings.LastIndexAny(withoutExt, "/\\")
	moduleName := withoutExt
	if ix != -1 {
		moduleName = moduleName[ix+1:]
	}
	if len(moduleName) == 0 {
		panic("Input file name is empty")
	}
	moduleName = strings.ToUpper(moduleName[0:1]) + strings.ToLower(moduleName[1:])
	tmpFn := withoutExt + ".tmp"
	outFn := withoutExt + ".go"
	input, inErr := os.Open(fn)
	if inErr != nil {
		panic(inErr)
	}
	// TODO: Use proper tempfiles
	tmp, tmpErr := os.Create(tmpFn)
	if tmpErr != nil {
		panic(tmpErr)
	}
	// TODO: Remove the tempfile on error / early exit
	/*
		defer {
			if tmp != nil {
				tmp.Close()
				os.Remove(tmpFn)
			}
		}
	*/
	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(tmp)
	fmt.Fprintf(writer, `
// Generated from %s
package runtime
import (
	. "sint/core"
	"math/big"
)
func dummy%s() {
	// Make sure the imports are used, or the Go compiler barfs.
	var _ Val = big.NewInt(0)
}
func init%s(c *Scheme) {
`, fn, moduleName, moduleName)
	id := 1
	for {
		v := runtime.Read(engine, reader)
		if v == engine.EofVal {
			break
		}
		prog, progErr := comp.CompileToplevel(v)
		if progErr != nil {
			os.Stderr.WriteString(progErr.Error() + "\n")
			os.Stderr.WriteString("Aborting\n")
			os.Exit(1)
		}
		initName := fmt.Sprintf("code%d", id)
		id++
		compiler.EmitGo(prog, initName, writer)
		fmt.Fprintf(writer, "c.EvalToplevel(%s)\n", initName)
	}
	writer.WriteString("}\n")
	writer.Flush()
	input.Close()
	tmp.Close()
	os.Rename(tmpFn, outFn)
}
