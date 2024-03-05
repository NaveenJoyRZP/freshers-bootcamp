package main

import "fmt"

func zeroVal(num int) {
  num = 0
}

func zeroPtr(num *int) {
  *num = 0
}

func main() {

  // Pointers
  i := 5
  fmt.Println("Initial i = ", i)
  zeroVal(i)
  fmt.Println("Zero val i = ", i)
  zeroPtr(&i)
  fmt.Println("Pointer val i = ", i)

  // Arrays
  arr := [5]int{1,2,3,4,5}
  slice := []int{1,2,3,4,5,6,7,8}
  fmt.Println("Lenght of array : ", len(arr))
  fmt.Println("Lenght of slice : ", len(slice))

  // 2D array
  var twoDArray [3][3]int
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      twoDArray[i][j] = i + j
    }
  }

  fmt.Println("2d Array : ", twoDArray)
  printSlice(twoDArray)

  // Slices
  strSlice := make([]string, 3)
  strSlice[0] = "a"
  strSlice[1] = "b"
  strSlice[2] = "c"
  strSlice = append(strSlice, "d")
  fmt.Println(strSlice)

  // Copying a slice
  copySlice := make([]string, len(strSlice))
  copy(copySlice, strSlice)
  fmt.Println(copySlice)

  // Slice slicing
  newStrSlice := []string{"n","a","v","e","e","n"," ","j","o","y"}
  fmt.Println(newStrSlice[:5])
  fmt.Println(newStrSlice[5:])

  // Maps
  map1 := make(map[int]string)
  map1[1] = "Naveen"
  map1[2] = "Joy"
  fmt.Println(map1)
  delete(map1, 1)
  fmt.Println(map1)
  map1[1] = "Geethu"
  fmt.Println(map1)
  value, isPresent := map1[2]
  fmt.Printf("Map[2] = %v, is-present = %v\n",value, isPresent)

  n := map[int]string{1: "Evelyn", 2: "Sara", 3: "Georgie"}
  fmt.Println(n)

  // Looping through a map
  for key, value := range n {
    fmt.Printf("Key = %v : value = %v\n", key, value)
  }

  sampleString := "Naveen"
  for index, val := range sampleString {
    fmt.Printf("index = %v, val = %v\n", index, val)
  }
   

}

func printSlice(slice [3][3]int) {
  for i := 0; i < len(slice); i++ {
    fmt.Print("[")
    for j := 0; j < len(slice); j++ {
      fmt.Print(slice[i][j])
    }
    fmt.Print("]\n")
  }
}
