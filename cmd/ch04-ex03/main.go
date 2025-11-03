package main

import (
	"fmt"
	"log"
	"week10/pkg/keyboard"
)

func main() {
	fmt.Print("화씨 온도 입력 :")
	celsius, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

}
