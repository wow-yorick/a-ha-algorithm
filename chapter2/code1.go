package main

import "fmt"

func main() {
	q := []int{0, 6, 3, 1, 7, 5, 8, 9, 2, 4}
	m := make([]int, 0)
	var head, tail int
	head = 1
	tail = 10
	for head < tail {
		fmt.Println(head)
		//fmt.Println(q[head])
		head++
		if head < tail {
			m = append(m, q[head])
		}
		head++
	}

	fmt.Println("result:", m)
}
