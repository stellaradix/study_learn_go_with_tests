// main package는 entry point를 가지는 package
package main

import "fmt"

// 함수 선언 형식 - 기본
// func 함수명 (parameter) (return) {}

/* 함수 이름 */
// 함수 이름이 대문자로 시작하면 public
// 함수 이름이 소문자로 시작하면 private

/* 함수 return */
// 반환 타입 지정
// 반환 변수 지정 가능 - return 뒤에 변수 생략
// 여러개 설정 가능

// 0개 단어 모두 생략
// 1개 단어 괄호 생략 가능 (권장)
// 1개 단어 초과 괄호 생략 불가능

// 반환 없을 때 정의부에서 return 없어도 된다
// 반환 있을 때 정의부에 return 필수

func Hello() string {
	return "Hello, World"
}

// main package 의 main 함수는 entry point
func main() {

	// fmt.Println() : 콘솔 메세지 출력
	fmt.Println(Hello())
}
