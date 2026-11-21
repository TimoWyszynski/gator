package main

import (
	"log"
	"os"

	"github.com/TimoWyszynski/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	newState := state{cfg: &cfg}

	newCommands := commands{cmdMap: make(map[string]func(*state, command) error)}

	newCommands.register("login", handlerLogin)

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
}
