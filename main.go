package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func main() {

	info()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
