package main

import (
	"fmt"
)

type matrix struct {
  rowCount int
  columnCount int
  elements [][]int
}

func (matrix matrix) getTotalRows() int {
  return matrix.rowCount
}

func (matrix matrix) getTotalColumns() int {
  return matrix.columnCount
}

func (matrix *matrix) setElements(i , j, value int) {
  matrix.elements[i][j] = value
}

func (matrix *matrix) addMatrix(matrixToAdd matrix) {
  for i, row := range matrix.elements {
    for j, _ := range row {
      matrix.elements[i][j] = matrixToAdd.elements[i][j] 
    }
  }
}

func (matrix *matrix) printMatrixInJson() {
  fmt.Println("{")
  for _, row := range matrix.elements {
    for _, column := range row {
      fmt.Printf("%v ", column)
    }
    fmt.Println("")
  }
  fmt.Println("}")
}

func main() {
  elements1 := [][]int{{0,1,2},{3,4,5},{5,6,7}}
  elements2 := [][]int{{1,2,5},{3,4,5},{5,6,7}}
  matrix1 := matrix{rowCount: 3, columnCount: 3, elements: elements1}
  matrix2 := matrix{rowCount: 3, columnCount: 3, elements: elements2}
  // Printing the matrix
  matrix1.printMatrixInJson()

  // Adding another matrix
  matrix1.addMatrix(matrix2)
  matrix1.printMatrixInJson()

  // Setting another value to a position
  matrix1.setElements(0,0,9)
  matrix1.printMatrixInJson()

}
