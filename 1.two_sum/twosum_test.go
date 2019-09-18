package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	nums     []int
	target   int
	expected []int
}{
	{
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},
	{
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	},
	{
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	},
	{
		nums:     []int{0},
		target:   6,
		expected: nil,
	},
	{
		nums:     nil,
		target:   6,
		expected: nil,
	},
}

func Test_TwoSum(t *testing.T) {
	calculators := []struct {
		name       string
		calculator Calculator
	}{
		{"brute force", &bruteForceCalculator{}},
		{"hash map", &mapCalculator{}},
	}
	for _, c := range calculators {
		c := c
		t.Run(c.name, func(t *testing.T) {
			for _, tt := range testData {
				tt := tt
				t.Run("", func(t *testing.T) {
					actual := c.calculator.TwoSum(tt.nums, tt.target)
					assert.Equal(t, tt.expected, actual)
				})
			}
		})
	}
}
