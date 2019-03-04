package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header  {
		fmt.Fprintf(w,"header[%q] = %s\n", k, v)
	}
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	if err := r.ParseForm(); err != nil {
		log.Printf("%s", err)
	}
	for k, v := range r.Form{
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
