package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type button struct {
	app.Compo
	Background string
}

func (h *button) Render() app.UI {
	return app.Button().
		Class("simon-button").
		Body(app.Span().Text(""))
}

func NewSimonSays() *simonSays {
	return &simonSays{}
}

type simonSays struct {
	app.Compo
}

func (h *simonSays) Render() app.UI {
	t := app.Div().Class("game-field")

	t.Body(
		&button{Background: "yellow-gradient"},
		&button{Background: "red-gradient"},
		&button{Background: "green-gradient"},
		&button{Background: "blue-gradient"},
	)

	return app.Div().Class("fill", "background").Body(
		t,
	)
}

func main() {
	app.Route("/", NewSimonSays())

	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

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

func GenerateSequence(l int) []int64 {
	seq := []int64{}
	for i := 0; i < l; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(4))
		if err != nil {
			panic(err)
		}
		seq = append(seq, n.Int64())

	}
	return seq
}
