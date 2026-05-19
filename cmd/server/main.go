package main

import (
	"log"

	"github.com/ajone239/star_denizen/internal"
)

const StaticDir string = "./frontend/build"

func main() {
	server := internal.NewServer(StaticDir)

	err := server.Run()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Closing up shop")
	}
}
