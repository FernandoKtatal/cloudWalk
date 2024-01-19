package main

import (
	"example.com/quake/internal/game"
	"fmt"
	"log"
)

func main() {
	filePath := "pkg/log/qgames.log"

	games, err := game.ParseLogFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, g := range games {
		result, err := game.GetGameResult(g)
		if err != nil {
			fmt.Println(err)
			break
		}

		weaponsResult, err := game.GetGameWeaponsResult(g)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(result)
		fmt.Println(weaponsResult)
	}
}
