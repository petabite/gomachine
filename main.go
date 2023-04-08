package main

import (
	m "github.com/petabite/gomachine/internal/machine"
	asm "github.com/petabite/gomachine/internal/assembler"
)

func main() {
	machine := m.NewGoMachine()
	subroutine, err := asm.Assemble("examples/loop.gm")
	if err != nil {
		panic(err)
	}
	machine.Run(subroutine)
}
