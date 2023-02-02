package main

import "testing"

// 두번째 파라미터 추가 : 언어 (기본값 영어)
// 스페인어 인사말 추가
// 프랑스어 인사말 추가
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Lauren", french)
		want := "Bonjour, Lauren"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(
	t testing.TB,
	got, want string) {

	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
