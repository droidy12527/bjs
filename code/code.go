// This package has code for generating definations for Code
// this has operand and opcode, In this package functions exist which can generate
// instructions
package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

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
	OpAdd
	OpPop
	OpSub
	OpMul
	OpDiv
)

// Opcode definations, We will use this to create further instructions for CPU and debug
type Defination struct {
	Name          string
	OperandWidths []int
}

// This is map which stores the defination for opcode.
var definations = map[Opcode]*Defination{
	OpConstant: {"OpConstant", []int{2}},
	OpAdd:      {"OpAdd", []int{}},
	OpSub:      {"OpSub", []int{}},
	OpMul:      {"OpMul", []int{}},
	OpDiv:      {"OpDiv", []int{}},
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

// Creates an slice of byte which has opcode and operand and returns it back
// Thsi function posses risk of generating empty bytecode instructions
func Make(op Opcode, operands ...int) []byte {
	def, ok := definations[op]
	if !ok {
		return []byte{}
	}
	// Define instruction lenght for the CPU
	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	// The first instruction byte should always be OpCode
	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	// Define OffSet and start pushing instructions
	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}

// Gets the defination and instructions and returns back operands and offsets
func ReadOperands(def *Defination, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0
	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}
		offset += width
	}
	return operands, offset
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

// Gets the instruction and returns back string of the instruction
func (ins Instructions) String() string {
	var out bytes.Buffer
	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}
		operands, read := ReadOperands(def, ins[i+1:])
		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))
		i += 1 + read
	}
	return out.String()
}

func (ins Instructions) fmtInstruction(def *Defination, operands []int) string {
	operandCount := len(def.OperandWidths)
	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n", len(operands), operandCount)
	}
	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}
	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}
