package virtualmachine

import (
	"compiler/code"
	"compiler/compiler"
	"compiler/object"
	"fmt"
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

// Returs back the object located on the stacktop
func (vm *VirtualMachine) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

// Returns back error and runs the program in Fetch, Decode, Execute cycle
func (vm *VirtualMachine) Run() error {
	// Setup instruction pointer and start going through the pointer
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])
		// Get the opcode and then start decoding in execute cycle
		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Pushes the object to stack of Virtual machine and increments the stackpointer
func (vm *VirtualMachine) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.sp] = o
	vm.sp++
	return nil
}
