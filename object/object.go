package object

import "fmt"

// ObjectType is the type of object
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

// Object is the base for all types
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer is an integer type
type Integer struct {
	Value int64
}

// Type returns object type of integer
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Inspect returns integer value as string
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Boolean is a boolean type
type Boolean struct {
	Value bool
}

// Type returns object type of boolean
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Inspect returns boolean value as string
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Null is a null type
type Null struct{}

// Type returns object type of null
func (n *Null) Type() ObjectType { return NULL_OBJ }

// Inspect returns null as string
func (n *Null) Inspect() string { return "null" }
