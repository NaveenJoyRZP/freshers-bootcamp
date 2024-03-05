package main

import "fmt"

type Person struct {
  name string
  age int
}

func newPerson(name string, age int) *Person {
  person := Person{name: name, age: age}
  return &person
}

func main() {
  person1 := newPerson("Naveen", 25);
  fmt.Println(*person1)

  // Anonymous inner structs
  dog := struct {
    name string
    isGood bool
  } {
    name: "Cassy",
    isGood: true,
  }
  fmt.Println(dog)

}
