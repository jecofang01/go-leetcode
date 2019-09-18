package twosum

type Calculator interface {
	TwoSum(nums []int, target int) []int
}

type bruteForceCalculator struct{}

func (c bruteForceCalculator) TwoSum(nums []int, target int) []int {
	for i, a := range nums {
		for j, b := range nums {
			if a+b == target && i != j {
				return []int{i, j}
			}
		}
	}
	return nil
}

type mapCalculator struct{}

func (c mapCalculator) TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
