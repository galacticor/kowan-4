package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr: "0.0.0.0:" + port,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World, Widy"))
	})

	log.Printf("Server running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error starting server: ", err)
		panic(err)
	}
}