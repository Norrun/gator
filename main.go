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
	stt, err := loadState()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	cmds := registerCommands()

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

func loadState() (state, error) {
	conf := config.Read()
	db, err := sql.Open("postgres", conf.DbURL)
	if err != nil {
		return state{}, err
	}
	queries := database.New(db)
	stt := state{queries, &conf}
	return stt, nil
}

func registerCommands() commands {
	cmds := commands{make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	return cmds
}
