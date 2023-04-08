package main

import (
	m "github.com/petabite/gomachine/internal/machine"
)

func main() {
	machine := m.NewGoMachine()
	subroutine := []m.Instruction{
		*m.NewImmediateInstruction(m.OpMovConst, 1, 3),
		*m.NewImmediateInstruction(m.OpMovConst, 2, 2),
		*m.NewDataInstruction(m.OpAddRegister, 3, 1, 2),

		*m.NewDataInstruction(m.OpAddConst, 4, 4, 1),
		*m.NewDataInstruction(m.OpAddRegister, 0, 0, 3),
		*m.NewImmediateInstruction(m.OpCmpConst, 4, 10),
		*m.NewLiteralInstruction(m.OpJmpNe, 3),
	}
	machine.Run(subroutine)
}
