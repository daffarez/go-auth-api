package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func setupRouter() http.Handler {
	muxServer := http.NewServeMux()
	muxServer.HandleFunc("/health", healthHandler)
	return muxServer
}

func main() {
	muxServer := setupRouter()

	log.Println("Server running on :8088")
	err := http.ListenAndServe(":8088", muxServer)
	if err != nil {
		log.Fatal(err)
	}
}
