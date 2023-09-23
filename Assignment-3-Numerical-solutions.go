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

// EulerMethod is implementation from slide 13
func EulerMethod(xi, yi float64) float64 {
	if xi == INTERVAL[0] {
		return Equation(xi)
	}
	return 0 // EulerMethod(yi-1) + STEP
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
func FirstOrderDifferentialEquation(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	sum := x + y
	return math.Pow(sum, 2)
}

func main() {
	fmt.Printf("x \t| y_true\n")
	for x := INTERVAL[0]; x < INTERVAL[1]; x += STEP {
		fmt.Println(Answer(x))
	}
}

func Answer(x float64) string {
	return fmt.Sprintf("%.2f\t| %.9f", x, Equation(x))
}
