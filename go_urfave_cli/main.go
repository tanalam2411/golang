package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
)

/*
go get go get github.com/urfave/cli
*/

func main() {

	app := cli.NewApp()
	app.Name = "Website lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers!"

	appFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "geekslib.com", // default value for arg - host
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the Name Servers for a particular Host",
			Flags: appFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
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
			Name:  "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Flags: appFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}

				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}

				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks CNAME for a particular host",
			Flags: appFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}

				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a particular host",
			Flags: appFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}

				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
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

/*
$ ./go_cli  help
NAME:
   Website lookup CLI - Let's you query IPs, CNAMEs, MX records and Name Servers!

USAGE:
   go_cli [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     ns       Looks up the Name Servers for a particular Host
     ip       Looks up the IP addresses for a particular host
     cname    Looks CNAME for a particular host
     mx       Looks up the MX records for a particular host
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version


$ go run go_cli/main.go ns --host geekslib.com
dns1.registrar-servers.com.
dns2.registrar-servers.com.

$ go run go_cli/main.go ip --host geekslib.com
18.191.216.130

$ go run go_cli/main.go cname --host geekslib.com
geekslib.com.

$ go run go_cli/main.go mx --host geekslib.com
eforward3.registrar-servers.com. 10
eforward2.registrar-servers.com. 10
eforward1.registrar-servers.com. 10
eforward4.registrar-servers.com. 15
eforward5.registrar-servers.com. 20

*/
