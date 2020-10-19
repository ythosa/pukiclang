package object

import "fmt"

type ObjectType string

const (
    INTEGER_OBJ = "INTEGER"
    BOOLEAN_OBJ = "BOOLEAN"
)

type Object interface {
    Type() ObjectType
    Inspect() string
}

type Integer struct {
    Value int64
}

func (i *Integer) Inspect() string {
    return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
    return INTEGER_OBJ
}

type Boolean struct {
    Value bool
}

func (b *Boolean) Inspect() string {
    return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
    return BOOLEAN_OBJ
}
