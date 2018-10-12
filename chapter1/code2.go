package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	input = strings.Replace(input, "\n", "", -1)
	nums := strings.Split(input, "|")
	fmt.Printf("num len %d, %v ", len(nums), nums)
	var numArr []int
	for k, v := range nums {
		iv, _ := strconv.Atoi(v)
		numArr = append(numArr, iv)
		fmt.Println(k, iv)
	}

	for i := 0; i <= len(numArr)-1; i++ {
		for j := 0; j < len(numArr)-1; j++ {
			if numArr[j] < numArr[j+1] {
				numArr[j], numArr[j+1] = numArr[j+1], numArr[j]
			}
		}
	}

	// for i := 1; i <= len(numArr); i++ {
	// 	fmt.Printf("%d ", numArr[i])
	// }
	fmt.Println(numArr)
}
