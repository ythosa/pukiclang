package evaluator

import (
	"github.com/ythosa/pukiclang/src/object"
)

var builtIns = map[string]*object.BuiltIn{
	"len":   &object.BuiltIn{Fn: lenBuiltIn},
	"first": &object.BuiltIn{Fn: first},
	"last":  &object.BuiltIn{Fn: last},
	"tail":  &object.BuiltIn{Fn: tail},
	"push":  &object.BuiltIn{Fn: push},
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

func last(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		if len(arg.Value) > 0 {
			return &object.String{
				Value: string(arg.Value[len(arg.Value)-1]),
			}
		}
		return NULL

	case *object.Array:
		if len(arg.Elements) > 0 {
			return arg.Elements[len(arg.Elements)-1]
		}

		return NULL

	default:
		return newError("argument to `last` not supported, got %s",
			args[0].Type())
	}
}

func tail(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		length := len(arg.Value)
		if length > 0 {
			tailOfString := make([]byte, length-1, length-1)

			copy([]byte(tailOfString), []byte(arg.Value[1:]))

			return &object.String{Value: string(tailOfString)}
		}
		return NULL

	case *object.Array:
		length := len(arg.Elements)
		if length > 0 {
			tailOfArray := make([]object.Object, length-1, length-1)

			copy(tailOfArray, arg.Elements[1:])

			return &object.Array{Elements: tailOfArray}
		}

		return NULL

	default:
		return newError("argument to `tail` not supported, got %s",
			args[0].Type())
	}
}

func push(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2",
			len(args))
	}

	if args[0].Type() != object.ArrayObj {
		return newError("argument to `push` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	newElements := make([]object.Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &object.Array{Elements: newElements}
}
