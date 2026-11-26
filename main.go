package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/TimoWyszynski/gator/internal/config"
	"github.com/TimoWyszynski/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	dbURL := cfg.DBUrl
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	newState := state{cfg: &cfg, db: database.New(db)}

	newCommands := commands{cmdMap: make(map[string]func(*state, command) error)}

	newCommands.register("login", handlerLogin)
	newCommands.register("register", handlerRegister)

	cliArgs := os.Args[1:]

	if len(cliArgs) < 2 {
		log.Fatal("Not enough cli arguments provided")
	}

	cmdName := cliArgs[0]
	cmdArgs := cliArgs[1:]

	newCommand := command{name: cmdName, args: cmdArgs}

	err = newCommands.run(&newState, newCommand)

	if err != nil {
		log.Fatal(err)
	}

}

type state struct {
	cfg *config.Config
	db  *database.Queries
}
