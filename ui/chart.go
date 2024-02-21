package ui

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func newChart(scores map[int]int) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Endless Score Distribution",
	}))

	max := 0
	for score := range scores {
		if score > max {
			max = score
		}
	}

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

	bar.SetXAxis(xAxis).AddSeries("", series)
	return bar
}

type GoAppBar struct {
	app.Compo
	Class           string
	Bar             *charts.Bar
	eChartsInstance app.Value
}

func (c *GoAppBar) OnMount(ctx app.Context) {
	ctx.After(50*time.Millisecond, func(ctx app.Context) {
		c.eChartsInstance = app.Window().Get("echarts").
			Call("init", c.JSValue(), c.Bar.Theme)
		c.UpdateConfig(ctx, c.Bar)
	})
}

func (c *GoAppBar) OnDismount() {
	if c.eChartsInstance != nil {
		c.eChartsInstance.Call("dispose")
	}
}

func (c *GoAppBar) UpdateConfig(ctx app.Context, config *charts.Bar) {
	config.Validate()
	c.Bar = config

	if c.eChartsInstance != nil {
		c.eChartsInstance.Call("dispose")
	}
	c.eChartsInstance = app.Window().Get("echarts").
		Call("init", c.JSValue(), c.Bar.Theme)

	ctx.Async(func() {
		jsonString, err := json.Marshal(c.Bar.JSON())
		if err != nil {
			panic(err)
		}
		options := app.Window().Get("JSON").Call("parse", string(jsonString))
		c.eChartsInstance.Call("setOption", options)
		c.Update()
	})
}

func (c *GoAppBar) Render() app.UI {
	fmt.Println("render")
	if c.Bar == nil {
		c.Bar = charts.NewBar()
		c.Bar.Validate()
	}
	return app.Div().Class(c.Class).ID(c.Bar.ID).
		Style("width", c.Bar.Initialization.Width).
		Style("height", c.Bar.Initialization.Height)
}
