package object

// NewEnclosedEnvironment returns a new enclosed env
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment returns an empty environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment has a map of objects and names
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get returns object from environment store
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets name in map to object and returns object
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
