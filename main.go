package main

import (
	"fmt"
	"monkey-ball/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, welcome to monkey.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
