package bytecode

import "fmt"

// This package has all of the bytecode we require for compiling the language.
// We will use Stack Compiling, Not using physical register as we don't have access to them

// Instructions are slice of Byte.
type Instructions []byte

// Opcode is OneByte instruction, We are specifically defining this as OpCode as byte cannot be used
// Makes easier to debug and attach additional functions
type Opcode byte

// Defining OpConstants, As we are going to generate the codes on fly.
const (
	OpConstant Opcode = iota
)

// Opcode definations, We will use this to create further instructions for CPU and debug
type Defination struct {
	Name          string
	OperandWidths []int
}

// This is map which stores the defination for opcode.
var definations = map[Opcode]*Defination{
	OpConstant: {"OpConstant", []int{2}},
}

// Lookup returns the defination pointer or error if the opcode does not exist
// This can be when the opcode is not in instruction set
func Lookup(op byte) (*Defination, error) {
	def, ok := definations[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d is undefined", op)
	}
	return def, nil
}
