package main

import (
	"flag"
	"fmt"
	"monkey-ball/repl"
	"os"
	"os/user"
)

var engine = flag.String("engine", "vm", "user 'vm' or 'eval'")

func main() {
	flag.Parse()
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, welcome to monkey. Currently using %s engine\n",
		user.Username, *engine)

	if *engine == "vm" {
		repl.StartCompiler(os.Stdin, os.Stdout)
	} else {
		repl.StartInterpreter(os.Stdin, os.Stdout)
	}
}
