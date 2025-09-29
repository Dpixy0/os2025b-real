package main

import (
	"fmt"
	"reflect"
)

func main() {
	totalPrice := 1000

	fmt.Println(totalPrice, reflect.TypeOf(totalPrice))
}
