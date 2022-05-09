package main

import (
	"fmt"
	"testing"
)

func TestIsConvergent(t *testing.T) {
	var tests = []struct {
        a, b float64
        want int
    }{
        {0, 0, 64},
        {1, 0, 2},
        {2, -2, 1},
        {0, -1, 64},
        {-1, 0, 64},
		{-0.7433183529628573, -0.11102957901891086, 59},
    }

	for _, tt := range tests {
        testname := fmt.Sprintf("%f,%f", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            _, ans := isConvergent(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %v, want %v", ans, tt.want)
            }
        })
    }
}

func BenchmarkIsConvergent(b *testing.B) {
	for i := 0; i < b.N; i++ {
        isConvergent(-0.7433183529628573, -0.11102957901891086)
    }
}