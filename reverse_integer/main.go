// https://leetcode.com/problems/reverse-integer/description

package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	var index int
	if x < 0 {
		index++
		x = x * (-1)
	}
	str := strconv.Itoa(x)
	var reverse []byte
	for i := len(str) - 1; i >= 0; i-- {
		reverse = append(reverse, str[i])
	}

	if index == 1 {
		x, _ = strconv.Atoi(string(reverse))
		x = x * (-1)
	} else {
		x, _ = strconv.Atoi(string(reverse))
	}
	if x <= -2147483648 || x >= 2147483647 {
		return 0
	}
	return x
}

func main() {

	fmt.Println(reverse(-123))

}
