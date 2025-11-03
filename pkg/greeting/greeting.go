package greeting

import "fmt"

//func hello() // 함수 이름이 소문자로 시작할 시 외부에 노출이 안됨
func Hello() {
	fmt.Printf("Hello!~\n")
}

func Hi() {
	fmt.Printf("Hi!~\n")
}
