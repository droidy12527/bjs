package evaluator

import (
	"compiler/ast"
	"compiler/constants"
	"compiler/object"
)

// This is predeclared variable block for memory allocation for true, false and null objects in memory
// as all of them are common and we do not need to create new memory allocation everytime
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
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	}
	return nil
}

// Eval Inflix expression returns back the infix expression, It checks if both of the right and left nodes of the ast
// are integer, If so then it returns back the integer object back
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == constants.INTEGER_OBJECT && right.Type() == constants.INTEGER_OBJECT:
		return evalIntegerInflixExpression(operator, left, right)
	case operator == "==":
		return nativeBooleanToBooleanObject(left == right)
	case operator == "!=":
		return nativeBooleanToBooleanObject(left != right)
	default:
		return NULL
	}
}

// Eval integer Inflix expression just switches the string that is given and then returns back the computation
func evalIntegerInflixExpression(operator string, left, right object.Object) object.Object {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftValue + rightValue}
	case "-":
		return &object.Integer{Value: leftValue - rightValue}
	case "*":
		return &object.Integer{Value: leftValue * rightValue}
	case "/":
		return &object.Integer{Value: leftValue / rightValue}
	case "<":
		return nativeBooleanToBooleanObject(leftValue < rightValue)
	case ">":
		return nativeBooleanToBooleanObject(leftValue > rightValue)
	case "==":
		return nativeBooleanToBooleanObject(leftValue == rightValue)
	case "!=":
		return nativeBooleanToBooleanObject(leftValue != rightValue)
	default:
		return NULL
	}
}

// check for prefix operator and then check for what to do next
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != constants.INTEGER_OBJECT {
		return NULL
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

// This function inverts the bang operator by sending back object
// Basically we check for ! and if we get the prefix then we invert back the results
// The results are then inverted and sent, Mainly this is used because we cannot just send back boolean
// We need to send back the address of the boolean we created in memory. This makes the things faster
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return FALSE
	default:
		return FALSE
	}
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
