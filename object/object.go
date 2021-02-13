package object

import (
	"bytes"
	"fmt"
	"strings"
	"zlang/ast"
)

// ObjectType is the type of object
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	FLOAT_OBJ        = "FLOAT"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
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

// Float is a float type
type Float struct {
	Value float64
}

// Type returns object type of float
func (i *Float) Type() ObjectType { return FLOAT_OBJ }

// Inspect returns integer value as float
func (i *Float) Inspect() string { return fmt.Sprintf("%f", i.Value) }

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

// None is an empty type
type None struct{}

// Type returns object type of none
func (n *None) Type() ObjectType { return NULL_OBJ }

// Inspect returns none as string
func (n *None) Inspect() string { return "" }

// ReturnValue is a return value type
type ReturnValue struct {
	Value Object
}

// Type returns object type of return value
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Inspect returns return value as string
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error is an error type
type Error struct {
	Message string
}

// Type returns object type of error
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// Inspect returns error message as string
func (e *Error) Inspect() string { return "bruh moment: " + e.Message }

// Function is a function type
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns object type of function
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

// Inspect returns function as string
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

// String is a string type
type String struct {
	Value string
}

// Type returns object type of string
func (s *String) Type() ObjectType { return STRING_OBJ }

// Inspect returns string as string
func (s *String) Inspect() string { return s.Value }

// BuiltinFunction is a built in function type
type BuiltinFunction func(args ...Object) Object

// Builtin is a BuiltinFunction wrapper type
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns object type of builtin
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }

// Inspect returns builtin string
func (b *Builtin) Inspect() string { return "builtin function" }

// Array is an Array type
type Array struct {
	Elements []Object
}

// Type returns object type of array
func (ao *Array) Type() ObjectType { return ARRAY_OBJ }

// Inspect returns array as string
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		s := e.Inspect()
		if e.Type() == STRING_OBJ {
			s = fmt.Sprintf(`"%s"`, e.Inspect())
		}
		elements = append(elements, s)
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
