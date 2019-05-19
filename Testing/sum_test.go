package Testing

import (
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1,2,3,4,5}
	expectedSum := 15
	s := sum(nums)

	if s != expectedSum {
		t.Errorf("The Sum should be %v but got %v",expectedSum,s)
	}
}
