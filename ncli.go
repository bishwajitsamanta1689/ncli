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
		Name:    "host",
		Value:   "Enter Host Details",
		Aliases: []string{"ho"},
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
	app.Version = "1.0.1"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Bishwajit",
			Email: "bishwajitsamanta1689@gmail.com",
		},
	}
}

func command() {
	app.Commands = []*cli.Command{
		{
			Name:    "nameserver",
			Usage:   "Looks up the Name Server for Particular Host",
			Flags:   myFlags,
			Aliases: []string{"ns"},
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
		{
			Name:    "ipAddress",
			Usage:   "Looks up the IP Addresses for a particular Host",
			Flags:   myFlags,
			Aliases: []string{"ip"},
			Action: func(context *cli.Context) error {
				ip, err := net.LookupIP(context.String("host"))
				handleError(err)
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:    "mxrecords",
			Usage:   "Looks up for MX Record for a particular Host",
			Flags:   myFlags,
			Aliases: []string{"mx"},
			Action: func(context *cli.Context) error {
				mxRecords, err := net.LookupMX(context.String("host"))
				handleError(err)
				for i := 0; i < len(mxRecords); i++ {
					fmt.Println(mxRecords[i].Host)
				}
				return nil
			},
		},
		{
			Name:    "cname",
			Usage:   "Looks up for the CNAME for a Particular Host",
			Flags:   myFlags,
			Aliases: []string{"cn"},
			Action: func(context *cli.Context) error {
				cname, err := net.LookupCNAME(context.String("host"))
				handleError(err)
				fmt.Println(cname)
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
