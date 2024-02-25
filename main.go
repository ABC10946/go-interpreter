package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ABC10946/go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello %s! This is the monkey programming language!\n", user.Username)
	fmt.Printf("feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
