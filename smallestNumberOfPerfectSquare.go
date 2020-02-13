package main

import (
	"fmt"
	"math"
)

/*
Write a program that determines the smallest number of perfect squares that sum up to N.

Here are a few examples:

Given N = 4, return 1 (4)
Given N = 17, return 2 (16 + 1)
Given N = 18, return 2 (9 + 9)
*/

// Dynamic Programming Solution
func smallestNumberOfPerfectSqaure(num int) int {
	dp := make([]int, num+1)

	// base case
	for i := 0; i < 4; i++ {
		dp[i] = i
	}

	for i := 4; i <= num; i++ {
		dp[i] = i

		for j := 1; j <= int(math.Ceil(math.Sqrt(float64(i)))); j++ {
			if temp := j * j; temp > i {
				break
			} else {
				dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-temp])))
			}
		}
	}
	return dp[num]
}

// Brute Force solution
func smallestNumberOfPerfectSqaure2(num int) int {
	floatnum := float64(num)
	sqrt := math.Sqrt(floatnum)

	if sqrtToInt := math.Trunc(sqrt); sqrtToInt == sqrt {
		return 1
	}
	if num < 4 {
		return num
	}

	result := floatnum
	for i := 1; i <= num; i++ {
		temp := i * i
		if temp > num {
			break
		} else {
			result = math.Min(result, float64(smallestNumberOfPerfectSqaure2(num-temp)+1))
		}
	}
	return int(result)
}

func main() {
	n1 := 4
	fmt.Println(smallestNumberOfPerfectSqaure(n1))
	n2 := 17
	fmt.Println(smallestNumberOfPerfectSqaure(n2))
	n3 := 18
	fmt.Println(smallestNumberOfPerfectSqaure2(n3))
	n4 := 22
	fmt.Println(smallestNumberOfPerfectSqaure(n4))
}
