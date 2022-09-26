// "sint" - an extended subset of R7RS scheme, embedded in Go, with many Go facilities.
//
// Command line parsing and command implementation.

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/pprof"
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

  sint [-cpuprofile|-memprofile] load filename.sch ...
    Load filename.sch: read its expressions, evaluate them in order and
    print their results, and exit after the last expression of the last file.
    A profile is written to sint.cprof or sint.mprof if the options are
    present.  (CPU profiling excludes the initialization of the built-in
    libraries, but memory profiling needs to include that.)

  sint compile filename.sch ...
    Compile each filename.sch into filename.go and exit.  The output will have
    a function initFilename() that takes a *Scheme and evaluates the
    expressions and definitions of filename.sch in order in that runtime.

  sint help
    Print help (this text)
`

func main() {
	cpuprofile := false
	memprofile := false

	args := os.Args[1:]

	if len(args) > 0 && args[0] == "-cpuprofile" {
		cpuprofile = true
	} else if len(args) > 0 && args[0] == "-memprofile" {
		memprofile = true
	}
	if cpuprofile || memprofile {
		args = args[1:]
	}

	engine := core.NewScheme(nil, nil)
	comp := compiler.NewCompiler(engine.Shared)

	if len(args) == 0 {
		args = []string{"repl"}
	}

	if args[0] != "load" && (cpuprofile || memprofile) {
		panic("Profiling only with the `load` verb")
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
		_, stdout, _ := runtime.StandardInitialization(engine)
		for _, ex := range args[1:] {
			err := evalExpr(engine, comp, ex, stdout)
			if err != nil {
				if unw, ok := err.(*core.UnwindPkg); ok {
					reportUnwinding(engine, unw)
					os.Exit(1)
				} else {
					panic(err)
				}
			}
		}
		stdout.Flush()
	case "load":
		if len(args) < 2 {
			panic("Bad 'load' command, at least one file name argument required")
		}
		_, stdout, _ := runtime.StandardInitialization(engine)
		if cpuprofile {
			f, err := os.Create("sint.cprof")
			if err != nil {
				panic(err)
			}
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}
		for _, fn := range args[1:] {
			err := loadFile(engine, comp, fn, stdout)
			if err != nil {
				if unw, ok := err.(*core.UnwindPkg); ok {
					reportUnwinding(engine, unw)
					os.Exit(1)
				} else {
					panic(err)
				}
			}
		}
		stdout.Flush()
	case "help":
		fmt.Print(HelpText)
	case "repl":
		stdin, stdout, stderr := runtime.StandardInitialization(engine)
		enterRepl(engine, comp, stdin, stdout, stderr)
	default:
		panic("Bad verb '" + args[0] + "', try `sint help`")
	}
	if memprofile {
		f, err := os.Create("sint.mprof")
		if err != nil {
			panic(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}
}

func reportUnwinding(engine *core.Scheme, unw *core.UnwindPkg) {
	if engine.UnwindReporter == nil {
		panic("UNHANDLED UNWINDING WITHOUT AN INSTALLED UNWIND REPORTER")
	}
	engine.UnwindReporter(engine, unw)
}

func enterRepl(engine *core.Scheme, comp *compiler.Compiler,
	stdin runtime.InputStream, stdout runtime.OutputStream, stderr runtime.OutputStream) {
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

func evalExpr(engine *core.Scheme, comp *compiler.Compiler, expr string, stdout runtime.OutputStream) error {
	sourceReader := bufio.NewReader(strings.NewReader(expr))
	form, rdrErr := runtime.Read(engine, sourceReader)
	if rdrErr != nil {
		return rdrErr
	}
	prog, progErr := comp.CompileToplevel(form)
	if progErr != nil {
		return progErr
	}
	results, unw := engine.EvalToplevel(prog)
	if unw != nil {
		return unw.(*core.UnwindPkg)
	}
	for _, result := range results {
		if result != engine.UnspecifiedVal {
			runtime.Write(result, false, stdout)
			stdout.WriteRune('\n')
		}
	}
	return nil
}

func loadFile(engine *core.Scheme, comp *compiler.Compiler, fn string, stdout runtime.OutputStream) error {
	input, inErr := os.Open(fn)
	if inErr != nil {
		panic(inErr)
	}
	sourceReader := bufio.NewReader(input)
	for {
		form, rdrErr := runtime.Read(engine, sourceReader)
		if rdrErr != nil {
			return rdrErr
		}
		if form == engine.EofVal {
			break
		}
		prog, progErr := comp.CompileToplevel(form)
		if progErr != nil {
			return progErr
		}
		results, unw := engine.EvalToplevel(prog)
		if unw != nil {
			return unw.(*core.UnwindPkg)
		}
		for _, result := range results {
			if result != engine.UnspecifiedVal {
				runtime.Write(result, false, stdout)
				stdout.WriteRune('\n')
			}
		}
	}
	input.Close()
	return nil
}

func compileFile(engine *core.Scheme, comp *compiler.Compiler, fn string) error {
	// TODO: Are there path name utilities that could be brought to bear here?
	// Note the spec explicitly prohibits \ from being interpreted as a path separator,
	// only / is a valid path separator, see fs.ValidPath.
	if strings.LastIndex(fn, ".sch") != len(fn)-4 {
		return compiler.NewCompilerError("Input file for 'compile' must have type '.sch': " + fn)
	}
	withoutExt := fn[:len(fn)-4]
	ix := strings.LastIndexAny(withoutExt, "/\\")
	// The module name is the base file name without any special characters
	moduleName := withoutExt
	if ix != -1 {
		moduleName = moduleName[ix+1:]
	}
	newModuleName := ""
	for _, c := range moduleName {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' || (c >= '0' && c <= '9') {
			newModuleName += string(c)
		}
	}
	if len(newModuleName) == 0 {
		return compiler.NewCompilerError("Module name would be empty: " + fn)
	}
	moduleName = strings.ToUpper(newModuleName[0:1]) + strings.ToLower(newModuleName[1:])
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
// Make sure the imports are used, or the Go compiler barfs.
var _ Val = big.NewInt(0)
func init%s(c *Scheme) {
`, fn, moduleName)
	id := 1
	for {
		form, rdrErr := runtime.Read(engine, reader)
		if rdrErr != nil {
			return rdrErr
		}
		if form == engine.EofVal {
			break
		}
		prog, progErr := comp.CompileToplevel(form)
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
