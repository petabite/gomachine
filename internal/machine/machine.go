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
	case OpMovRegister:
		m.Mov(i.dest, m.registerFile[i.operand])
	case OpAddConst:
		m.Add(i.dest, i.source, i.operand)
	case OpAddRegister:
		m.Add(i.dest, i.source, m.registerFile[i.operand])
	case OpCmpConst:
		m.Cmp(i.dest, i.operand)
	case OpJmp:
		m.Jump(int(i.dest))
	case OpJmpNe:
		if m.cond != 0 {
			m.Jump(int(i.dest))
		}
	case OpJmpLt:
		if m.cond < 0 {
			m.Jump(int(i.dest))
		}
	case OpJmpGt:
		if m.cond > 0 {
			m.Jump(int(i.dest))
		}
	case OpAndConst:
		m.And(i.dest, i.source, i.operand)
	case OpAndRegister:
		m.And(i.dest, i.source, m.registerFile[i.operand])
	case OpOrConst:
		m.Or(i.dest, i.source, i.operand)
	case OpOrRegister:
		m.Or(i.dest, i.source, m.registerFile[i.operand])
	case OpNotRegister:
		m.Not(int(i.dest))
	case OpXorConst:
		m.Xor(i.dest, i.source, i.operand)
	case OpXorRegister:
		m.Xor(i.dest, i.source, m.registerFile[i.operand])
	}
}

func (m *GoMachine) Mov(dest, value uint64) {
	m.registerFile[dest] = value
}

func (m *GoMachine) Add(dest, src, value uint64) {
	m.registerFile[dest] = m.registerFile[src] + value
}

func (m *GoMachine) Cmp(dest, value uint64) {
	m.cond = int64(m.registerFile[dest] - value)
}

func (m *GoMachine) Jump(dest int) {
	m.pc = dest - 1
}

func (m *GoMachine) And(dest, src, value uint64) {
	m.registerFile[dest] = m.registerFile[src] & value
}

func (m *GoMachine) Or(dest, src, value uint64) {
	m.registerFile[dest] = m.registerFile[src] | value
}

func (m *GoMachine) Not(dest int) {
	m.registerFile[dest] = ^m.registerFile[dest]
}

func (m *GoMachine) Xor(dest, src, value uint64) {
	m.registerFile[dest] = m.registerFile[src] ^ value
}
