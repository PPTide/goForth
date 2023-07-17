package main

import (
	"errors"
	"testing"
)

func TestParseAdd(t *testing.T) {
	input := "4 2 + ."
	want := "6"

	checkParse(t, input, want)
}

func TestParseMult(t *testing.T) {
	input := "4 2 * ."
	want := "8"

	checkParse(t, input, want)
}

func TestParseSub(t *testing.T) {
	input := "4 2 - ."
	want := "2"

	checkParse(t, input, want)
}

func TestParseDiv(t *testing.T) {
	input := "4 2 / ."
	want := "2"

	checkParse(t, input, want)
}

func TestParseMore(t *testing.T) {
	input := "4 2 + 3 - ."
	want := "3"

	checkParse(t, input, want)
}

func TestParseEmit(t *testing.T) {
	input := "66 EMIT"
	want := "B"

	checkParse(t, input, want)
}

func TestParseCR(t *testing.T) {
	input := "66 EMIT"
	want := "B"

	checkParse(t, input, want)
}

func checkParse(t *testing.T, input string, want string) {
	got, err := parse(input)

	if want != got || err != nil {
		t.Fatalf(`parse("%s") = (%s, %v), want match for (%s, nil)`, input, got, err, want)
	}
}

func TestParseInvalidWord(t *testing.T) {
	input := "4 2 dsgb"

	got, err := parse(input)

	if errors.Is(errors.New("not implemented"), err) {
		t.Fatalf(`parse("%s") = (%s, %v), want match for (%s, nil)`, input, got, err, "")
	}
}

func checkArray[T int](a1 []T, a2 []T) bool {
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
