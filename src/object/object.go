package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
)

// NewEnvironment returns new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment is type for environment
type Environment struct {
	store map[string]Object
}

// Get returns object in environment and is environment contains it with passed name
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set in environment variable with key = `name` and value = `value`
func (e *Environment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}

// Object interface
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer interface
type Integer struct {
	Value int64
}

// Inspect returns string representation of object
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type returns type of object
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Boolean is type for boolean expressions
type Boolean struct {
	Value bool
}

// Inspect returns string representation of object
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Type returns type of object
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Null is type for null expressions
type Null struct{}

// Inspect returns string representation of object
func (n *Null) Inspect() string {
	return "null"
}

// Type returns type of object
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// ReturnValue is type for return expressions
type ReturnValue struct {
	Value Object
}

// Inspect returns string representation of object
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Type returns type of object
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// Error is type for errors handling
type Error struct {
	Message string
}

// Inspect returns string representation of object
func (e *Error) Inspect() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

// Type returns type of object
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
