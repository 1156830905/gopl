package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:]  {
		reps, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(reps.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
		}
		fmt.Printf("%s", b)
	}
}
