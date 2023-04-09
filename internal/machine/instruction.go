package machine

const (
	OpMovConst = iota
	OpMovRegister
	OpAddConst
	OpAddRegister
	OpDecRegister
	OpIncRegister
	OpCmpConst
	OpJmp
	OpJmpNe
	OpJmpEq
	OpJmpLt
	OpJmpGt
	OpAndConst
	OpAndRegister
	OpOrConst
	OpOrRegister
	OpNotConst
	OpNotRegister
	OpXorConst
	OpXorRegister
)

type Instruction struct {
	operation int
	arguments []uint64
}

func NewInstruction(operation int, arguments ...uint64) *Instruction {
	return &Instruction{operation: operation, arguments: arguments}
}
