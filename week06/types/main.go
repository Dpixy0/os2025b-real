package main

import (
	"fmt"
	"reflect"
)

// ctrl + F5
func main() {
	var id int
	var gpa float32 = 3.99
	// var name string
	// gpa := 3.99 // float64

	name := "kim Inha"
	// id := 1000
	id = 1000

	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(id, reflect.TypeOf(id))
	fmt.Println(gpa, reflect.TypeOf(gpa))
}
