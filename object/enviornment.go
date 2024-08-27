// This file manages the object storage, Basically managing string to object mapping
// This is used to make sure the object associated with string is found back.
package object

type Enviornment struct {
	store map[string]Object
}

func NewEnviornment() *Enviornment {
	s := make(map[string]Object)
	return &Enviornment{store: s}
}

func (e *Enviornment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}

func (e *Enviornment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}
