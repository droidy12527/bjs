package object

import (
	"bytes"
	"compiler/ast"
	"compiler/constants"
	"fmt"
	"strings"
)

type ObjectType string

type BuiltinFunction func(args ...Object) Object

// Object representation for transpiling and also object representation in base language, ie GO for making sure
// that the object has given memory location, as further the objects will be needed to access the values

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

type Boolean struct {
	Value bool
}

type ReturnValue struct {
	Value Object
}

type Error struct {
	Message string
}

type Null struct {
}

type String struct {
	Value string
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Enviornment
}

type Builtin struct {
	Fn BuiltinFunction
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return constants.INTEGER_OBJECT }

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return constants.BOOLEAN_OBJECT }

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return constants.NULL_OBJECT }

func (rv *ReturnValue) Type() ObjectType { return constants.RETURN_VALUE_OBJECT }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

func (e *Error) Type() ObjectType { return constants.ERROR_OBJECT }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

func (f *Function) Type() ObjectType { return constants.FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

func (s *String) Type() ObjectType { return constants.STRING_OBJECT }
func (s *String) Inspect() string  { return s.Value }

func (b *Builtin) Type() ObjectType { return constants.BUILTIN_OBJECT }
func (b *Builtin) Inspect() string  { return "builtin function" }
