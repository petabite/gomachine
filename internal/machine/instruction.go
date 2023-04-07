package machine

const (
	OpMovConst = iota
	OpMovRegister
	OpAddConst
	OpAddRegister
	OpCmpConst
	OpJmp
	OpJmpNe
	// OpAnd
	// OpOr
	// OpNot
	// OpJz
	// OpNop
)

type Instruction struct {
	operation int
	source    uint64
	dest      uint64
	operand   uint64
}

func NewJumpInstruction(operation int, dest uint64) *Instruction {
	return &Instruction{operation: operation, source: 0, dest: dest, operand: 0}
}

func NewImmediateInstruction(operation int, dest, operand uint64) *Instruction {
	return &Instruction{operation: operation, source: 0, dest: dest, operand: operand}
}

func NewDataInstruction(operation int, dest, source, operand uint64) *Instruction {
	return &Instruction{operation: operation, source: source, dest: dest, operand: operand}
}
