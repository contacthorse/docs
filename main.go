package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const googleDocURL = "https://docs.google.com/document/d/e/2PACX-1vSsmM8yaTYx5tQst5tOLT85r7JjvEb0G-abrOZiblYvG2Eymkp4qZYjoSFr9_qemrRy44JNklGgmv96/pub"

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, googleDocURL, http.StatusTemporaryRedirect)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", redirectHandler)

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
