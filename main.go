package main

import (
	"fmt"
	"os"

	"github.com/Norrun/gator/internal/config"
)

func main() {
	conf := config.Read()
	stt := state{&conf}
	cmds := commands{make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("expected a command")
		os.Exit(1)
	}
	cmd := command{args[1], args[2:]}
	err := cmds.run(&stt, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
