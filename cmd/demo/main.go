package main

import (
	"fmt"
	"math"

	"github.com/mrumyantsev/randmatrix"
)

func main() {
	cols := 3
	rows := 10
	randCap := math.MaxInt

	mat, err := randmatrix.RandMatrix(cols, rows, randCap)
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		return
	}

	printMatrix(mat, cols, rows)
}

func printMatrix(mat [][]int, cols int, rows int) {
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			fmt.Printf("%23d", mat[i][j])
		}

		fmt.Println()
	}
}
