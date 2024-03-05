package main

import (
	"fmt"
	"math"
	"runtime"
)

// Variables at package level
var bool1, bool2 bool

func main() {
  fmt.Println("Sample code")

  // Variables at function level
  var var1, var2 string
  fmt.Printf("var1 = %v, var2 = %v\n", var1, var2)
  fmt.Printf("bool1 = %v, bool2 = %v\n", bool1, bool2)

  // Type inferred variable declarations
  var typeDeclaredStr1, typeDeclaredStr2 = "Hi", "Hello"
  fmt.Printf("typeDeclaredStr1 = %v, typeDeclaredStr2 = %v\n", typeDeclaredStr1, typeDeclaredStr2)


  // Short variable declaration - Here type is inferred from 
  shortDeclaredStr1 := "I am declared without var keyword"
  fmt.Printf("Variable declared without var keyoword : %v\n", shortDeclaredStr1)

  // Float variables manipulation
  fmt.Println("7.0/3.0", 7.0/3.0)

  // For loop
  // Simple for loop
  for i := 0; i < 10; i++ {
    fmt.Printf("i : %v ", i)
  }

  fmt.Println("")

  // While loop (for)
  whilei := 2
  for whilei < 20 {
    fmt.Printf("while i : %v ", whilei)
    whilei += 2
  }

  fmt.Println("")

  // If Else statements
  if  whilei > 20 {
    fmt.Println("i is greater than 20")
  } else {
    fmt.Println("i is less thatn 20")
  }

  // Use of a function
  sqrt1 := findSqrt(-16)
  sqrt2 := findSqrt(25)
  fmt.Printf("sqrt(-16) = %v and sqrt(25) = %v\n", sqrt1, sqrt2)


  // Switch statements
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("System running MacOS- OS X")
  case "linux":
    fmt.Println("System running linux")
  default:
    fmt.Println("Cannot determine OS")
  }

  // Defer statements
  deferFunction()

  // Funcrtions
  var swapStr1, swapStr2 = "Naveen", "Joy"
  swappedStr1, swappedStr2 := swap(swapStr1, swapStr2)
  fmt.Printf("Swappped %v to %v and %v to %v", swapStr1, swappedStr1, swapStr2, swappedStr2)

  // Vardiac functions
  sum := sumNums(1,2,3,4,5)
  fmt.Println("\nSUM = ", sum)
  nums := []int{1,2,3,4}
  sum2 := sumNums(nums...)
  fmt.Println("SUM2 = ", sum2)


  // Closures
  nextInt := closureFuncNextInt()
  fmt.Print(nextInt())
  fmt.Print(nextInt())
  fmt.Print(nextInt())
  fmt.Print(nextInt())

  nextInt2 := closureFuncNextInt()
  fmt.Print(nextInt2())
  fmt.Print(nextInt2())
  fmt.Print(nextInt2())
  fmt.Print(nextInt2())

  factNum := 5
  fmt.Printf("Factorial of %v is %v\n", factNum, factorial(factNum))

}

func factorial(num int) float64 {
  if num == 0 {
    return 1
  }
  return float64(num) * factorial(num - 1)
}

func closureFuncNextInt() func() int {
  var i int = 0
  return func() int {
    i++
    return i
  }
}

func sumNums(nums ...int) int {
  sum := 0
  for _, num := range nums {
    sum += num
  }
  return sum
}

func swap(str1, str2 string) (string, string) {
  return str2, str1
}

func findSqrt(num float64) string {
  if num < 0 {
    return fmt.Sprint(math.Sqrt(-num))
  } else {
    return fmt.Sprint(math.Sqrt(num))
  }
}

func deferFunction() {
  defer fmt.Println("Defere statement executed")
  fmt.Println("Fucntion called..")

}
