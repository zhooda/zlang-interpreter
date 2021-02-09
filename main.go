package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"zlang/evaluator"
	"zlang/file"
	"zlang/lexer"
	"zlang/object"
	"zlang/parser"
	"zlang/repl"
)

const (
	version     = "0.1.0"
	buildString = "0x000047"
	buildDate   = "02/07/2020"
	goVersion   = "go version go1.15.8"
)

// z 0.1.0 (v0.1.0:build string, build_date)
// [go version go1.15.8]

func main() {
	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	if len(os.Args) == 1 {
		fmt.Printf("z %s (v%s:%s, %s)\n", version, version, buildString, buildDate)
		fmt.Printf("[%s]\n\n", goVersion)
		fmt.Printf("hello %s, type some commands\n", user.Username)
		repl.Start(os.Stdin, os.Stdout)
	}
	if len(os.Args) == 2 {
		fname := os.Args[1]
		file := file.NewFile(fname)
		env := object.NewEnvironment()
		l := lexer.New(file.String())
		p := parser.New(l)

		program := p.ParseProgram()

		evaluated := evaluator.Eval(program, env)
		fmt.Println(evaluated.Inspect())
	}
}
