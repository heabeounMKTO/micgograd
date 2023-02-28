package main

import (
	"math/rand"
	"micgograd/quads"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"gonum.org/v1/gonum/mat"
)

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	data := make([]float64, 49)
	for i := range data {
		data[i] = float64(rand.Int31n(20))
	}
	a := mat.NewDense(7, 7, data)
	x, y := a.Dims()
	datas := quads.GenerateLineItems(x, y, a)
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeRomantic}),
		charts.WithTitleOpts(opts.Title{
			Title:    "ayylmao",
			Subtitle: "Line chart",
		}))

	line.SetXAxis([]string{"1", "2", "3", "4", "5", "6"}).
		AddSeries("Number", datas).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}
