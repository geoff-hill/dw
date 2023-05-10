package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/geoff-hill/dw/pkg/lookup"
)

func main() {
	app := &cli.App{
		Name: "DW - Diceware Tools",
		Usage: `
		Helps you create diceware passwords with custom wordlists,
		and using your own dice or built in random dice`,
		Commands: []*cli.Command{
			{
				Name:        "lookup",
				HelpName:    "lookup",
				Aliases: []string{"l"},
				Action:      LookupAction,
				ArgsUsage:   ` `,
				Usage:       `Lookup a word in the diceware list`,
				Description: `Lookup a word in a diceware list`,
				Flags: []cli.Flag{},
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "verbose",
				Aliases:  []string{"v"},
				Usage:    "shows messages",
			},
		},
}
  
    err := app.Run(os.Args)
    if err !=  nil  {
        log.Fatal(err)
    }
}

func LookupAction(c *cli.Context)  error  {
	if  c.Args().Len()  ==  0  {
		return errors.New("At least one argument is expected")
	}

	// validate all args are the right size
	// read fordlists 
    if c.Bool("verbose") {
		fmt.Println(lookup.Keys())
	}
	fmt.Print(lookup.LookupByKey(c.Args().Get(0)))
	return  nil
}

