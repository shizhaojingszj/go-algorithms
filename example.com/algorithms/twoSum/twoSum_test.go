package twoSum

import (
	"fmt"
	"testing"
)


func TestRunSolutionWithMap(t *testing.T) {
	fmt.Println("This is two sum.")

	// first input
	numbers := []int{ 2, 7, 11, 15 }

	for i, num := range numbers {
		fmt.Println(i, ":", num)
	}

	// another input
	two_sum := 18

	fmt.Println(runSolutionWithMap(numbers, two_sum))

}


func TestRunSolutionWithTwoPointers(t *testing.T) {
	fmt.Println("This is two sum.")

	// first input
	numbers := []int{ 2, 7, 11, 15 }

	for i, num := range numbers {
		fmt.Println(i, ":", num)
	}

	// another input
	two_sum := 18

	fmt.Println(runSolutionWithTwoPointers(numbers, two_sum))

}


func BenchmarkTwoSumRunSolutionWithMap(b *testing.B) {
	b.ResetTimer()
	col := []int{2, 7, 11, 15}
	for n := 0; n < b.N; n++ {
		runSolutionWithMap(col, 18)
	}
	b.StartTimer()
}


func BenchmarkTwoSumRunSolutionWithTwoPointers(b *testing.B) {
	b.ResetTimer()
	col := []int{2, 7, 11, 15}
	for n := 0; n < b.N; n++ {
		runSolutionWithTwoPointers(col, 18)		
	}
	b.StartTimer()
}