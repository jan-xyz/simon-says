// Package main is a game of simon-says. It provides a server and client application
// as well as a static-resource generator for the game.
package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/jan-xyz/simon-says/game"
	"github.com/jan-xyz/simon-says/ui"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	serve := flag.Bool("serve", false, "set to serve instead of generating resources")
	flag.Parse()

	l := game.New()
	// TODO: improve statistics
	// TODO: for endless mode add histogram of how far you got.
	// TODO: add tests
	// TODO: add dependabot
	// TODO: add linter
	app.Handle(ui.EventClick, l.HandleClick)
	app.Handle(ui.EventNewGame, l.HandleNewGame)

	g := ui.NewUI()
	app.Route("/", g)

	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	h := &app.Handler{
		Name:        "Simon Says",
		Description: "A game of simon says",
		Styles: []string{
			"/web/styles.css",
		},
		Icon: app.Icon{
			Default:    "/web/icon.png",
			Large:      "/web/icon.png",
			SVG:        "/web/icon.svg",
			AppleTouch: "/web/icon.png",
		},
		LoadingLabel:       "Loading...",
		AutoUpdateInterval: 15 * time.Minute,
	}
	if !*serve {
		h.Resources = app.GitHubPages("simon-says")
		err := app.GenerateStaticWebsite("_site", h)
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
	http.Handle("/", h)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
