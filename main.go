package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const N = 9

func isSafe(board [N][N]int, row, col, num int) bool {
	for x := 0; x < N; x++ {
		if board[row][x] == num || board[x][col] == num ||
			board[row-row%3+x/3][col-col%3+x%3] == num {
			return false
		}
	}
	return true
}


func solveSudoku(board *[N][N]int) bool {
	row, col := -1, -1
	found := false

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 0 {
				row, col = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isSafe(*board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

type SudokuRequest struct {
	Board [N][N]int `json:"board"`
}


func sudokuHandler(c *gin.Context) {
	var request SudokuRequest

	
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Sudoku board"})
		return
	}

	
	board := request.Board
	fmt.Println("Received board:", board) 

	if solveSudoku(&board) {
		c.JSON(http.StatusOK, gin.H{"board": board}) 
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No solution exists"})
	}
}

func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	
	r.POST("/solve", sudokuHandler)

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	r.Run(":1010")
}
