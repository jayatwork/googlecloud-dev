package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//indexhandler will reesponse to req with our lil shoutout
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {

	}
	fmt.Fprint(w, "Hello, World a GOGO")
}

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
