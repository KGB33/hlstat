package cli

import (
	"flag"
	"fmt"
	"os"
)

type appEnv struct {
	// args

	// flags
	routerPing bool
	gatewayPing bool
	raspberryPi bool
	checkDNS bool
	checkDiscordBot bool
}

func CLI(args []string) int {
	var app appEnv

	switch	err := app.fromArgs(args); err {
	case nil:
		// Do nothing
	case flag.ErrHelp:
		return 0 // Fail nicely
	default:
		fmt.Fprintf(os.Stderr, "Could not parse args: %v\nErr: %v\n", args, err)
		return 2
	}

	return app.run()
}

func (a *appEnv) run() int {
	fmt.Printf("%v", a)
	return 0
}

func (a *appEnv) fromArgs(args []string) error {
	// Flags
	fl := flag.NewFlagSet("hlstat", flag.ContinueOnError)
	fl.BoolVar(&a.routerPing, "router", false, "Ping the /status endpoint of the Google Wifi router")
	fl.BoolVar(&a.gatewayPing, "gateway", false, "Pings the apartnet gateway")
	fl.BoolVar(&a.raspberryPi, "rpi", false, "Pings the raspberrPis")
	fl.BoolVar(&a.checkDNS, "dns", false, "Checks a public DNS server")
	fl.BoolVar(&a.checkDiscordBot, "discord", false, "Checks the DiscordBot")

	err := fl.Parse(args)
	return err
}
