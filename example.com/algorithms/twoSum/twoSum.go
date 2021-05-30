/*
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/

package twoSum


func runSolutionWithMap(col []int, twoSum int) (otherIndex int, thisIndex int) {
	map1 := make(map[int]int)
	for thisIndex, num := range col {
		map1[num] = thisIndex
		otherNum := twoSum - num
		otherIndex, ok := map1[otherNum]
		if ok {
			return otherIndex, thisIndex
		}
	}
	return -1, -1
}


/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.

这种方法的好处是，基于已有的条件，就是数组已经是排过序的。

*/
func runSolutionWithTwoPointers(col []int, twoSum int) (otherIndex int, thisIndex int) {
	// first and last indexes
	l, r := 0, len(col) - 1

	currentSum := 0
	for {
		currentSum = col[l] + col[r]
		if l >= r {
			break
		}
		if currentSum < twoSum {
			l++
		} else if currentSum > twoSum {
			r--
		} else {
			return l, r
		}
	}
	return -1, -1
}


