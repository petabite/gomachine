package machine

const (
	OpMovConst = iota
	OpMovRegister
	OpAddConst
	OpAddRegister
	OpCmpConst
	OpJmp
	OpJmpNe
	OpJmpLt
	OpJmpGt
	OpAndConst
	OpAndRegister
	OpOrConst
	OpOrRegister
	OpNotRegister
	OpXorConst
	OpXorRegister
)

type Instruction struct {
	operation             int
	source, dest, operand uint64
}

func NewLiteralInstruction(operation int, dest uint64) *Instruction {
	return &Instruction{operation: operation, source: 0, dest: dest, operand: 0}
}

func NewImmediateInstruction(operation int, dest, operand uint64) *Instruction {
	return &Instruction{operation: operation, source: 0, dest: dest, operand: operand}
}

func NewDataInstruction(operation int, dest, source, operand uint64) *Instruction {
	return &Instruction{operation: operation, source: source, dest: dest, operand: operand}
}
