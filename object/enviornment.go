// This file manages the object storage, Basically managing string to object mapping
// This is used to make sure the object associated with string is found back.
package object

type Enviornment struct {
	store map[string]Object
	outer *Enviornment
}

func NewEnviornment() *Enviornment {
	s := make(map[string]Object)
	return &Enviornment{store: s, outer: nil}
}

func (e *Enviornment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}

func (e *Enviornment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func NewEnclosedEnviornment(outer *Enviornment) *Enviornment {
	env := NewEnviornment()
	env.outer = outer
	return env
}
