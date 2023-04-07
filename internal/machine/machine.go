package machine

import "fmt"

func NewGoMachine() *GoMachine {
	return &GoMachine{}
}

type GoMachine struct {
	registerFile [10]uint64
	pc           int
	cond         int64
}

func (m *GoMachine) Run(subroutine []Instruction) {
	for m.pc < len(subroutine) {
		m.Execute(subroutine[m.pc])
		m.pc++
	}
	fmt.Println(m.registerFile[0])
}

func (m *GoMachine) Execute(i Instruction) {
	// fmt.Println("PC:", m.pc)
	// fmt.Println("Register File:", m.registerFile)
	// fmt.Println("Condition Register:", m.cond)
	switch i.operation {
	case OpMovConst:
		m.Mov(i.dest, i.operand)
	case OpAddConst:
		m.Add(i.dest, i.source, i.operand)
	case OpAddRegister:
		m.Add(i.dest, i.source, m.registerFile[i.operand])
	case OpCmpConst:
		m.Cmp(i.dest, i.operand)
	case OpJmpNe:
		if m.cond != 0 {
			m.Jump(int(i.dest))
		}
	case OpJmp:
		m.Jump(int(i.dest))
	}
}

func (m *GoMachine) Mov(dest, operand uint64) {
	m.registerFile[dest] = operand
}

func (m *GoMachine) Add(dest, src, operand uint64) {
	m.registerFile[dest] = m.registerFile[src] + operand
}

func (m *GoMachine) Cmp(dest, operand uint64) {
	m.cond = int64(m.registerFile[dest] - operand)
}

func (m *GoMachine) Jump(dest int) {
	m.pc = dest - 1
}
