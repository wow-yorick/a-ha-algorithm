package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input some num:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("there ware errors reading", err)
		return
	}
	fmt.Printf("your num is %s", input)
	nums := strings.Split(input, " ")
	fmt.Printf("num len %d ", len(nums))
	numArr := make([]int, 5)
	for v := range nums {
		numArr = append(numArr, v)
	}
}
