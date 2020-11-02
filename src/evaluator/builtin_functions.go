package evaluator

import (
	"github.com/ythosa/pukiclang/src/object"
)

var builtIns = map[string]*object.BuiltIn{
	"len":   &object.BuiltIn{Fn: lenBuiltIn},
	"first": &object.BuiltIn{Fn: first},
}

func lenBuiltIn(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{
			Value: int64(len(arg.Value)),
		}

	case *object.Array:
		return &object.Integer{
			Value: int64(len(arg.Elements)),
		}

	default:
		return newError("argument to `len` not supported, got %s",
			args[0].Type())
	}
}

func first(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		if len(arg.Value) > 0 {
			return &object.String{
				Value: string(arg.Value[0]),
			}
		}
		return NULL

	case *object.Array:
		if len(arg.Elements) > 0 {
			return arg.Elements[0]
		}

		return NULL

	default:
		return newError("argument to `len` not supported, got %s",
			args[0].Type())
	}
}
