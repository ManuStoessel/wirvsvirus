package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

var (
	clientid string
	issuer   string
)

func main() {

	app := &cli.App{
		Name:  "wirvsvirusbackend",
		Usage: "This will start the wirvsvirus backend http server.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "port",
				Value: "8080",
				Usage: "Start HTTP server on port `PORT`",
			},
			&cli.StringFlag{
				Name:        "oauth2-clientid",
				Value:       "xxxxxxxx",
				Usage:       "Okta OAuth2 client ID to use.",
				Required:    true,
				Destination: &clientid,
			},
			&cli.StringFlag{
				Name:        "oauth2-issuer",
				Value:       "xxxxxxxx",
				Usage:       "Okta OAuth2 issuer to use.",
				Required:    true,
				Destination: &issuer,
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

	router := mux.NewRouter().StrictSlash(true)

	// definition of global middlewares to use
	router.Use(loggingMiddleware)
	router.Use(verificationMiddleware)

	// definition of routes
	router.HandleFunc("/", homeLink)
	return http.ListenAndServe(":"+c.String("port"), router)
}
