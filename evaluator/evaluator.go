package evaluator

import (
	"compiler/ast"
	"compiler/constants"
	"compiler/object"
	"fmt"
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
func Eval(node ast.Node, env *object.Enviornment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.Boolean:
		return nativeBooleanToBooleanObject(node.Value)
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Function{Parameters: params, Env: env, Body: body}
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)
	case *ast.ReturnStatement:
		value := Eval(node.ReturnValue, env)
		if isError(value) {
			return value
		}
		return &object.ReturnValue{Value: value}
	}
	return nil
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok {
		return newError("not a function: %s", fn.Type())
	}
	extendedEnviornment := extendedFunctionEnviornment(function, args)
	evaluated := Eval(function.Body, extendedEnviornment)
	return unwrapReturnValue(evaluated)
}

func extendedFunctionEnviornment(fn *object.Function, args []object.Object) *object.Enviornment {
	env := object.NewEnclosedEnviornment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalExpressions(exps []ast.Expression, env *object.Enviornment) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

func evalProgram(node *ast.Program, env *object.Enviornment) object.Object {
	var result object.Object
	for _, statement := range node.Statements {
		result = Eval(statement, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Enviornment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return newError("identifier not found: " + node.Value)
	}
	return val
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Enviornment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if result != nil {
			rt := result.Type()
			if rt == constants.RETURN_VALUE_OBJECT || rt == constants.ERROR_OBJECT {
				return result
			}
		}
	}
	return result
}

func evalIfExpression(ie *ast.IfExpression, env *object.Enviornment) object.Object {
	condition := Eval(ie.Condition, env)
	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

// Checks if the object is error object or not
func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == constants.ERROR_OBJECT
	}
	return false
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
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	case left.Type() == constants.STRING_OBJECT && right.Type() == constants.STRING_OBJECT:
		return evalStringInfixExpression(operator, left, right)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
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
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
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
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != constants.INTEGER_OBJECT {
		return newError("unknown operator: -%s", right.Type())
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	if operator != "+" {
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
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

// Function returns new error object to the caller, This is used to send the error to the user.
// This function basically formats the string and returns back the error object address stored in memory.
// Uses pointer to do the work.
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
