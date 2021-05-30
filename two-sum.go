/*
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/

package main

import "fmt" 


func main () {
	fmt.Println("This is two sum.")

	// first input
	numbers := []int{ 2, 7, 11, 15 }

	for i, num := range numbers {
		fmt.Println(i, ":", num)
	}

	// another input
	two_sum := 25

	fmt.Println("=-----=")
	fmt.Println(runSolution(numbers, two_sum))
}


func runSolution(col []int, twoSum int) (otherIndex int, thisIndex int) {
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