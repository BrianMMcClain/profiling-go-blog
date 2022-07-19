package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(n int) string {
	ret := ""
	if n%3 == 0 {
		ret += "Fizz"
	}
	if n%5 == 0 {
		ret += "Buzz"
	}
	if len(ret) == 0 {
		return strconv.Itoa(n)
	}
	return ret
}
