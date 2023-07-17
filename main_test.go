package goForth

import (
	"testing"
)

func TestParseAdd(t *testing.T) {
	input := "4 2 + ."
	want := "6 "

	checkParse(t, input, want)
}

func TestParseMult(t *testing.T) {
	input := "4 2 * ."
	want := "8 "

	checkParse(t, input, want)
}

func TestParseSub(t *testing.T) {
	input := "4 2 - ."
	want := "2 "

	checkParse(t, input, want)
}

func TestParseDiv(t *testing.T) {
	input := "4 2 / ."
	want := "2 "

	checkParse(t, input, want)
}

func TestParseMore(t *testing.T) {
	input := "4 2 + 3 - ."
	want := "3 "

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

func TestParseSwap(t *testing.T) {
	input := "1 2 SWAP . ."
	want := "1 2 "

	checkParse(t, input, want)
}

func TestParseDup(t *testing.T) {
	input := "2 DUP . ."
	want := "2 2 "

	checkParse(t, input, want)
}

func checkParse(t *testing.T, input string, want string) {
	got, err := Parse(input)

	if want != got || err != nil {
		t.Fatalf(`Parse("%s") = (%s, %v), want match for (%s, nil)`, input, got, err, want)
	}
}

func TestParseWords(t *testing.T) {
	input := "WORDS"

	got, err := Parse(input)
	t.Log(got)

	if err != nil {
		t.Fatalf(`Parse("%s") = (%s, %v)`, input, got, err)
	}
}
