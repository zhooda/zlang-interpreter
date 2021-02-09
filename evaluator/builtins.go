package evaluator

import (
	"os"
	"zlang/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"exit": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError("wrong number of arguments. got=%d, want=1 (optional)", len(args))
			}

			if len(args) == 1 {
				switch arg := args[0].(type) {
				case *object.Integer:
					os.Exit(int(arg.Value))
				default:
					return newError("argument to `exit` not supported, got %s", args[0].Type())
				}
			}

			os.Exit(0)
			return &object.Integer{Value: int64(0)}
		},
	},
}
