package evaluator

import (
	"compiler/ast"
	"compiler/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

// Function evaluator mrecieves ast.Node and stores into memory for representation
// Eval returns object which is stored into memory which is represented in golang struct
// For debugging purpose the struct takes more memory in ram.
// This will be fixed in upcoming versions
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBooleanToBooleanObject(node.Value)
	}
	return nil
}

// This function is optimisation to return back the value of boolean
// As boolean can be of only type true and false we can keep them in memory once someone allocates them
// Then we can reference via address to the same boolean rather than creating new boolean in memory location every time
// This saves us time and also makes sure that more memory is not allocated for garbage collector to sweep
func nativeBooleanToBooleanObject(value bool) object.Object {
	if value {
		return TRUE
	}
	return FALSE
}

// If eval statement receives ast statements then it parses and returns the object result back
func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
	}
	return result
}
