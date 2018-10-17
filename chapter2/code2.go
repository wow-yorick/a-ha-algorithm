package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputreader := bufio.NewReader(os.Stdin)
	fmt.Println("please input something:")
	input, err := inputreader.ReadString('\n')
	if err != nil {
		fmt.Println("error")
	}
	input = strings.Replace(input, "\n", "", -1)

	var s []byte
	len := len(input)
	mid := len/2 - 1
	top := 0
	for i := 0; i < mid; i++ {
		s = append(s, input[i])
	}
	top = bytes.Count(s, nil) - 1
	var next int
	if len%2 == 0 {
		next = mid + 1
	} else {
		next = mid + 2
	}

	for i := next; i <= len-1; i++ {
		if input[i] != s[top] {
			break
		}
		top--
	}
	if top == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
