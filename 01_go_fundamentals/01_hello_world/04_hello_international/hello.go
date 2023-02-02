package main

import "fmt"

const englishHelloPrefix = "Hello, "

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

const french = "French"
const frenchHelloPrefix = "Bonjour, "

// 두 번째 파라미터로 언어 추가
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// 언어에 따른 인사말 접두사 반환
// go 함수: 반환 타입, 이름을 지정할 수 있다
func greetingPrefix(language string) (prefix string) {

	// switch 문
	// switch 의 값과 case 의 값이 일치하면 실행문 작동

	// case 값 여러개 (or) 넣고 싶으면 , 콤마로 연결.
	// case 값 (if else 처럼) 조건문 써도 된다.

	// break 키워드 기본으로 작동: 안씀
	// fallthrough 키워드: 다음 케이스 조건문 무시하고 다음 실행문 작동

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
