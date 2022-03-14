package equations

import (
	"fmt"
	"math"
)

const (
	// accuracy level
	epsilon = 1e-5
	maxUint = ^uint(0)
	minUint = 0
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

func solveSquareEquation(a, b, c float64) (result []float64, err error) {
	if inEpsilon(a, 0) {
		err = fmt.Errorf("not a square equation")
		return
	}

	d := b*b - 4*a*c
	if d < 0 {
		return
	}

	if inEpsilon(d, 0) {
		x := -b / (2 * a)
		result = append(result, x, x)
		return
	}

	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)
	result = append(result, x1, x2)
	return
}

func inEpsilon(a, b float64) bool {
	return math.Abs(a-b) <= epsilon
}
