package object

import (
	"compiler/constants"
	"fmt"
)

type ObjectType string

// Object representation for transpiling

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
