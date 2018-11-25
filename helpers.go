package drow

func Sum(rolls []int) int {
	sum := 0
	for _, addend := range rolls {
		sum += addend
	}
	return sum
}
