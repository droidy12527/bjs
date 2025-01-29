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

// Setting global values of True and False as they are immutable and do not change
// Defining them everytime gains memory space and has to gc it again and again
var True = &object.Boolean{Value: true}
var False = &object.Boolean{Value: false}

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
		case code.OpTrue:
			err := vm.push(True)
			if err != nil {
				return err
			}
		case code.OpFalse:
			err := vm.push(False)
			if err != nil {
				return err
			}
		case code.OpAdd, code.OpDiv, code.OpMul, code.OpSub:
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return err
			}
		case code.OpPop:
			vm.pop()
		}
	}
	return nil
}

// Does basic checks on the opcode and values and returns back if there is any error
func (vm *VirtualMachine) executeBinaryOperation(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	if right.Type() == constants.INTEGER_OBJECT && left.Type() == constants.INTEGER_OBJECT {
		return vm.executeBinaryIntegerOperation(op, left, right)
	}
	return fmt.Errorf("unsupported types %s %s", left.Type(), right.Type())
}

// Does the operations on left and right operator and then pushed them on the VM stack and returns back error
// if found
func (vm *VirtualMachine) executeBinaryIntegerOperation(op code.Opcode, left, right object.Object) error {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	var result int64
	switch op {
	case code.OpAdd:
		result = leftVal + rightVal
	case code.OpSub:
		result = leftVal - rightVal
	case code.OpDiv:
		result = leftVal / rightVal
	case code.OpMul:
		result = leftVal * rightVal
	default:
		return fmt.Errorf("unknown integer operator : %d", op)
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

// Returns back the last popped element from the stack
func (vm *VirtualMachine) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}
