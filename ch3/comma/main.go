package main

import "fmt"

func main()  {
	s := "123456789"
	t := comma(s)
	fmt.Println(t)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
