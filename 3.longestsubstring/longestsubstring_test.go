package longestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  int
	}{
		{"abcabcbb", "abcabcbb", 3},
		{"bbbbb", "bbbbb", 1},
		{"pwwkew", "pwwkew", 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := LengthOfLongestSubstring(tt.in)

			assert.Equal(t, tt.out, actual)
		})
	}
}
