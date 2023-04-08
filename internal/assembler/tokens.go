package assembler

type Tokens struct {
	operation          string
	arguments          []uint64
	hasRegisterOperand bool
}
