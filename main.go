package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", printHostname)
	http.ListenAndServe(":8080", nil)
}

func printHostname(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Your hostname of host node is %s!", name)
}
