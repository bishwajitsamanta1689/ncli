package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
)

//Variable Declaration
var app = cli.NewApp()
var myFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "host",
		Value: "",
	},
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func info() {
	app.Name = "Network CLI"
	app.Usage = "Let's you Query IP's, CNAME"
	app.Version = "0.1.0"
}

func command() {
	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the Name Server for Particular Host",
			Flags: myFlags,
			Action: func(context *cli.Context) error {
				ns, err := net.LookupNS(context.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
	}
}
func main() {
	// Importing Functions Declared
	info()
	command()

	err := app.Run(os.Args)
	handleError(err)
}
