package ui

import (
	"testing"

	"github.com/jan-xyz/simon-says/storage"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/stretchr/testify/assert"
)

func TestComponentLifcycle(t *testing.T) {
	tests := []struct {
		Name       string
		matches    []app.TestUIDescriptor
		difficulty storage.Difficulty
	}{
		{
			Name: "test initial",
			matches: []app.TestUIDescriptor{
				{
					Path:     app.TestPath(0),
					Expected: app.Div().Body(),
				},
				{
					Path: app.TestPath(0, 0),
					Expected: app.Button().
						Class("simon-button", "new-game").
						OnClick(func(_ app.Context, _ app.Event) {}),
				},
				{
					Path: app.TestPath(0, 0, 0),
					Expected: app.Span().
						Text("New Game"),
				},
				{
					Path: app.TestPath(0, 1),
					Expected: app.Div().
						Class("difficulty"),
				},
				{
					Path: app.TestPath(0, 1, 0),
					Expected: app.Input().
						Type("radio").
						Name("difficulty-setting").
						ID("difficultyeasy").
						Value("easy").
						Checked(true).
						OnClick(func(_ app.Context, _ app.Event) {}),
				},
				{
					Path: app.TestPath(0, 1, 1),
					Expected: app.Label().
						For("difficultyeasy"),
				},
				{
					Path: app.TestPath(0, 1, 2),
					Expected: app.Input().
						Type("radio").
						Name("difficulty-setting").
						ID("difficultymedium").
						Value("medium").
						OnClick(func(_ app.Context, _ app.Event) {}),
				},
				{
					Path: app.TestPath(0, 1, 3),
					Expected: app.Label().
						For("difficultymedium"),
				},
				{
					Path: app.TestPath(0, 1, 4),
					Expected: app.Input().
						Type("radio").
						Name("difficulty-setting").
						ID("difficultyhard").
						Value("hard").
						OnClick(func(_ app.Context, _ app.Event) {}),
				},
				{
					Path: app.TestPath(0, 1, 5),
					Expected: app.Label().
						For("difficultyhard"),
				},
				{
					Path: app.TestPath(0, 1, 6),
					Expected: app.Input().
						Type("radio").
						Name("difficulty-setting").
						ID("difficultyendless").
						Value("endless").
						OnClick(func(_ app.Context, _ app.Event) {}),
				},
				{
					Path: app.TestPath(0, 1, 7),
					Expected: app.Label().
						For("difficultyendless"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			compo := &menu{}

			disp := app.NewClientTester(compo)
			defer disp.Close()

			for _, match := range tt.matches {
				err := app.TestMatch(compo, match)

				assert.NoError(t, err, "Path: %v", match.Path)
			}
		})
	}
}
