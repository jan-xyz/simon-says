package ui

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/jan-xyz/simon-says/storage"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func newChart(scores map[int]int) *charts.Bar {
	max := 0
	for score := range scores {
		if score > max {
			max = score
		}
	}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Endless Score Distribution",
	}))

	xAxis := []string{}
	series := []opts.BarData{}
	for i := 1; i <= max; i++ {
		xAxis = append(xAxis, strconv.Itoa(i))
		val, ok := scores[i]
		if !ok {
			val = 0
		}
		series = append(series, opts.BarData{Value: val})
	}
	fmt.Println(scores)

	bar.SetXAxis(xAxis).
		AddSeries("", series)
	return bar
}

type GoAppBar struct {
	app.Compo
	Class           string
	Options         *charts.Bar
	eChartsInstance app.Value
}

func (c *GoAppBar) OnMount(ctx app.Context) {
	s := storage.LoadScores(ctx)
	ctx.After(50*time.Millisecond, func(ctx app.Context) {
		c.eChartsInstance = app.Window().Get("echarts").
			Call("init", c.JSValue(), c.Options.Theme)

		c.UpdateConfig(ctx, newChart(s.Endless))
	})
	ctx.Handle(storage.EventScoreUpdate, c.HandleScoreUpdate)
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

func (c *GoAppBar) HandleScoreUpdate(ctx app.Context, a app.Action) {
	s, ok := a.Value.(storage.Scores)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	c.UpdateConfig(ctx, newChart(s.Endless))
}
