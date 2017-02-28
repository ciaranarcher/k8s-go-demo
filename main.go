package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func movie(w http.ResponseWriter, r *http.Request) {
	movie := "La La Land"
	t := time.Now()
	fmt.Println(t.Format("20060102150405"), movie)
	io.WriteString(w, movie)
}

func main() {
	http.HandleFunc("/", movie)
	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(":8000", nil)
}
