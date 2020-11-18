package cli

import (
	"flag"
	"fmt"
	"os"
)

type appEnv struct {
	// args
	services map[string]bool

	// flags
	longList bool
}

func CLI(args []string) int {
	var app appEnv

	switch err := app.fromArgs(args); err {
	case nil:
		// Do nothing
	case flag.ErrHelp:
		return 0 // Fail nicely
	default:
		fmt.Fprintf(os.Stderr, "Could not parse args: %v\n\tErr: %v\n", args, err)
		return 2
	}

	return app.run()
}

func (a *appEnv) run() int {
	defaultFormat(workerDispacher(a.services))
	fmt.Printf("%v\n", a)
	return 0
}

func (a *appEnv) fromArgs(args []string) error {
	// Flags
	fl := flag.NewFlagSet("hlstat", flag.ContinueOnError)
	fl.BoolVar(&a.longList, "long", false, "Display details for each service")
	fl.BoolVar(&a.longList, "l", false, "Shorthand for -long")
	// fl.BoolVar(&a.routerPing, "router", false, "Ping the /status endpoint of the Google Wifi router")
	// fl.BoolVar(&a.gatewayPing, "gateway", false, "Pings the apartnet gateway")
	// fl.BoolVar(&a.raspberryPi, "rpi", false, "Pings the raspberrPis")
	// fl.BoolVar(&a.checkDNS, "dns", false, "Checks a public DNS server")
	// fl.BoolVar(&a.checkDiscordBot, "discord", false, "Checks the DiscordBot")

	if err := fl.Parse(args); err != nil {
		return err
	}

	// Args

	a.services = map[string]bool{
		"router":  false,
		"gateway": false,
		"rpi":     false,
		"dns":     false,
		"d-bot":   false,
	}

	for _, arg := range fl.Args() {
		if _, ok := a.services[arg]; !ok {
			return fmt.Errorf("Unexpected argument '%s'", arg)
		}
		a.services[arg] = true
	}

	return nil

}
