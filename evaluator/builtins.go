package evaluator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	"print": {
		Fn: func(args ...object.Object) object.Object {
			var str string
			for _, arg := range args {
				str = str + arg.Inspect() + " "
			}
			fmt.Printf("%s\n", str)
			return NONE
		},
	},
	"str": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() == object.STRING_OBJ {
				return args[0]
			}

			return &object.String{Value: args[0].Inspect()}
		},
	},
	"int": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Integer:
				return args[0]
			case *object.String:
				if i, err := strconv.ParseInt(arg.Value, 10, 64); err == nil {
					return &object.Integer{Value: int64(i)}
				}
				return newError("could not convert type STRING to INTEGER")
			case *object.Boolean:
				if arg == TRUE {
					return &object.Integer{Value: int64(1)}
				}
				return &object.Integer{Value: int64(0)}
			}
			return newError("argument to `int` not supported, got %s", args[0].Type())
		},
	},
	"type": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			return &object.String{Value: string(args[0].Type())}
		},
	},
	"input": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 1 {
				fmt.Fprintf(os.Stdout, args[0].Inspect())
			}

			scanner := bufio.NewScanner(os.Stdin)
			scanned := scanner.Scan()
			if !scanned {
				return &object.String{Value: ""}
			}
			return &object.String{Value: scanner.Text()}
		},
	},
	"set": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("invalid array type. got=%s, want=ARRAY", args[0].Type())
			}
			arr := args[0].(*object.Array)
			if args[1].Type() != object.INTEGER_OBJ {
				return newError("invalid index type. got=%s. want=INTEGER", args[1].Type())
			}
			index := args[1].(*object.Integer)
			if int(index.Value) >= len(arr.Elements) || index.Value < 0 {
				return newError("index %d out of range", index.Value)
			}
			args[0].(*object.Array).Elements[int(index.Value)] = args[2]
			return NONE
		},
	},
	"append": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("invalid array type. got=%s, want=ARRAY", args[0].Type())
			}
			elements := args[0].(*object.Array).Elements
			elements = append(elements, args[1])
			args[0].(*object.Array).Elements = elements
			return NONE
		},
	},
	"split": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 2 || len(args) < 1 {
				return newError("wrong number of arguments. got=%d, want=1 or 2", len(args))
			}
			sep := " "
			arr := &object.Array{}
			if len(args) == 2 {
				if args[1].Type() != object.STRING_OBJ {
					return newError("invalid seperator type. got=%s, want=STRING", args[0].Type())
				}
				sep = args[1].(*object.String).Value
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("invalid string type. got=%s, want=STRING", args[0].Type())
			}
			str := args[0].(*object.String).Value
			tempElems := strings.Split(str, sep)
			for _, elem := range tempElems {
				arr.Elements = append(arr.Elements, &object.String{Value: elem})
			}
			return arr
		},
	},
}
