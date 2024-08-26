package object

import (
	"compiler/constants"
	"fmt"
)

type ObjectType string

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
