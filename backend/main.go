package main

import (
	"os"

	"github.com/ManuStoessel/wirvsvirus/backend/api"
	"github.com/ManuStoessel/wirvsvirus/backend/entity"
	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

var dbString string

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {

	app := &cli.App{
		Name:  "wirvsvirusbackend",
		Usage: "This will start the wirvsvirus backend http server.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "port",
				Value:       "8080",
				Usage:       "Start HTTP server on port `PORT`",
				DefaultText: "8080",
			},
			&cli.StringFlag{
				Name:        "log-level",
				Value:       "warn",
				Usage:       "Set log level to trace, debug, info, warn, error or fatal",
				DefaultText: "warn",
			},
			&cli.StringFlag{
				Name:        "db",
				Value:       "root:my-secret-pw@/wirvsvirus?charset=utf8&parseTime=True&loc=Local",
				Usage:       "Set db connection",
				DefaultText: "root:my-secret-pw@/wirvsvirus?charset=utf8&parseTime=True&loc=Local",
			},
		},
		Action: Run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// Run will start the http server
func Run(c *cli.Context) error {
	loglvl, err := log.ParseLevel(c.String("log-level"))
	if err != nil {
		log.WithField("inputLevel", c.String("log-level")).Warn("Could not parse input for log level. Defaulting to warn.")
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(loglvl)
	}

	// init db layer
	entity.Initialize(c.String("db"))

	return api.StartRouter(c.String("port"))
}
