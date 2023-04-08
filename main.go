package main

import (
	"fmt"
	"os"

	asm "github.com/petabite/gomachine/internal/assembler"
	m "github.com/petabite/gomachine/internal/machine"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No source file provided!")
		return
	}
	source := args[1]
	machine := m.NewGoMachine()
	subroutine, err := asm.Assemble(source)
	if err != nil {
		panic(err)
	}
	machine.Run(subroutine)
}
