package main

import "testing"

func TestHello(t *testing.T) {
	// 함수 파라미터에 인자 넣기
	got := Hello("Chris")

	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
