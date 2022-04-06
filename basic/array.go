package basic

import (
	"sort"
)

// Array 基础 - 数组
type Array struct{}

// NewBasicArray 创建结构体实例
func NewBasicArray() *Array {
	return &Array{}
}

// TwoSum 两数之和 - 给定一个整数数组 nums 和一个整数目标值 target，找出和为目标值 target 的那两个整数，返回它们的数组下标，链接：https://leetcode-cn.com/problems/two-sum/
// nums 整数数组
// target 目标值
// indexs 下标
func (array *Array) TwoSum(nums []int, target int) (indexs []int) {
	// 可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。------ 说实话，这个但是没看懂，是指数组内的元素不能重复吗？

	// 解法一：遍历 nums，当前数与下一个数相加，需要循环 n(n-1)/2 + n - 1 次，执行结果：执行用时 28ms，内存消耗 3.4MB
	// length := len(nums)
	// sum := 0
	// for i := 0; i < (length - 1); i++ {
	// 	for j := i + 1; j <= (length - 1); j++ {
	// 		sum = nums[i] + nums[j]
	// 		if sum == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }

	// 解法二：遍历 nums，反向找符合数据的另一个值，执行结果：执行用时 4ms，内存消耗 4.1MB
	length := len(nums)
	indexMap := make(map[int]int)
	for i := 0; i < length; i++ {
		num := nums[i]
		if v, isExist := indexMap[target-num]; isExist {
			indexs = []int{v, i}
			return
		}
		indexMap[num] = i
	}
	return
}

// FindMedianSortedArrays 寻找两个正序数组的中位数 - 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。找出并返回这两个正序数组的中位数。算法的时间复杂度应该为 O(log (m+n))，链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
// nums1 数组 1
// nums2 数组 2
// result 中位数
func (array *Array) FindMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
	// 合并数组，重新排序，计算中位数

	// 解法一：判断两个数组谁大谁小，假如 m > n，则遍历 m。取 n 的第一个值与遍历 m 的当前值进行比较，若此值大于当前值，
	// largeArray := nums1
	// smallArray := nums2
	// if len(nums1) < len(nums2) {
	// 	largeArray = nums2
	// 	smallArray = nums1
	// }

	// largeArrayLength := len(largeArray)
	// smallArrayLength := len(smallArray)

	// mergeArray := []int{}

	// if largeArrayLength == 0 || smallArrayLength == 0 {
	// 	mergeArray = largeArray
	// } else {
	// 	// 如果最小的比另一个最大的还大或者最大的比另一个最小的还小，则直接合并
	// 	if largeArray[0] >= smallArray[smallArrayLength-1] {
	// 		mergeArray = append(smallArray, largeArray...)
	// 	} else if largeArray[largeArrayLength-1] <= smallArray[0] {
	// 		mergeArray = append(largeArray, smallArray...)
	// 	} else {
	// 		smallArrayNextIndex := 0
	// 		for i := 0; i < largeArrayLength; i++ {
	// 			// 判断当前值与 smallArrayNextIndex 对应的值进行比较
	// 			numForLarge := largeArray[i]
	// 			if smallArrayLength != smallArrayNextIndex {
	// 				// 因为中间还可能包含其他的，所以，还要判断一下
	// 				for {
	// 					numForSmall := smallArray[smallArrayNextIndex]
	// 					if numForLarge < numForSmall {
	// 						break
	// 					}
	// 					mergeArray = append(mergeArray, numForSmall)
	// 					smallArrayNextIndex++
	// 					if smallArrayLength == smallArrayNextIndex {
	// 						break
	// 					}
	// 				}
	// 			}
	// 			mergeArray = append(mergeArray, numForLarge)
	// 		}
	// 		// 如果小数组的长度减 1 后并不等于下标，则说明还没全部合并进来
	// 		if smallArrayLength != smallArrayNextIndex {
	// 			mergeArray = append(mergeArray, smallArray[smallArrayNextIndex:]...)
	// 		}
	// 	}
	// }

	// 解法二：适用无序的数组找中位数，通过 sort 包进行排序
	mergeArray := append(nums1, nums2...)
	sort.Ints(mergeArray)
	totalLength := len(mergeArray)
	// 判断是否是偶数
	if (totalLength)%2 == 0 {
		result = (float64(mergeArray[totalLength/2-1]) + float64(mergeArray[totalLength/2])) / 2
	} else {
		result = float64(mergeArray[(totalLength+1)/2-1])
	}

	return
}
