package main

import (
	"fmt"
	"math"
)

/*
Solve the ODE numerically using the four given methods
(Lecture 4, slides 13, 17, 24, 28). Attach a Python script and one graph of
solutions for all methods along with the exact solution specified in the
problem statement.
*/

var INTERVAL = [2]float64{0, 0.5}

const STEP float64 = 0.1

//var n = int(INTERVAL[1] / STEP)
//var ys = make([]float64, n)

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

func EulerCauchyMethod(interval [2]float64, h float64) {

}

// slide 24
func RungeKuttaMethod(interval [2]float64, h float64) {

}

func AdamsMethod(interval [2]float64, h float64) {

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
	xs, ys, yds := getValues()
	yme := EulerMethod(xs, ys, yds)
	//fmt.Printf("x \t| y_true\n")
	for _, val := range yme {
		fmt.Println(val)
	}
}

func getValues() ([]float64, []float64, []float64) {
	var xs, ys, yds []float64

	for x := INTERVAL[0]; x < INTERVAL[1]; x += STEP {
		xs = append(xs, x)
		ys = append(ys, Equation(x))
		yds = append(yds, FirstOrderDerivative(x, ys[len(ys)-1]))
	}
	return xs, ys, yds
}

func Answer(x, yTrue float64) string {
	return fmt.Sprintf("%.2f\t| %.9f", x, yTrue)
}
