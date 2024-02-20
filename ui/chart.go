package ui

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

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
