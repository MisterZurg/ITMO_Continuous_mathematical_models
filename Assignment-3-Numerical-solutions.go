package main

import (
	"fmt"
	"github.com/wcharczuk/go-chart/v2"
	"math"
	"os"
)

// go get gonum.org/v1/plot/...

var INTERVAL = [2]float64{0, 0.5}
var INTERVAL_ADAMS = [2]float64{0, 1}

const STEP float64 = 0.1

func EulerMethod(x, y, dy []float64) []float64 {
	n := len(x)
	yAns := make([]float64, n)
	yAns[0] = y[0]
	for i := 1; i < n; i++ {
		yAns[i] = yAns[i-1] + STEP*dy[i-1]
	}
	return yAns
}

func EulerCauchyMethod(x, y, dy []float64) []float64 {
	n := len(x)
	yTilde := make([]float64, n)
	for i := 1; i < n; i++ {
		yTilde[i] = y[i-1] + STEP*dy[i-1]
	}
	yAns := make([]float64, n)
	yAns[0] = y[0]
	for i := 1; i < n; i++ {
		der := Derivative(x[i], yTilde[i])
		yAns[i] = yAns[i-1] + STEP*(dy[i-1]+der)/2
	}
	return yAns
}

func RungeKuttaMethod(x, y, dy []float64) []float64 {
	n := len(x)
	yAns := make([]float64, n)
	yAns[0] = y[0]
	for i := 0; i < n-1; i++ {
		k1 := STEP * dy[i]
		k2 := STEP * Derivative(x[i]+0.5*STEP, y[i]+0.5*k1)
		k3 := STEP * Derivative(x[i]+0.5*STEP, y[i]+0.5*k2)
		k4 := STEP * Derivative(x[i]+STEP, y[i]+k3)
		deltaY := (k1 + 2*k2 + 2*k3 + k4) / 6
		yAns[i+1] = y[i] + deltaY
	}
	return yAns
}

func MultistepAdamsMethod(x, y, dy []float64) []float64 {
	n := len(x)
	adams := make([]float64, n)
	adams[0] = y[0]
	adams[1] = y[1]
	adams[2] = y[2]
	adams[3] = y[3]
	for i := 4; i < n; i++ {
		temp := 55*dy[i-1] - 59*dy[i-2] + 37*dy[i-3] - 9*dy[i-4]
		adams[i] = adams[i-1] + STEP*temp/24
	}
	return adams
}

// Equation() computes y = tg(x) - x
func Tan(x float64) float64 {
	if x == 0 {
		return 0
	}
	return math.Tan(x) - x
}

// Derivative() computes y' = (y + x)^2
func Derivative(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	return math.Pow(x+y, 2)
}

func getValues(interval [2]float64) ([]float64, []float64, []float64) {
	var x, y, dy []float64

	for xi := interval[0]; xi <= interval[1]; xi += STEP {
		x = append(x, xi)
		y = append(y, Tan(xi))
		dy = append(dy, Derivative(xi, y[len(y)-1]))
	}
	return x, y, dy
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
	x, y, dy := getValues(INTERVAL)
	yEuler := EulerMethod(x, y, dy)

	fmt.Println("Euler Method")
	fmt.Printf(" k | x   |\ty \t| y_true\n")
	for i := range yEuler {
		fmt.Printf("%2d | %.1f | %.9f  | %.9f \n", i, x[i], yEuler[i], y[i])
	}

	fmt.Println("Euler Cauchy Method")
	fmt.Printf(" k | x   |\ty \t| y_true\n")
	yCauchy := EulerCauchyMethod(x, y, dy)
	for i := range yCauchy {
		fmt.Printf("%2d | %.1f | %.9f  | %.9f \n", i, x[i], yCauchy[i], y[i])
	}

	fmt.Println("Runge Kutta Method")
	fmt.Printf(" k | x   |\ty \t| y_true\n")
	yRunge := RungeKuttaMethod(x, y, dy)
	for _, val := range yRunge {
		fmt.Printf("y_k %.9f\n", val)
	}

	fmt.Println("Multistep Adams Method")
	fmt.Printf(" k | x   |\ty \t| y_true\n")
	xAdam, yAdam, dyAdam := getValues(INTERVAL_ADAMS)
	yNewAdam := MultistepAdamsMethod(xAdam, yAdam, dyAdam)
	for i := range yNewAdam {
		fmt.Printf("%2d | %.1f | %.9f  | %.9f \n", i, xAdam[i], yNewAdam[i], yAdam[i])
	}

	defStyle := chart.Style{StrokeWidth: 5}

	series := []chart.Series{
		chart.ContinuousSeries{Style: defStyle, XValues: x, YValues: yEuler, Name: "Euler"},
		chart.ContinuousSeries{Style: defStyle, XValues: x, YValues: yCauchy, Name: "Euler Cauchy"},
		chart.ContinuousSeries{Style: defStyle, XValues: x, YValues: yRunge, Name: "Runge Kutta"},
		chart.ContinuousSeries{Style: defStyle, XValues: xAdam[:len(x)], YValues: yNewAdam[:len(x)], Name: "Adam"},
	}

	PlotGraphs(series)
}
