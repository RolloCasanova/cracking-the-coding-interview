package main

import (
	"fmt"
	"os"

	"github.com/RolloCasanova/cracking-the-coding-interview/1_arrays_and_strings/utils"
)

func main() {
	if len(os.Args) < 2 {
		panic("usage: go run main.go <row1> <row2> ... <rowN>")
	}

	rows := len(os.Args[1:])
	columns := len(os.Args[1])
	matrix := make([][]string, rows)

	for i := 0; i < rows; i++ {
		if len(os.Args[i+1]) != columns {
			panic("all rows must have the same number of columns")
		}

		matrix[i] = make([]string, columns)
		for j := 0; j < columns; j++ {
			if len(os.Args[i+1]) != columns {
				panic("all rows must have the same number of columns")
			}

			if os.Args[i+1][j] != '0' && os.Args[i+1][j] != '1' {
				panic("all elements must be 0 or 1")
			}

			matrix[i][j] = string(os.Args[i+1][j])
		}
	}

	fmt.Println("Input matrix:")
	utils.PrintStringMatrix(matrix, 1)

	zeroedMatrix := zeroMatrix(matrix)

	fmt.Println("Zeroed matrix:")
	utils.PrintStringMatrix(zeroedMatrix, 1)
}

// zeroMatrix zeroes all rows and columns in a matrix that contain a "0"; this is done in place
func zeroMatrix(matrix [][]string) [][]string {
	// using 0th row and 0th column to indicate if the column contain a "0" in the matrix
	var firstRowHasZero, firstColHasZero bool

	// validate if any element in the 0th column has a "0"
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == "0" {
			firstColHasZero = true
			break
		}
	}

	// validate if any element in the 0th row has a "0"
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == "0" {
			firstRowHasZero = true
			break
		}
	}

	// omitting 0th row and column, check if there's a "0" in the matrix
	// use i-th position of 0th column to indicate if there's a zero in the i-th row
	// use j-th position of 0th row to indicate if there's a zero in the j-th column
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == "0" {
				matrix[0][j] = "0"
				matrix[i][0] = "0"
			}
		}
	}

	// traverse 0-th column and make all rows to "0" in matrix
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == "0" {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = "0"
			}
		}
	}

	// traverse 0-th row and make all columns to "0" in matrix
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == "0" {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = "0"
			}
		}
	}

	// if first column had a "0" make all elements in 0th column to "0"
	if firstColHasZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = "0"
		}
	}

	// if 0th row had a "0" make all elements in 0th row to "0"
	if firstRowHasZero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = "0"
		}
	}

	return matrix
}
