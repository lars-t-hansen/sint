// "sint" - an extended subset of R7RS scheme, embedded in Go, with many Go facilities.
//
// Command line parsing and command implementation.

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
"sint" - an extended subset of Scheme, embedded in Go
v0.1 (pre-mvp)

Usage:

  sint
  sint repl
    Enter the interactive repl

  sint eval expr ...
    Evaluate the expressions, print their result(s), and exit after the last.

  sint load filename.sch ...
    Load filename.sch: read its expressions, evaluate them in order and
    print their results, and exit after the last expression of the last file.

  sint compile filename.sch ...
    Compile each filename.sch into filename.go and exit.  The output will have
    a function initFilename() that takes a *Scheme and evaluates the
    expressions and definitions of filename.sch in order in that runtime.

  sint help
    Print help (this text)
`

func main() {
	engine := core.NewScheme(nil, nil)
	comp := compiler.NewCompiler(engine.Shared)

	args := os.Args[1:]

	if len(args) == 0 {
		enterRepl(engine, comp)
		return
	}

	switch args[0] {
	case "compile":
		if len(args) < 2 {
			panic("Bad 'compile' command, at least one file name argument required")
		}
		for _, fn := range args[1:] {
			err := compileFile(engine, comp, fn)
			if err != nil {
				panic(err)
			}
		}
	case "eval":
		if len(args) < 2 {
			panic("Bad 'eval' command, at least one expression argument required")
		}
		for _, ex := range args[1:] {
			err := evalExpr(engine, comp, ex)
			if err != nil {
				if unw, ok := err.(*core.UnwindPkg); ok {
					reportUnwinding(engine, unw)
					os.Exit(1)
				} else {
					panic(err)
				}
			}
		}
	case "load":
		if len(args) < 2 {
			panic("Bad 'load' command, at least one file name argument required")
		}
		for _, fn := range args[1:] {
			err := loadFile(engine, comp, fn)
			if err != nil {
				if unw, ok := err.(*core.UnwindPkg); ok {
					reportUnwinding(engine, unw)
					os.Exit(1)
				} else {
					panic(err)
				}
			}
		}
	case "help":
		fmt.Print(HelpText)
	case "repl":
		enterRepl(engine, comp)
	default:
		panic("Bad verb '" + args[0] + "', try `sint help`")
	}
}

type errorReporter interface {
	WriteString(s string) (int, error)
}

func reportUnwinding(engine *core.Scheme, unw *core.UnwindPkg) {
	if engine.UnwindReporter == nil {
		panic("UNHANDLED UNWINDING WITHOUT AN INSTALLED UNWIND REPORTER")
	}
	engine.UnwindReporter(engine, unw)
}

func enterRepl(engine *core.Scheme, comp *compiler.Compiler) {
	stdin, stdout, stderr := runtime.StandardInitialization(engine)
	nextResultId := 1
	for {
		stdout.WriteString("> ")
		form, rdrErr := runtime.Read(engine, stdin)
		if rdrErr != nil {
			stderr.WriteString(rdrErr.Error() + "\n")
			continue
		}
		if form == engine.EofVal {
			stdout.WriteRune('\n')
			break
		}
		prog, progErr := comp.CompileToplevel(form)
		if progErr != nil {
			stderr.WriteString(progErr.Error() + "\n")
			continue
		}
		results, unw := engine.EvalToplevel(prog)
		if unw != nil {
			reportUnwinding(engine, unw.(*core.UnwindPkg))
			continue
		}
		for _, result := range results {
			if result != engine.UnspecifiedVal {
				rName := nextResultId
				nextResultId++
				name := fmt.Sprintf("$%d", rName)
				fmt.Printf("%s = ", name)
				runtime.Write(result, false, stdout)
				stdout.WriteRune('\n')
				engine.DefineToplevel(name, result)
			}
		}
	}
}

func evalExpr(engine *core.Scheme, comp *compiler.Compiler, expr string) error {
	_, stdout, _ := runtime.StandardInitialization(engine)
	sourceReader := bufio.NewReader(strings.NewReader(expr))
	v, rdrErr := runtime.Read(engine, sourceReader)
	if rdrErr != nil {
		return rdrErr
	}
	prog, progErr := comp.CompileToplevel(v)
	if progErr != nil {
		return progErr
	}
	results, unw := engine.EvalToplevel(prog)
	if unw != nil {
		return unw.(*core.UnwindPkg)
	}
	for _, r := range results {
		if r != engine.UnspecifiedVal {
			runtime.Write(r, false, stdout)
			stdout.WriteRune('\n')
			stdout.Flush()
		}
	}
	return nil
}

func loadFile(engine *core.Scheme, comp *compiler.Compiler, fn string) error {
	_, stdout, _ := runtime.StandardInitialization(engine)
	input, inErr := os.Open(fn)
	if inErr != nil {
		panic(inErr)
	}
	sourceReader := bufio.NewReader(input)
	for {
		v, rdrErr := runtime.Read(engine, sourceReader)
		if rdrErr != nil {
			return rdrErr
		}
		if v == engine.EofVal {
			break
		}
		prog, progErr := comp.CompileToplevel(v)
		if progErr != nil {
			return progErr
		}
		results, unw := engine.EvalToplevel(prog)
		if unw != nil {
			return unw.(*core.UnwindPkg)
		}
		for _, r := range results {
			if r != engine.UnspecifiedVal {
				runtime.Write(r, false, stdout)
				stdout.WriteRune('\n')
			}
		}
	}
	input.Close()
	return nil
}

func compileFile(engine *core.Scheme, comp *compiler.Compiler, fn string) error {
	if strings.LastIndex(fn, ".sch") != len(fn)-4 {
		return compiler.NewCompilerError("Input file for 'compile' must have type '.sch': " + fn)
	}
	withoutExt := fn[:len(fn)-4]
	ix := strings.LastIndexAny(withoutExt, "/\\")
	moduleName := withoutExt
	if ix != -1 {
		moduleName = moduleName[ix+1:]
	}
	if len(moduleName) == 0 {
		return compiler.NewCompilerError("Input file name is empty after stripping suffix: " + fn)
	}
	moduleName = strings.ToUpper(moduleName[0:1]) + strings.ToLower(moduleName[1:])
	tmpFn := withoutExt + ".tmp"
	outFn := withoutExt + ".go"
	input, inErr := os.Open(fn)
	if inErr != nil {
		return inErr
	}
	// TODO: Use proper tempfiles
	tmp, tmpErr := os.Create(tmpFn)
	if tmpErr != nil {
		return tmpErr
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
			return rdrErr
		}
		if v == engine.EofVal {
			break
		}
		prog, progErr := comp.CompileToplevel(v)
		if progErr != nil {
			return progErr
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
	return nil
}
