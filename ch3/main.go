package main

import (
	"fmt"
	"strconv"
)

func main()  {
	x, err := strconv.Atoi("123")             // x is an int
	if err != nil {
		fmt.Println(err)
	}
	y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(x, y)
}
