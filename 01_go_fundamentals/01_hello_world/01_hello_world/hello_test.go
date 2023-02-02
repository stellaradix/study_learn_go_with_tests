package main

import "testing"

func TestHello(t *testing.T) {

	// := 변수 선언과 할당을 동시에 하는 연산자
	got := Hello()
	want := "Hello, World"

	// if문은 다른 프로그래밍 언어와 유사
	// 조건문에 괄호 없다
	// 반드시 boolean 식으로 판별 되어야 한다 (0, 1 불가)
	if got != want {

		// t.Errorf() : 테스트 실패, 콘솔 에러 메세지 출력

		// 문자열 포맷 - %s 문자열 대입
		// 문자열 포맷 - %q 문자열을 쌍따옴표 감싸 대입
		t.Errorf("got %q want %q", got, want)
	}

	// t.Fail() : 테스트 실패 코드
}

/* 테스트 코드 작성

테스트 파일 이름 xxx_test.go
테스트 함수 이름 Test 로 시작
테스트 함수 매개변수는 오직 t *testing.T 만 존재

*testing.T 를 쓰기 위해 testing 모듈 import

테스트 실패 코드 : t.Fail()

*/
