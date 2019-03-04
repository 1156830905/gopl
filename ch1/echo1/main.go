package main

import (
	"fmt"
	"os"
)

func main()  {
	var s, sep string
	l := len(os.Args)
	for i := 0; i < l; i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s);
}
