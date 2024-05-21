package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

func main() {
	Connect()
}

func Connect() {
	addr := os.Getenv("ADDR")

	if len(addr) == 0 {
		addr = ":3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	log.Printf("Server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
	fmt.Println("Connection...")
}
