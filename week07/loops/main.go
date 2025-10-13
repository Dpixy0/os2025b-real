package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n') //ignore error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
}
