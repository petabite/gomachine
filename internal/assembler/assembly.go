package assembler

import (
	m "github.com/petabite/gomachine/internal/machine"
)

type Assembly struct {
	source             string
	sourceTokens       []Tokens
	labels             map[string]int
	subroutine         []m.Instruction
	staticInstructions int
}

type Tokens struct {
	keyword   string
	arguments []string
}

type OperationKey struct {
	operation         string
	registerOperation bool
}
