package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := getPort()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Printf("Listening on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func getPort() int {
	port := 8080
	if portEnv, portExist := os.LookupEnv("PORT"); portExist {
		var err error
		port, err = strconv.Atoi(portEnv)
		if err != nil {
			log.Printf("Failed to parse port : %v", err)
			port = 8080
		}
	}

	return port
}
