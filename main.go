package main

import (
	"fmt"
	"os"

	"github.com/petabite/gomachine/internal/assembler"
	"github.com/petabite/gomachine/internal/machine"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No source file provided!")
		return
	}
	source := args[1]
	subroutine, err := assembler.Assemble(source)
	if err != nil {
		panic(err)
	}
	machine.Run(subroutine)
}
