package main

import (
	"fmt"
	"net/http"
	"strconv"
	// "time"
)

const width = 3
const height = 3

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

func tick(w http.ResponseWriter, req *http.Request) {
	fmt.Println("[INFO]: connection received")
	defer fmt.Println("[INFO]: responded")
	id := req.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	board, ok := memory[idInt]
	if !ok {
		fmt.Println("[INFO]: creating new")
		data := [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}
		memory[idInt] = data
		i += 1
		board = memory[idInt]

	}
	fmt.Println("[INFO]: start: ", board)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			board[y][x] = 1
		}
	}
	fmt.Println("[INFO]: after: ", board)
}

func main() {

	http.HandleFunc("/start", start)
	http.HandleFunc("/tick", tick)
	http.ListenAndServe(":8090", nil)
}
