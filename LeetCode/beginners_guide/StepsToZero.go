package main

// 1342. Number of Steps to Reduce a Number to Zero
func ToZero(num int) int {
	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num -= 1
		}
		steps++
	}
	return steps
}
