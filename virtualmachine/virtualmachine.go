package virtualmachine

import (
	"compiler/code"
	"compiler/compiler"
	"compiler/object"
)

// Defining stacksize to also check with stack overflow
// Limited stack size has been used, for in depth recursive operations this stack size can be increased
const StackSize = 2048

type VirtualMachine struct {
	constants    []object.Object
	instructions code.Instructions
	stack        []object.Object // Virtual machine stack
	sp           int             // StackPointer always points to the top of the stack
}

// Creates new virtual machine and returns back for execution
func New(bytecode *compiler.ByteCode) *VirtualMachine {
	// Create a virtual machine using the basic params given in the bytecode
	// Set the stack size to the one defined above.
	return &VirtualMachine{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,
		stack:        make([]object.Object, StackSize),
		sp:           0,
	}
}
