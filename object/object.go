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

type Null struct {
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() string    { return constants.INTEGER_OBJECT }

func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() string    { return constants.BOOLEAN_OBJECT }

func (n *Null) Inspect() string { return "null" }
func (n *Null) Type() string    { return constants.NULL_OBJECT }
