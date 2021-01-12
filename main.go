package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"zlang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("hello %s, welcome to z\n", user.Username)
	fmt.Printf("type some commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
