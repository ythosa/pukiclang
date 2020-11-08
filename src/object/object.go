package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/ythosa/pukiclang/src/ast"
)

// Type is type of object which represented with string
type Type string

// NewEnvironment returns new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{
		store: s,
		outer: nil,
	}
}

// NewEnclosedEnvironment returns new environment with pointer on outer environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

// Environment is type for environment
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get returns object in environment and is environment contains it with passed name
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

// Set in environment variable with key = `name` and value = `value`
func (e *Environment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}

// Object interface
type Object interface {
	Type() Type
	Inspect() string
}

// BuiltInFunction is type for built in interpeter function
type BuiltInFunction func(args ...Object) Object

// BuiltIn is type for built in functionality into interpreter
type BuiltIn struct {
	Fn BuiltInFunction
}

// Inspect returns string representation of object
func (bi *BuiltIn) Inspect() string {
	return fmt.Sprintf("built in function")
}

// Type returns type of object
func (bi *BuiltIn) Type() Type {
	return BuiltInObj
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
func (i *Integer) Type() Type {
	return IntegerObj
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
func (b *Boolean) Type() Type {
	return BooleanObj
}

// String is type for string expressions
type String struct {
	Value string
}

// Inspect returns string representation of object
func (s *String) Inspect() string {
	return s.Value
}

// Type returns type of object
func (s *String) Type() Type {
	return StringObj
}

// Null is type for null expressions
type Null struct{}

// Inspect returns string representation of object
func (n *Null) Inspect() string {
	return "null"
}

// Type returns type of object
func (n *Null) Type() Type {
	return NullObj
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
func (rv *ReturnValue) Type() Type {
	return ReturnValueObj
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
func (e *Error) Type() Type {
	return ErrorObj
}

// Function is type for function object
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns type of object
func (f *Function) Type() Type {
	return FunctionObj
}

// Inspect returns string representation of object
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

// Array is type for array objects
type Array struct {
	Elements []Object
}

// Type returns type of object
func (a *Array) Type() Type {
	return ArrayObj
}

// Inspect returns string representation of object
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// Hash is type for hash map objects
type Hash struct {
	Pairs map[Object]Object
}

// Type returns type of object
func (h *Hash) Type() Type {
	return HashObj
}

// Inspect returns string representation of object
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for k, v := range h.Pairs {
		pairs = append(pairs, k.Inspect()+":"+v.Inspect())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// HashKey is type for keys in hash object
type HashKey struct {
	Type  Type
	Value uint64
}

// HashKey return HashKey object for the current object
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{
		Type:  b.Type(),
		Value: value,
	}
}

// HashKey return HashKey object for the current object
func (i *Integer) HashKey() HashKey {
	return HashKey{
		Type:  i.Type(),
		Value: uint64(i.Value),
	}
}

// HashKey return HashKey object for the current object
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{
		Type:  s.Type(),
		Value: h.Sum64(),
	}
}
