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

const (
	dark  = "#05062d"
	light = "#a4a4a4"
)

func newBarChart(scores map[int]int) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:      "Endless Score Distribution",
			TitleStyle: &opts.TextStyle{Color: dark},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			BackgroundColor: light,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			SplitArea: &opts.SplitArea{Show: false},
			SplitLine: &opts.SplitLine{Show: false},
			AxisLabel: &opts.AxisLabel{Show: true, Color: dark},
			AxisLine:  &opts.AxisLine{Show: true, LineStyle: &opts.LineStyle{Color: dark}},
		}, 0),
		charts.WithXAxisOpts(opts.XAxis{
			SplitArea: &opts.SplitArea{Show: false},
			SplitLine: &opts.SplitLine{Show: false},
			AxisLabel: &opts.AxisLabel{Show: true, Color: dark},
			AxisTick:  &opts.AxisTick{Show: true},
		}, 0),
		charts.WithColorsOpts(opts.Colors{dark}),
	)
	bar.SetSeriesOptions(
		charts.WithItemStyleOpts(opts.ItemStyle{
			BorderColor: dark,
			BorderWidth: 0,
		}),
	)

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

	bar.SetXAxis(xAxis).AddSeries("", series)
	return bar
}

type barChart struct {
	app.Compo
	Class           string
	Bar             *charts.Bar
	eChartsInstance app.Value
}

func (c *barChart) OnMount(ctx app.Context) {
	ctx.After(50*time.Millisecond, func(ctx app.Context) {
		ctx.Defer(func(_ app.Context) {
			c.eChartsInstance = app.Window().Get("echarts").
				Call("init", c.JSValue())
			c.UpdateBarChart(c.Bar)
		})
	})
	ctx.Handle(storage.EventScoreUpdate, c.HandleScoreUpdate)
}

func (c *barChart) HandleScoreUpdate(_ app.Context, a app.Action) {
	s, ok := a.Value.(storage.Scores)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	c.UpdateBarChart(newBarChart(s.Endless))
}

func (c *barChart) OnDismount() {
	if c.eChartsInstance != nil {
		c.eChartsInstance.Call("dispose")
	}
}

func (c *barChart) OnResize(ctx app.Context) {
	if c.eChartsInstance != nil {
		c.eChartsInstance.Call("resize")
	}
}

func (c *barChart) UpdateBarChart(config *charts.Bar) {
	config.Validate()
	c.Bar = config

	if c.eChartsInstance != nil {
		c.eChartsInstance.Call("dispose")
	}
	c.eChartsInstance = app.Window().Get("echarts").
		Call("init", c.JSValue())

	jsonString, err := json.Marshal(c.Bar.JSON())
	if err != nil {
		panic(err)
	}
	options := app.Window().Get("JSON").Call("parse", string(jsonString))
	c.eChartsInstance.Call("setOption", options)
	c.Update()
}

func (c *barChart) Render() app.UI {
	if c.Bar == nil {
		c.Bar = charts.NewBar()
		c.Bar.Validate()
	}
	return app.Div().Class("chart").ID(c.Bar.ID)
}
