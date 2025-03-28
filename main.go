package main

import (
	"fmt"

	"github.com/nonezerone/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("ERROR WHILE READING FILE")
	}
	err = cfg.SetUser()
	if err != nil {
		fmt.Println("ERROR WHILE SETTING USER")
	}

	newCfg, err := config.Read()
	if err != nil {
		fmt.Println("ERROR WHILE READING FILE")
	}

	fmt.Println(newCfg.DBUrl)
	fmt.Println(newCfg.CurrentUserName)
}
