package main

import (
	"bufio"
	"os"
	"sint/core"
	"sint/runtime"
)

func main() {
	c := core.NewScheme()
	runtime.InitPrimitives(c)
	runtime.InitCompiled(c)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for {
		writer.WriteString("> ")
		writer.Flush()
		v := runtime.Read(c, reader)
		if v == c.EofVal {
			break
		}
		runtime.Write(v, writer)
		writer.WriteRune('\n')
		writer.Flush()
	}
}
