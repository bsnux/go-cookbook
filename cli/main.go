package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "IPs Lookup"
	app.Usage = "Getting IPv4 and IPv6 addresses for given hostname"

	app.Author = "Arturo Fernandez"
	app.Email = "artufm at yahoo dot es"
	// Changing default values
	app.Version = "0.1"
	app.UsageText = fmt.Sprintf("%s hostname", os.Args[0])

	app.Action = func(c *cli.Context) error {
		// Checking we're getting one argument
		if len(os.Args) < 2 {
			cli.ShowAppHelpAndExit(c, 0)
		}

		ip, err := net.LookupIP(c.Args().Get(0))
		if err != nil {
			fmt.Println(err)
		}
		// Priting addresses
		for i := 0; i < len(ip); i++ {
			fmt.Println(ip[i])
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
