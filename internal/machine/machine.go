package machine

import "fmt"

type GoMachine struct {
	registerFile        [10]uint64
	pc                  int
	cond                int64
	dynamicInstructions uint64
}

func Run(subroutine []Instruction) {
	machine := &GoMachine{}
	for machine.pc < len(subroutine) {
		machine.execute(subroutine[machine.pc])
		machine.dynamicInstructions++
		machine.pc++
	}
	fmt.Println(machine.registerFile[0])
}

func (m *GoMachine) execute(i Instruction) {
	// fmt.Println("PC:", m.pc)
	// fmt.Println("Register File:", m.registerFile)
	// fmt.Println("Condition Register:", m.cond)
	switch i.operation {
	case OpMovConst:
		m.move(i.arguments[0], i.arguments[1])
	case OpMovRegister:
		m.move(i.arguments[0], m.registerFile[i.arguments[1]])
	case OpAddConst:
		m.add(i.arguments[0], i.arguments[1], i.arguments[2])
	case OpAddRegister:
		m.add(i.arguments[0], i.arguments[1], m.registerFile[i.arguments[2]])
	case OpDecRegister:
		m.add(i.arguments[0], i.arguments[0], ^uint64(0))
	case OpIncRegister:
		m.add(i.arguments[0], i.arguments[0], 1)
	case OpCmpConst:
		m.compare(i.arguments[0], i.arguments[1])
	case OpJmp:
		m.jump(int(i.arguments[0]))
	case OpJmpNe:
		if m.cond != 0 {
			m.jump(int(i.arguments[0]))
		}
	case OpJmpEq:
		if m.cond == 0 {
			m.jump(int(i.arguments[0]))
		}
	case OpJmpLt:
		if m.cond < 0 {
			m.jump(int(i.arguments[0]))
		}
	case OpJmpGt:
		if m.cond > 0 {
			m.jump(int(i.arguments[0]))
		}
	case OpAndConst:
		m.and(i.arguments[0], i.arguments[1], i.arguments[2])
	case OpAndRegister:
		m.and(i.arguments[0], i.arguments[1], m.registerFile[i.arguments[2]])
	case OpOrConst:
		m.or(i.arguments[0], i.arguments[1], i.arguments[2])
	case OpOrRegister:
		m.or(i.arguments[0], i.arguments[1], m.registerFile[i.arguments[2]])
	case OpNotConst:
		m.not(i.arguments[0], i.arguments[1])
	case OpNotRegister:
		m.not(i.arguments[0], m.registerFile[i.arguments[1]])
	case OpXorConst:
		m.xor(i.arguments[0], i.arguments[1], i.arguments[2])
	case OpXorRegister:
		m.xor(i.arguments[0], i.arguments[1], m.registerFile[i.arguments[2]])
	}
}

func (m *GoMachine) move(dest, value uint64) {
	m.registerFile[dest] = value
}

func (m *GoMachine) add(dest, src, value uint64) {
	m.registerFile[dest] = m.registerFile[src] + value
}

func (m *GoMachine) compare(dest, value uint64) {
	m.cond = int64(m.registerFile[dest] - value)
}

func (m *GoMachine) jump(dest int) {
	m.pc = dest - 1
}

func (m *GoMachine) and(dest, src, value uint64) {
	m.registerFile[dest] = m.registerFile[src] & value
}

func (m *GoMachine) or(dest, src, value uint64) {
	m.registerFile[dest] = m.registerFile[src] | value
}

func (m *GoMachine) not(dest, value uint64) {
	m.registerFile[dest] = ^value
}

func (m *GoMachine) xor(dest, src, value uint64) {
	m.registerFile[dest] = m.registerFile[src] ^ value
}
