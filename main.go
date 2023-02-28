package main

import (
	"micgograd/quads"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	data := floats.Span(make([]float64, 50), 4000, 300)

	a := mat.NewDense(25, 2, data)
	x, y := a.Dims()
	datas := quads.GenerateLineItems(x, y, a)
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeRomantic}),
		charts.WithTitleOpts(opts.Title{
			Title:    "ayylmao",
			Subtitle: "Line chart",
		}))

	line.SetXAxis([]string{}).
		AddSeries("Number", datas).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: false}))
	line.Render(w)
}
