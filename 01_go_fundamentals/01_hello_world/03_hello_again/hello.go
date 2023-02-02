package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello 함수에 이름을 넣어 개인화한 인사말로 변경
// 빈 문자열을 이름으로 넣을 때, 기본값으로 Hello, World 반환
func Hello(name string) string {

	// 이름이 빈 문자열 이면 기본값 설정.
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
