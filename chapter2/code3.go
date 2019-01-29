package main

import "fmt"

func main() {
	var a, b, c, m, sum int = 0, 0, 0, 0, 0
	m = 18
	for a = 0; a <= 1111; a++ {
		for b = 0; b <= 1111; b++ {
			c = a + b
			if result(a)+result(b)+result(c) == m-4 {
				fmt.Printf("%d+%d=%d", a, b, c)
				sum++
			}
		}
	}
	fmt.Printf("一共可以拼出%d个不同的等式", sum)
}

func result(x int) int {
	num := int(0)
	f := []int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	for x/10 != 0 {
		num += f[x%10]
		x = x / int(10)
	}
	num += f[x]
	return num
}
