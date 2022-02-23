package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Requested Host: %s\n", r.Host)
	fmt.Fprintf(w, "used path: %s\n", r.URL.Path[0:])

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "Hostname: %s\n", hostname)
	}
}
