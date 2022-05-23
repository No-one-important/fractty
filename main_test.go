package main

import (
	"fmt"
	"testing"
)

func TestIsConvergentMandelbrot(t *testing.T) {
	var tests = []struct {
        a, b float64
        want int
    }{
        {0, 0, 64},
        {1, 0, 1},
        {2, -2, 0},
        {0, -1, 64},
        {-1, 0, 64},
		{-0.7433183529628573, -0.11102957901891086, 58},
    }

	for _, tt := range tests {
        testname := fmt.Sprintf("%f,%f", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            _, ans := isConvergentMandelbrot(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %v, want %v", ans, tt.want)
            }
        })
    }
}

func BenchmarkIsConvergentMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
        isConvergentMandelbrot(-0.7433183529628573, -0.11102957901891086)
    }
}