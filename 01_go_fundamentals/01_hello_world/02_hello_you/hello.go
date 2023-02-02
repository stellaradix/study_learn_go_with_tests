package main

import "fmt"

// 상수 선언
// 1. 값의 의미를 상수 이름으로 표시
// 2. 성능 향상
const englishHelloPrefix = "Hello, "

// 함수 parameter
// 괄호 생략 불가능
// 연속된 parameter 타입이 같을 경우 마지막에만 타입 선언 가능

// Hello 함수에 이름을 넣어 개인화한 인사말로 변경
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
