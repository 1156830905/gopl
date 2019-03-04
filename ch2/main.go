package main

import (
	"log"
	"os"
)

var cwd string

func main()  {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed %v\n", err)
	}
	log.Printf("hello")
	log.Printf("Working directory = %s\n", cwd)
}
