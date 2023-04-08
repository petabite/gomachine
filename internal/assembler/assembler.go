package assembler

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	m "github.com/petabite/gomachine/internal/machine"
)

func Assemble(file string) ([]m.Instruction, error) {
	subroutine := []m.Instruction{}
	lines, err := readFileLines(file)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		tokens, err := tokenizeLine(line)
		if err != nil {
			return nil, errors.New("Error tokenizing file: " + err.Error())
		}
		instruction, err := tokensToInstruction(tokens)
		if err != nil {
			return nil, errors.New("Error converting tokens to instruction: " + err.Error())
		}
		subroutine = append(subroutine, instruction)
	}

	return subroutine, nil
}

func tokensToInstruction(t Tokens) (m.Instruction, error) {
	var instruction m.Instruction
	var err error
	switch t.operation {
	case "mov":
		operation := m.OpMovConst
		if t.hasRegisterOperand {
			operation = m.OpMovRegister
		}
		instruction = *m.NewImmediateInstruction(operation, t.arguments[0], t.arguments[1])
	case "add":
		operation := m.OpAddConst
		if t.hasRegisterOperand {
			operation = m.OpAddRegister
		}
		instruction = *m.NewDataInstruction(operation, t.arguments[0], t.arguments[1], t.arguments[2])
	default:
		err = errors.New("Invalid instruction: " + t.operation)
	}

	return instruction, err
}

func tokenizeLine(line string) (Tokens, error) {
	stringTokens := strings.Split(strings.ToLower(line), " ")
	// parse operation
	tokens := Tokens{operation: stringTokens[0]}
	// parse arguments
	for index, stringArg := range stringTokens[1:] {
		if strings.HasPrefix(stringArg, "r") {
			if index == len(stringTokens)-1 {
				tokens.hasRegisterOperand = true
			}
			// strip the 'r' prefix
			stringArg = stringArg[1:]
		}
		arg, err := strconv.ParseUint(stringArg, 10, 64)
		if err != nil {
			return Tokens{}, err
		}
		tokens.arguments = append(tokens.arguments, arg)
	}
	return tokens, nil
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
