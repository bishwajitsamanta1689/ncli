package main
import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
)
func main() {
	app:= cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you Query IP's, CNAME"
	myFlags:= []cli.Flag {
		&cli.StringFlag{
			Name: "host",
			Value: "",
		},
	}
	app.Commands =[]*cli.Command {
		{
			Name: "ns",
			Usage: "Looks up the Name Server for Particular Host",
			Flags: myFlags,
			Action: func(context *cli.Context) error {
				ns, err := net.LookupNS(context.String("host"))
				if err != nil {
					return err
				}
				for i:=0; i<len(ns);i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}