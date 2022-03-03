package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		z := 1.0
		for math.Abs(z*z-x) >= 0.000001 {
			z -= (z*z - x) / (2 * z)
		}
		return z, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}
