package virtualmachine

import (
	"compiler/code"
	"compiler/compiler"
	"compiler/constants"
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
		case code.OpAdd, code.OpDiv, code.OpMul, code.OpSub:
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return err
			}
		// Just pop the element off stack
		case code.OpPop:
			vm.pop()
		}
	}
	return nil
}

func (vm *VirtualMachine) executeBinaryOperation(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()
	if left.Type() == constants.INTEGER_OBJECT && right.Type() == constants.INTEGER_OBJECT {
		return vm.executeBinaryIntegerOperation(op, left, right)
	}
	return fmt.Errorf("types are unsupported %s, %s", left.Type(), right.Type())
}

func (vm *VirtualMachine) executeBinaryIntegerOperation(op code.Opcode, left object.Object, right object.Object) error {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value
	var result int64
	switch op {
	case code.OpAdd:
		result = leftValue + rightValue
	case code.OpDiv:
		result = leftValue / rightValue
	case code.OpMul:
		result = leftValue * rightValue
	case code.OpSub:
		result = leftValue - rightValue
	default:
		return fmt.Errorf("unsupported operation code %d", op)
	}
	return vm.push(&object.Integer{Value: result})
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

// Pops from the stack of Virtual machine and returns back object
func (vm *VirtualMachine) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}

// Returns the last popped element in stack
func (vm *VirtualMachine) LastPoppedStackElement() object.Object {
	return vm.stack[vm.sp]
}
