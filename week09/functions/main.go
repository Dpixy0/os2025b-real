package main

import "fmt"

func swap(first *int, second *int) {
	temp := first
	first = second
	second = temp
	fmt.Println(first, second)
}

func main() {
	//fmt.Println(math.Sqrt(-16.0))
	a, b := 10, 20
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}
