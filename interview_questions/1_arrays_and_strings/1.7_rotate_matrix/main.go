package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// get the matrix as a slice of slices from stdin
	// first parameter is the number of rows
	// following parameters are the rows
	// each row is a string of space separated integers
	//
	// example:
	// 3 "1 2 3" "4 5 6" "7 8 9"
	// 4 "1 2 3 4" "5 6 7 8" "9 10 11 12" "13 14 15 16"
	// 5 "1 2 3 4 5" "6 7 8 9 10" "11 12 13 14 15" "16 17 18 19 20" "21 22 23 24 25"
	// 6 "1 2 3 4 5 6" "7 8 9 10 11 12" "13 14 15 16 17 18" "19 20 21 22 23 24" "25 26 27 28 29 30" "31 32 33 34 35 36"
	//
	// program execution will fail if:
	// - the number of rows is not provided
	// - the number of rows provided does not match the number of rows provided
	// - any of the rows provided does not contain the same number of elements as the number of rows provided
	//
	// for better display of matrix at most two characters are allowed per element
	// it will work for any number of characters per element, but the matrix will not be displayed properly

	if len(os.Args) < 3 {
		panic("usage: go run main.go <size (N)> <row1> <row2> ... <rowN>")
	}

	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("size should be an integer")
	}

	if len(os.Args) != size+2 {
		panic("number of rows provided does not match the number of rows provided")
	}

	matrix := make([][]string, size)
	for i, row := range os.Args[2:] {
		// separate each row by space
		matrix[i] = strings.Split(row, " ")

		if len(matrix[i]) != size {
			panic("number of elements in row " + strconv.Itoa(i) + " does not match the number of rows provided")
		}
	}

	fmt.Printf("Original matrix:\n")
	printMatrix(matrix)

	rotatedMatrix := rotateMatrix(matrix)

	fmt.Printf("\nRotated matrix:\n")
	printMatrix(rotatedMatrix)

}

func rotateMatrix(matrix [][]string) [][]string {
	if len(matrix) <= 1 {
		return matrix
	}

	// rotate will be in place, starting from the outter layer, and going inside layer by layer
	// maintaining the index of the first and last element of the current layer
	// as the matrix is a square matrix there's no need to track row and column separately

	first := 0
	last := len(matrix) - 1

	// ech iteration here is a layer
	for first < last {
		// rotate each element in the layer
		for i := first; i < last; i++ {
			// offset is the distance from the first element of the layer
			// needed to calculate the position of the element to be rotated
			offset := i - first

			// save top
			top := matrix[first][i]

			// left -> top
			matrix[first][i] = matrix[last-offset][first]

			// bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]

			// right -> bottom
			matrix[last][last-offset] = matrix[i][last]

			// top -> right
			matrix[i][last] = top
		}

		// move to the next layer
		first++
		last--
	}

	return matrix
}

func printMatrix(matrix [][]string) {
	var sep string

	for _, row := range matrix {
		for _, c := range row {
			fmt.Printf("%02s ", c)
		}
		sep += "---"
		fmt.Println()
	}

	fmt.Println(sep[:len(sep)-1])
}
