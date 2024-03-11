package main

import (
	"fmt"
	"net/http"
	// "time"
)

var i = 0
var memory = make(map[int][][]int)

func start(w http.ResponseWriter, req *http.Request) {

	// ctx := req.Context()
	fmt.Println("[INFO]: connection received")
	defer fmt.Println("[INFO]: responded")

	data := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	memory[i] = data
	i += 1

	fmt.Println("memory:", memory)

	// select {
	// case <-time.After(10 * time.Second):
	// 	fmt.Fprintf(w, "hello\n")
	// case <-ctx.Done():
	// 	err := ctx.Err()
	// 	fmt.Println("server:", err)
	// 	internalError := http.StatusInternalServerError
	// 	http.Error(w, err.Error(), internalError)
	// }
}

func main() {

	http.HandleFunc("/start", start)
	http.ListenAndServe(":8090", nil)
}
