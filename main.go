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

	ifaces, err := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
					ip = v.IP
			case *net.IPAddr:
					ip = v.IP
			}
			fmt.Fprintf(w, "IP(s): %s\n", hostname)
		}
	}
}
