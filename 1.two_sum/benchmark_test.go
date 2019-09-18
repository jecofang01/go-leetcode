package twosum

import (
	"testing"
)

func Benchmark_TwoSum(b *testing.B) {
	calculators := []struct {
		name       string
		calculator Calculator
	}{
		{"brute force", &bruteForceCalculator{}},
		{"hash map", &mapCalculator{}},
	}
	for _, c := range calculators {
		c := c
		b.Run(c.name, func(b *testing.B) {
			for _, tt := range testData {
				tt := tt
				b.Run("", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						c.calculator.TwoSum(tt.nums, tt.target)
					}
				})
			}
		})
	}
}
