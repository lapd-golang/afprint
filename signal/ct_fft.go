package signal

import (
	"math"
	"math/cmplx"
)

// Returns the fast fourier transformation of x. If the length of x is not a power of 2 returns []
func FFT(x []float32) []complex128 {
	if !isPowerof2(len(x)) {
		return []complex128{}
	}

	y := toComplex(x)
	ctFFT(x, y, len(x), 1)
	return y
}

// Implementation of the recursive Cooley–Tukey algorithm
// Results are returned in the preallocated slice passed as a parameter
func ctFFT(x []float32, y []complex128, n, s int) {
	if n == 1 {
		y[0] = complex(float64(x[0]), 0)
		return
	}
	ctFFT(x, y, n/2, 2*s)
	ctFFT(x[s:], y[n/2:], n/2, 2*s)
	for k := 0; k < n/2; k++ {
		tf := cmplx.Rect(1, -2*math.Pi*float64(k)/float64(n)) * y[k+n/2]
		y[k], y[k+n/2] = y[k]+tf, y[k]-tf
	}
}

// Returns a []complex128 representation of the slice
func toComplex(x []float32) []complex128 {
	y := make([]complex128, len(x))
	for n, v := range x {
		y[n] = complex(float64(v), 0)
	}
	return y
}

// Check if n is a power of 2
func isPowerof2(n int) bool {
	return n&(n-1) == 0
}
