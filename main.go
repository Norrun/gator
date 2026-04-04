package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Norrun/gator/internal/config"
	"github.com/Norrun/gator/internal/database"

	//drivers
	_ "github.com/lib/pq"
)

func main() {
	conf := config.Read()
	db, err := sql.Open("postgres", conf.DbURL)
	queries := database.New(db)
	stt := state{queries, &conf}
	cmds := commands{make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("expected a command")
		os.Exit(1)
	}
	cmd := command{args[1], args[2:]}
	err = cmds.run(&stt, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
