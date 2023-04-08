package assembler

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	m "github.com/petabite/gomachine/internal/machine"
)

type OperationKey struct {
	operation         string
	registerOperation bool
}

var operationMap = map[OperationKey]int{
	{"mov", false}: m.OpMovConst,
	{"mov", true}:  m.OpMovRegister,
	{"add", false}: m.OpAddConst,
	{"add", true}:  m.OpAddRegister,
	{"cmp", false}: m.OpCmpConst,
	{"j", false}:   m.OpJmp,
	{"jne", false}: m.OpJmpNe,
	{"jlt", false}: m.OpJmpLt,
	{"jgt", false}: m.OpJmpGt,
	{"and", false}: m.OpAndConst,
	{"and", true}:  m.OpAndRegister,
	{"or", false}:  m.OpOrConst,
	{"or", true}:   m.OpOrRegister,
	{"not", false}: m.OpNotConst,
	{"not", true}:  m.OpNotRegister,
	{"xor", false}: m.OpXorConst,
	{"xor", true}:  m.OpXorRegister,
}

func Assemble(file string) ([]m.Instruction, error) {
	// in: assembly
	lines, err := readFileLines(file)
	if err != nil {
		return nil, err
	}

	// assembly line to assemble the assembly
	subroutine := []m.Instruction{}
	for _, line := range lines {
		tokens := tokenizeLine(line)
		if err != nil {
			return nil, errors.New("Error tokenizing file: " + err.Error())
		}
		instruction, err := tokensToInstruction(tokens)
		if err != nil {
			return nil, errors.New("Error converting tokens to instruction: " + err.Error())
		}
		subroutine = append(subroutine, instruction)
	}

	// out: "machine code"
	return subroutine, nil
}

func tokensToInstruction(t Tokens) (m.Instruction, error) {
	operationKey := OperationKey{}
	var arguments []uint64

	// parse operation
	operationKey.operation = t.keyword

	// parse arguments
	for index, stringArg := range t.arguments {
		if strings.HasPrefix(stringArg, "r") {
			if index == len(t.arguments)-1 {
				operationKey.registerOperation = true
			}
			// strip the 'r' prefix
			stringArg = stringArg[1:]
		}
		arg, err := strconv.ParseUint(stringArg, 10, 64)
		if err != nil {
			return m.Instruction{}, err
		}
		arguments = append(arguments, arg)
	}

	operation, exists := operationMap[operationKey]
	if !exists {
		return m.Instruction{}, errors.New("Invalid instruction: " + t.keyword)
	}

	return *m.NewInstruction(operation, arguments...), nil
}

func tokenizeLine(line string) Tokens {
	stringTokens := strings.Split(strings.ToLower(line), " ")
	return Tokens{keyword: stringTokens[0], arguments: stringTokens[1:]}
}

func readFileLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
