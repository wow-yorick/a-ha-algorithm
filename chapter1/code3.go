package main

import "fmt"

var (
	a []int
	n int
)

func quicksort(left, right int) {
	var i, j, temp int
	if left > right {
		return
	}
	temp = a[left]
	i = left
	j = right
	for i != j {
		for a[j] >= temp && i < j {
			j--
		}
		for a[i] <= temp && i < j {
			i++
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	a[left] = a[i]
	a[i] = temp
	quicksort(left, i-1)
	quicksort(i+1, right)
}

func main() {
	a = []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}

	quicksort(0, len(a)-1)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
