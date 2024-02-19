package ui

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/jan-xyz/simon-says/storage"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type scoreBoard struct {
	app.Compo

	easyWinRatio     float64
	mediumWinRatio   float64
	hardWinRatio     float64
	endlessHighscore int
}

// OnMount implements the Mounter interface to run this on mounting the component.
func (b *scoreBoard) OnMount(ctx app.Context) {
	ctx.Handle(storage.EventScoreUpdate, b.HandleScoreUpdate)
	s := storage.LoadScores(ctx)
	b.storeScores(s)
}

func (b *scoreBoard) storeScores(s storage.Scores) {
	b.easyWinRatio = float64(s.Basic[storage.Easy].Win) / float64(s.Basic[storage.Easy].Win+s.Basic[storage.Easy].Loss)
	b.mediumWinRatio = float64(s.Basic[storage.Medium].Win) / float64(s.Basic[storage.Medium].Win+s.Basic[storage.Medium].Loss)
	b.hardWinRatio = float64(s.Basic[storage.Hard].Win) / float64(s.Basic[storage.Hard].Win+s.Basic[storage.Hard].Loss)

	max := 0
	for score := range s.Endless {
		if score > max {
			max = score
		}
	}
	b.endlessHighscore = max
}

// Render implements the interface for go-app to render the component.
func (b *scoreBoard) Render() app.UI {
	easyText := "no game yet"
	if !math.IsNaN(b.easyWinRatio) {
		easyText = fmt.Sprintf("%.1f%%", b.easyWinRatio*100)
	}
	mediumText := "no game yet"
	if !math.IsNaN(b.mediumWinRatio) {
		mediumText = fmt.Sprintf("%.1f%%", b.mediumWinRatio*100)
	}
	hardText := "no game yet"
	if !math.IsNaN(b.hardWinRatio) {
		hardText = fmt.Sprintf("%.1f%%", b.hardWinRatio*100)
	}

	return app.Table().Class("scores").Body(
		app.Tr().Body(app.Td().Text("Easy"), app.Td().Text(easyText)),
		app.Tr().Body(app.Td().Text("Medium"), app.Td().Text(mediumText)),
		app.Tr().Body(app.Td().Text("Hard"), app.Td().Text(hardText)),
		app.Tr().Body(app.Td().Text("Endless"), app.Td().Text(b.endlessHighscore)),
		&GoAppBar{Class: "chart1-cls"},
	)
}

func (b *scoreBoard) HandleScoreUpdate(_ app.Context, a app.Action) {
	scores, ok := a.Value.(storage.Scores)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	b.storeScores(scores)
}

func newChart() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "go-echarts used with go-app",
		Subtitle: "Subtitle goes here",
	}))

	generateBarItems := func() []opts.BarData {
		items := make([]opts.BarData, 0)
		for i := 0; i < 7; i++ {
			items = append(items, opts.BarData{Value: rand.Intn(300)})
		}
		return items
	}

	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

type GoAppBar struct {
	app.Compo
	Class           string
	Options         *charts.Bar
	eChartsInstance app.Value
}

func (c *GoAppBar) OnMount(ctx app.Context) {
	ctx.After(time.Second, func(context app.Context) {
		c.eChartsInstance = app.Window().Get("echarts").
			Call("init", c.JSValue(), c.Options.Theme)
		c.UpdateConfig(context, c.Options)
	})
}

func (c *GoAppBar) OnDismount() {
	if c.eChartsInstance != nil {
		c.eChartsInstance.Call("dispose")
	}
}

func (c *GoAppBar) UpdateConfig(ctx app.Context, config *charts.Bar) {
	config.Validate()
	c.Options = config

	if c.eChartsInstance != nil {
		c.eChartsInstance.Call("dispose")
	}
	c.eChartsInstance = app.Window().Get("echarts").
		Call("init", c.JSValue(), c.Options.Theme)

	ctx.Async(func() {
		jsonString, _ := json.Marshal(c.Options.JSON())
		options := app.Window().Get("JSON").Call("parse", string(jsonString))
		c.eChartsInstance.Call("setOption", options)
		c.Update()
	})
}

func (c *GoAppBar) Render() app.UI {
	if c.Options == nil {
		c.Options = charts.NewBar()
		c.Options.Validate()
	}
	return app.Div().Class(c.Class).ID(c.Options.ID).
		Style("width", c.Options.Initialization.Width).
		Style("height", c.Options.Initialization.Height)
}
