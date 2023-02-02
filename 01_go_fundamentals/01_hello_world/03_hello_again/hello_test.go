package main

import "testing"

// Hello 함수 인자로 빈 문자열 넣었을 때,
// 기본값 "Hello, World" 가 출력
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

// assertion 을 새 함수로 리팩토링
// 1. 중복 감소
// 2. 테스트 가독성 향상
func assertCorrectMessage(
	// testing.TB : *testing.T, *testing.B의 인터페이스
	t testing.TB,
	got, want string) {

	// 이 메서드가 helper 임을 test suite 에 알림
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
