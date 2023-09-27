package main

import (
	"fmt"
	"math"
	"os"

	"github.com/wcharczuk/go-chart/v2"
)

var (
	INTERVAL       = [2]float64{0, 0.5}
	INTERVAL_ADAMS = [2]float64{0, 1}
)

const STEP float64 = 0.1

// EulerMethod is implementation from slide 13
func EulerMethod(xs, ys, yds []float64) []float64 {
	n := len(xs)
	yem := make([]float64, n)
	yem[0] = ys[0]
	for i := 1; i < n; i++ {
		yem[i] = yem[i-1] + STEP*yds[i-1]
	}
	return yem
}

// EulerCauchyMethod is implementation of Euler-Cauchy method from slide 16
func EulerCauchyMethod(xs, ys, yds []float64) ([]float64, []float64) {
	n := len(xs)

	yhats := make([]float64, n)
	// DOESN'T EXIST so yhats[0] is 0 by default
	for i := 1; i < n; i++ {
		yhats[i] = ys[i-1] + STEP*yds[i-1]
	}

	yecm := make([]float64, n)
	yecm[0] = ys[0]
	for i := 1; i < n; i++ {
		hatDirr := FirstOrderDerivative(xs[i], yhats[i])
		fraction := (yds[i-1] + hatDirr) / 2
		yecm[i] = yecm[i-1] + STEP*fraction
	}
	return yecm, yhats
}

// RungeKuttaMethod is implementation of Runge-Kutta method from slide 24
func RungeKuttaMethod(xs, ys, yds []float64) []float64 {
	n := len(xs)
	yrkts := make([]float64, n)
	yrkts[0] = ys[0]
	for i := 0; i < n-1; i++ {
		k1 := STEP * yds[i]
		k2 := STEP * FirstOrderDerivative(xs[i]+0.5*STEP, ys[i]+0.5*k1)
		k3 := STEP * FirstOrderDerivative(xs[i]+0.5*STEP, ys[i]+0.5*k2)
		k4 := STEP * FirstOrderDerivative(xs[i]+STEP, ys[i]+k3)
		deltaY := (k1 + 2*k2 + 2*k3 + k4) / 6
		yrkts[i+1] = ys[i] + deltaY
	}
	return yrkts
}

// MultistepAdamsMethod is implementation of Multistep Adams method from slide 28
func MultistepAdamsMethod(xs, ys, yds []float64) []float64 {
	n := len(xs)
	adams := make([]float64, n)
	// Method requires three previous values of y
	adams[0] = ys[0]
	adams[1] = ys[1]
	adams[2] = ys[2]
	adams[3] = ys[3]

	for i := 4; i < n; i++ {
		bracket := 55*yds[i-1] - 59*yds[i-2] + 37*yds[i-3] - 9*yds[i-4]
		adams[i] = adams[i-1] + STEP/24*bracket
	}
	return adams
}

// Equation() computes y = tg(x) - x
func Equation(x float64) float64 {
	if x == 0 {
		return 0
	}
	return math.Tan(x) - x
}

// FirstOrderDifferentialEquation() computes y' = (y + x)^2
func FirstOrderDerivative(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	sum := x + y
	return math.Pow(sum, 2)
}

func getValues(interval [2]float64) ([]float64, []float64, []float64) {
	var xs, ys, yds []float64

	for x := interval[0]; x <= interval[1]; x += STEP {
		xs = append(xs, x)
		ys = append(ys, Equation(x))
		yds = append(yds, FirstOrderDerivative(x, ys[len(ys)-1]))
	}
	return xs, ys, yds
}

func printParamsTable(method string, xs, ys, dy, ym, yhat []float64) {
	fmt.Println(method)
	if yhat != nil {
		fmt.Print(" k | x   |\ty \t| y_hats       |y_true         | eps         | dy\n")
		for i := range ym {
			fmt.Printf("%2d | %.1f | %.9f  | %.9f  | %.9f   | %.9f | %.9f   \n", i, xs[i], ym[i], yhat[i], ys[i], math.Abs(ym[i]-ys[i]), dy[i])
		}
	} else {
		fmt.Printf(" k | x   |\ty \t| y_true\t| eps \t\t| dy\n")
		for i := range ym {
			fmt.Printf("%2d | %.1f | %.9f  | %.9f   | %.9f   | %.9f \n", i, xs[i], ym[i], ys[i], math.Abs(ym[i]-ys[i]), dy[i])
		}
	}
	fmt.Println("----------------------------------------------------------------------")
}

func PlotGraphs(series []chart.Series) {
	graph := chart.Chart{
		Title:  "Comparison of different methods",
		Series: series,
		DPI:    300.0,
		Width:  1600,
		Height: 900,
	}
	graph.XAxis.Name = "X"
	graph.YAxis.Name = "Y"
	graph.TitleStyle.FontSize = 10.0
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func main() {
	xs, ys, yds := getValues(INTERVAL)

	yme := EulerMethod(xs, ys, yds)
	printParamsTable("EulerMethod", xs, ys, yds, yme, nil)

	ymce, yhats := EulerCauchyMethod(xs, ys, yds)
	printParamsTable("EulerCauchyMethod", xs, ys, yds, ymce, yhats)

	yrkts := RungeKuttaMethod(xs, ys, yds)
	printParamsTable("RungeKuttaMethod", xs, ys, yds, yrkts, nil)

	xas, yas, yads := getValues(INTERVAL_ADAMS)
	yadams := MultistepAdamsMethod(xas, yas, yads)
	printParamsTable("MultistepAdamsMethod", xas, yas, yads, yadams, nil)

	defStyle := chart.Style{
		StrokeWidth:     5,
		StrokeDashArray: []float64{0.1, 0.2, 0.3, 0.4},
	}
	series := []chart.Series{
		chart.ContinuousSeries{Style: defStyle, XValues: xs, YValues: yme, Name: "Euler"},
		chart.ContinuousSeries{Style: defStyle, XValues: xs, YValues: ymce, Name: "Euler Cauchy"},
		chart.ContinuousSeries{Style: defStyle, XValues: xs, YValues: yrkts, Name: "Runge Kutta"},
		chart.ContinuousSeries{Style: defStyle, XValues: xas, YValues: yadams, Name: "Adam"},
	}

	PlotGraphs(series)
}
