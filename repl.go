package main

import (
	"bufio"
	"os"
	"sint/compiler"
	"sint/core"
	"sint/runtime"
)

func main() {
	engine := core.NewScheme()
	runtime.InitPrimitives(engine)
	runtime.InitCompiled(engine)
	comp := compiler.NewCompiler(engine)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for {
		writer.WriteString("> ")
		writer.Flush()
		v := runtime.Read(engine, reader)
		// TODO: This test for EOF does not work, interestingly, but where is the error?
		if v == engine.EofVal {
			break
		}
		// TODO: Recover from compilation error
		prog := comp.CompileToplevel(v)
		writer.WriteString(prog.String() + "\n")
		writer.Flush()
		// TODO: Recover from runtime error
		result := engine.EvalToplevel(prog)
		// TODO: Maybe not write if unspecified?
		runtime.Write(result, writer)
		writer.WriteRune('\n')
		writer.Flush()
	}
}
