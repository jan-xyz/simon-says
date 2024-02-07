package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Action handler that is called on a separate goroutine when a "greet" action
// is created.
func handleClick(_ app.Context, a app.Action) {
	button, ok := a.Value.(int) // Checks if a name was given.
	if !ok {
		fmt.Println("wrong type")
		return
	}

	fmt.Println("clicked button:", button)
}
