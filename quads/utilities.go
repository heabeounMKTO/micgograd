package quads

import (
	"math/rand"

	"github.com/go-echarts/go-echarts/v2/opts"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot/plotter"
)

func RandomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

func GenerateLineItems(x int, y int, matrix *mat.Dense) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			items = append(items, opts.LineData{Value: matrix.At(i, j)})
		}

	}
	return items
}

func GenerateBarItems(x int, y int, matrix *mat.Dense) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			items = append(items, opts.BarData{Value: matrix.At(i, j)})
		}

	}
	return items
}
