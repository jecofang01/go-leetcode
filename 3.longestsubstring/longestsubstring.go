package longestsubstring

func LengthOfLongestSubstring(s string) int {
	r, b := 0, 0
	m := map[rune]int{}
	for i, v := range s {
		l := i - b + 1
		if idx, ok := m[v]; ok && idx >= b {
			l = i - idx
			b = idx + 1
		}
		m[v] = i
		if l > r {
			r = l
		}
	}
	return r
}
