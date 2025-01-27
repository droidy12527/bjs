// This package contains compiler for compiling the instructions
package compiler

import (
	"compiler/ast"
	"compiler/code"
	"compiler/object"
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
	return nil
}

// Returns bytecode struct, This struct has bytecode
// Code has functions that generates
func (c *Compiler) ByteCode() *ByteCode {
	return &ByteCode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}
