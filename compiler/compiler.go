// This package contains compiler for compiling the instructions
package compiler

import (
	"compiler/ast"
	"compiler/code"
	"compiler/object"
	"fmt"
)

// Compiler is a struct that contains bytecode instructions and constants
type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

// This is higher level abstraction for code, This is bytecode which has constant
// This has
type ByteCode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

// Creates and returns new compiler to compile the code
// This is also used for testing
func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

// Compiles the code and returns back if there is error
func (c *Compiler) Compile(node ast.Node) error {
	// Get the node type
	switch node := node.(type) {
	// Check if it is AST Program node if so traverse the node statements
	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}
	// If expression Statement exist then check the statement and traverse
	// recursively on expression
	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
	// Get the left and right node for infix expression and compile them
	case *ast.InfixExpression:
		err := c.Compile(node.Left)
		if err != nil {
			return err
		}
		err = c.Compile(node.Right)
		if err != nil {
			return err
		}
		// Check which operator is there, According to that emit to instruction stack
		switch node.Operator {
		case "+":
			c.emit(code.OpAdd)
		default:
			return fmt.Errorf("unknown operator founf %s", node.Operator)
		}

	// Get the node value and assign to integer
	// Once that is done push it to constant stack to evaluate further
	case *ast.IntegerLiteral:
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(integer))
	}
	return nil
}

// Returns bytecode struct, This struct has bytecode
// Code has functions that generates instructions and constants
func (c *Compiler) ByteCode() *ByteCode {
	return &ByteCode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

// Add constant and then return back the position at which the constant is added
func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1
}

// Creates instruction and returns back the position where instruction is set
func (c *Compiler) emit(op code.Opcode, operands ...int) int {
	// Create instruction set from opcode and operands
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins)
	return pos
}

// Pushed instruction into slice and then return back the position of instruction where it is stored
func (c *Compiler) addInstruction(ins []byte) int {
	// Get the start of the instruction where it will be set
	posNewInstruction := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	return posNewInstruction
}
