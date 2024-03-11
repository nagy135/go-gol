package main

import (
	"fmt"
	"net/http"
	"strconv"
	// "time"
)

const width = 10
const height = 10

var i = 0
var memory = make(map[int][][]int)

func iterate(board [][]int) [][]int {
	indexes := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			alive := 0
			for _, pair := range indexes {
				targetX := x + pair[0]
				targetY := y + pair[1]
				if targetX > 0 && targetY > 0 && targetX < width && targetY < height && board[targetY][targetX] == 1 {
					alive += 1
				}
			}

			// rules
			if board[y][x] == 1 && (alive == 2 || alive == 3) {
				board[y][x] = 1
			} else if board[y][x] == 0 && alive == 3 {
				board[y][x] = 1
			} else {
				board[y][x] = 0
			}
		}
	}
	return board
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
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		memory[idInt] = data
		i += 1
		board = memory[idInt]

	}
	fmt.Println("[INFO]: start: ", board)
	board = iterate(board)
	fmt.Println("[INFO]: after: ", board)
}

func main() {

	http.HandleFunc("/tick", tick)
	http.ListenAndServe(":8090", nil)
}
