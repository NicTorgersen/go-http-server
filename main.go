package main

import (
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/NicTorgersen/go-http-server/controllers"
)

func run(port string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Home) 

    log.Printf("Listening on port %s", port)

	return http.ListenAndServe(":" + port, mux)
}

func main() {
    port := "8080"

	app := &cli.App{
		Name:  "http-server",
		Usage: "Run the web server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "port",
				Value: "8080",
				Usage: "Set the port the http-server should listen on.",
                Destination: &port,
			},
		},
		Action: func(ctx *cli.Context) error {
			return run(port)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
