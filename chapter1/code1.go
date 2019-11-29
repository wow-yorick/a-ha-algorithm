package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		a    [11]int
		i, j int
	)
	num := os.Args
	if len(num) <= 1 || len(num) >= 7 {
		log.Fatal("请输入五个数字")
	}
	num = num[1:6]
	log.Println(num)
	for i = 0; i < len(num); i++ {
		fmt.Println("arg:", num[i])
		ind, _ := strconv.ParseInt(num[i], 10, 64)
		a[ind]++
		fmt.Println("a[ind]:", a[ind])
	}

	log.Println(a)
	for i = 0; i < len(a); i++ {
		for j = 1; j <= a[i]; j++ {
			fmt.Printf("%d", i)
		}
	}
}
