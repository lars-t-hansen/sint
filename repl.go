// Command line parsing and command implementation.  This is still pretty rough.

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

var HelpText string = `
Usage:

  sint
  sint repl
    Enter the interactive repl

  sint eval expr
    Evaluate the expression, print its result(s), and exit.

  sint load filename.sch
    Load filename.sch: read its expressions, evaluate them in order and
	print their results, and exit after the last.

  sint compile filename.sch
    Compile filename.sch into filename.go and exit.  The output will have
    a function initFilename() that takes a *Scheme and evaluates the
    expressions and definitions of filename.sch in order in that runtime.
    Punctuation in the filename is removed.

  sint help
    Print help (this text)
`

func main() {
	engine := core.NewScheme(nil)
	comp := compiler.NewCompiler(engine.Shared)

	args := os.Args[1:]

	if len(args) == 0 {
		enterRepl(engine, comp)
		return
	}

	switch args[0] {
	case "compile":
		// Obviously it would be meaningful to have multiple file names.
		if len(args) != 2 {
			panic("Bad 'compile' command, one file name argument required")
		}
		compileFile(engine, comp, args[1])
	case "eval":
		// An idea is that "eval" is the default verb if the first letter of
		// the verb is left paren.  That way, `sint 'some expr'` will evaluate
		// and print it.
		//
		// Another idea is that there could be a sequence of expressions, not just one.
		if len(args) != 2 {
			panic("Bad 'eval' command, exactly one expression argument required")
		}
		evalExpr(engine, comp, args[1])
	case "help":
		fmt.Print(HelpText)
	case "load":
		// There could be multiple files too
		if len(args) != 2 {
			panic("Bad 'load' command, exactly one file required")
		}
		loadFile(engine, comp, args[1])
	case "repl":
		enterRepl(engine, comp)
	default:
		panic("Bad command arguments, try `sint help`")
	}
}

func enterRepl(engine *core.Scheme, comp *compiler.Compiler) {
	runtime.InitPrimitives(engine)
	runtime.InitCompiled(engine)
	reader := runtime.NewStdinReader()
	writer := runtime.NewStdoutWriter()
	engine.SetTlsValue(core.CurrentOutputPort, core.NewOutputPort(writer, true /* isText */, "<standard output>"))
	engine.SetTlsValue(core.CurrentInputPort, core.NewInputPort(reader, true /* isText */, "<standard input>"))
	for {
		writer.WriteString("> ")
		v, rdrErr := runtime.Read(engine, reader)
		if rdrErr != nil {
			os.Stderr.WriteString(rdrErr.Error() + "\n")
			continue
		}
		if v == engine.EofVal {
			writer.WriteRune('\n')
			break
		}
		prog, progErr := comp.CompileToplevel(v)
		if progErr != nil {
			os.Stderr.WriteString(progErr.Error() + "\n")
			continue
		}
		//writer.WriteString(prog.String() + "\n")
		results, unw := engine.EvalToplevel(prog)
		if unw != nil {
			// Last-ditch error handler.  With a little more sophistication, there
			// will be a call/cc to catch the error and we won't reach this code.
			os.Stderr.WriteString("ERROR: " + unw.String() + "\n")
			continue
		}
		for _, r := range results {
			if r != engine.UnspecifiedVal {
				runtime.Write(r, false, writer)
				writer.WriteRune('\n')
			}
		}
	}
}

func evalExpr(engine *core.Scheme, comp *compiler.Compiler, expr string) {
	runtime.InitPrimitives(engine)
	runtime.InitCompiled(engine)
	reader := bufio.NewReader(strings.NewReader(expr))
	writer := runtime.NewStdoutWriter()
	v, rdrErr := runtime.Read(engine, reader)
	if rdrErr != nil {
		os.Stderr.WriteString(rdrErr.Error() + "\n")
		os.Stderr.WriteString("Aborting\n")
		os.Exit(1)
	}
	prog, progErr := comp.CompileToplevel(v)
	if progErr != nil {
		os.Stderr.WriteString(progErr.Error() + "\n")
		os.Stderr.WriteString("Aborting\n")
		os.Exit(1)
	}
	results, unw := engine.EvalToplevel(prog)
	if unw != nil {
		// Last-ditch error handler.  With a little more sophistication, there
		// will be a call/cc to catch the error and we won't reach this code.
		os.Stderr.WriteString("ERROR: " + unw.String() + "\n")
		os.Stderr.WriteString("Aborting\n")
		os.Exit(1)
	}
	for _, r := range results {
		if r != engine.UnspecifiedVal {
			runtime.Write(r, false, writer)
			writer.WriteRune('\n')
			writer.Flush()
		}
	}
}

func loadFile(engine *core.Scheme, comp *compiler.Compiler, fn string) {
	runtime.InitPrimitives(engine)
	runtime.InitCompiled(engine)
	input, inErr := os.Open(fn)
	if inErr != nil {
		panic(inErr)
	}
	reader := bufio.NewReader(input)
	writer := runtime.NewStdoutWriter()
	for {
		v, rdrErr := runtime.Read(engine, reader)
		if rdrErr != nil {
			os.Stderr.WriteString(rdrErr.Error() + "\n")
			os.Stderr.WriteString("Aborting\n")
			os.Exit(1)
		}
		if v == engine.EofVal {
			break
		}
		prog, progErr := comp.CompileToplevel(v)
		if progErr != nil {
			os.Stderr.WriteString(progErr.Error() + "\n")
			os.Stderr.WriteString("Aborting\n")
			os.Exit(1)
		}
		results, unw := engine.EvalToplevel(prog)
		if unw != nil {
			// Last-ditch error handler.  With a little more sophistication, there
			// will be a call/cc to catch the error and we won't reach this code.
			os.Stderr.WriteString("ERROR: " + unw.String() + "\n")
			os.Stderr.WriteString("Aborting\n")
			os.Exit(1)
		}
		for _, r := range results {
			if r != engine.UnspecifiedVal {
				runtime.Write(r, false, writer)
				writer.WriteRune('\n')
			}
		}
	}
	input.Close()
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
		v, rdrErr := runtime.Read(engine, reader)
		if rdrErr != nil {
			os.Stderr.WriteString(rdrErr.Error() + "\n")
			os.Stderr.WriteString("Aborting\n")
			os.Exit(1)
		}
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
		fmt.Fprintf(writer, "_, unw%s := c.EvalToplevel(%s)\n", initName, initName)
		fmt.Fprintf(writer, "if unw%s != nil { panic(unw%s.String()) }\n", initName, initName)
	}
	writer.WriteString("}\n")
	writer.Flush()
	input.Close()
	tmp.Close()
	os.Rename(tmpFn, outFn)
}
