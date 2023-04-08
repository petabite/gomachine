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
		m.Move(i.arguments[0], i.arguments[1])
	case OpMovRegister:
		m.Move(i.arguments[0], m.registerFile[i.arguments[1]])
	case OpAddConst:
		m.Add(i.arguments[0], i.arguments[1], i.arguments[2])
	case OpAddRegister:
		m.Add(i.arguments[0], i.arguments[1], m.registerFile[i.arguments[2]])
	case OpDecRegister:
		m.Add(i.arguments[0], i.arguments[0], ^uint64(0))
	case OpIncRegister:
		m.Add(i.arguments[0], i.arguments[0], 1)
	case OpCmpConst:
		m.Compare(i.arguments[0], i.arguments[1])
	case OpJmp:
		m.Jump(int(i.arguments[0]))
	case OpJmpNe:
		if m.cond != 0 {
			m.Jump(int(i.arguments[0]))
		}
	case OpJmpLt:
		if m.cond < 0 {
			m.Jump(int(i.arguments[0]))
		}
	case OpJmpGt:
		if m.cond > 0 {
			m.Jump(int(i.arguments[0]))
		}
	case OpAndConst:
		m.And(i.arguments[0], i.arguments[1], i.arguments[2])
	case OpAndRegister:
		m.And(i.arguments[0], i.arguments[1], m.registerFile[i.arguments[2]])
	case OpOrConst:
		m.Or(i.arguments[0], i.arguments[1], i.arguments[2])
	case OpOrRegister:
		m.Or(i.arguments[0], i.arguments[1], m.registerFile[i.arguments[2]])
	case OpNotConst:
		m.Not(i.arguments[0], i.arguments[1])
	case OpNotRegister:
		m.Not(i.arguments[0], m.registerFile[i.arguments[1]])
	case OpXorConst:
		m.Xor(i.arguments[0], i.arguments[1], i.arguments[2])
	case OpXorRegister:
		m.Xor(i.arguments[0], i.arguments[1], m.registerFile[i.arguments[2]])
	}
}

func (m *GoMachine) Move(dest, value uint64) {
	m.registerFile[dest] = value
}

func (m *GoMachine) Add(dest, src, value uint64) {
	m.registerFile[dest] = m.registerFile[src] + value
}

func (m *GoMachine) Compare(dest, value uint64) {
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

func (m *GoMachine) Not(dest, value uint64) {
	m.registerFile[dest] = ^value
}

func (m *GoMachine) Xor(dest, src, value uint64) {
	m.registerFile[dest] = m.registerFile[src] ^ value
}
