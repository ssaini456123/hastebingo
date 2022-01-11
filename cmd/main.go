package main

import (
	"os"
	"github.com/urfave/cli/v2"
	"github.com/sa111n111/hastebingo"
	"fmt"
)

func main() {
    app := &cli.App{
		Name: "hastebingo",
	    	Usage: "Create a hastebin document. Wrap your input in double quotes (""). ",
		Action: func(c *cli.Context) error {
			h := hastebingo.Hastebin{}
			h.Post(c.Args().Get(0))
			fmt.Println("Created hastebin document.\nYour generated key is: " + h.GetKey())
			
			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		panic(err)
	}
}
