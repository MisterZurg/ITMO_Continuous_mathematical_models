package main

import (
	"fmt"
	"github.com/wcharczuk/go-chart/v2"
	"math"
	"os"
)

// o get gonum.org/v1/plot/...
/*
Solve the ODE numerically using the four given methods
(Lecture 4, slides 13, 17, 24, 28). Attach a Python script and one graph of
solutions for all methods along with the exact solution specified in the
problem statement.
*/

var INTERVAL = [2]float64{0, 0.5}
var INTERVAL_ADAMS = [2]float64{0, 1}

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
// TODO check yhats incorrect
func EulerCauchyMethod(xs, ys, yds []float64) []float64 {
	n := len(xs)

	// TODO L4 Slide 16 check second line
	yhats := make([]float64, n)
	yhats[0] = 0 // DOENST EXIST
	for i := 1; i < n; i++ {
		yhats[i] = ys[i-1] + STEP*yds[i-1]
	}
	// fmt.Println(yhats)

	yecm := make([]float64, n)
	yecm[0] = ys[0]
	for i := 1; i < n; i++ {
		hatDirr := FirstOrderDerivative(xs[i], yhats[i])
		fraction := (yds[i-1] + hatDirr) / 2
		yecm[i] = yecm[i-1] + STEP*fraction
	}
	return yecm
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

// MultistepAdamsMethod
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

func main() {
	xs, ys, yds := getValues(INTERVAL)
	yme := EulerMethod(xs, ys, yds)
	fmt.Println("EulerMethod")
	fmt.Printf(" k | x   |\ty \t| y_true\n")
	for i := range yme {
		fmt.Printf("%2d | %.1f | %.9f  | %.9f \n", i, xs[i], yme[i], ys[i])
	}

	fmt.Println("EulerCauchyMethod")
	fmt.Printf(" k | x   |\ty \t| y_true\n")
	ymce := EulerCauchyMethod(xs, ys, yds)
	for i := range ymce {
		fmt.Printf("%2d | %.1f | %.9f  | %.9f \n", i, xs[i], ymce[i], ys[i])
	}

	fmt.Println("RungeKuttaMethod")
	fmt.Printf(" k | x   |\ty \t| y_true\n")
	yrkts := RungeKuttaMethod(xs, ys, yds)
	for _, val := range yrkts {
		fmt.Printf("y_k %.9f\n", val)
	}

	fmt.Println("MultistepAdamsMethod")
	fmt.Printf(" k | x   |\ty \t| y_true\n")
	xas, yas, yads := getValues(INTERVAL_ADAMS)
	yadams := MultistepAdamsMethod(xas, yas, yads)
	for i := range yadams {
		fmt.Printf("%2d | %.1f | %.9f  | %.9f \n", i, xas[i], yadams[i], yas[i])
	}

	series := []chart.Series{
		chart.ContinuousSeries{XValues: xs, YValues: yme},
		chart.ContinuousSeries{XValues: xs, YValues: ymce},
		chart.ContinuousSeries{XValues: xs, YValues: yrkts},
		chart.ContinuousSeries{XValues: xas, YValues: yadams},
	}

	PlotGraphs(series)
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

func Answer(x, yTrue float64) string {
	return fmt.Sprintf("%.2f\t| %.9f", x, yTrue)
}

func PlotGraphs(series []chart.Series) {
	graph := chart.Chart{
		Series: series,
	}
	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
