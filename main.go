package main

import (
	"fmt"
	"log"

	"github.com/TimoWyszynski/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("timo")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.DBUrl)
	fmt.Println(cfg.CurrentUserName)
}
