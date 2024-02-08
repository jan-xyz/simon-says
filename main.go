package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	serve := flag.Bool("serve", false, "set serve to serve instead of generating resources")
	flag.Parse()

	logic := NewSimonSaysLogic()
	app.Handle(click, logic.handleClick)

	ui := NewSimonSaysUI()
	app.Route("/", ui)

	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	if !*serve {
		err := app.GenerateStaticWebsite("_site", &app.Handler{
			Name:        "Hello",
			Description: "An Hello World! example",
			Resources:   app.GitHubPages("REPOSITORY_NAME"),
		})
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Simon Says",
		Description: "A game of simon says",
		Styles: []string{
			"/web/styles.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
