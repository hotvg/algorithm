package basic

import (
	"fmt"
	"testing"

	"haolele.cn/algorithm/basic"
)

func TestTwoSum(t *testing.T) {
	array := basic.NewBasicArray()
	indexs := array.TwoSum([]int{1, 2, 5, 4}, 3)
	fmt.Printf("TestTwoSum result: %v\n", indexs)
}

func TestFindMedianSortedArrays(t *testing.T) {
	array := basic.NewBasicArray()
	result := array.FindMedianSortedArrays([]int{1, 4, 6, 7}, []int{2, 3, 5})
	fmt.Printf("FindMedianSortedArrays result: %v\n", result)
}
