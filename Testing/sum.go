package Testing

func sum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}
