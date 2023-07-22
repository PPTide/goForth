package main

import (
	"strings"
	"testing"
)

func parseTest(testString string, st *state, t *testing.T) {
	//TODO: pass state as an argument record stack depth so it can be eliminated
	st2 := *st

	// parse Test Harness
	if !(testString[:3] == "T{ " && testString[len(testString)-3:] == " }T") {
		t.Fatalf(`Test Harness "%s" not parsed correctly`, testString)
	}
	testStringP := testString[3 : len(testString)-3]

	splitTestString := strings.Split(testStringP, " -> ")
	if len(splitTestString) != 2 {
		t.Fatalf(`Test Harness "%s" not parsed correctly`, testString)
	}

	runString := splitTestString[0]
	err := interpret(runString, st)
	if err != nil {
		t.Fatalf(`Error "%v" while interpreting "%s"`, err, runString)
	}

	resultString := splitTestString[1]
	err = interpret(resultString, &st2)
	if err != nil {
		t.Fatalf(`Error "%v" while interpreting "%s"`, err, resultString)
	}

	if !slicesEqual((st2).dataStack, (*st).dataStack) {
		t.Fatalf(`Running "%s", expected "%v", got "%v"`, runString, (st2).dataStack, (*st).dataStack)
	}
}

func parseTests(tests string, st *state, t *testing.T) {
	for _, testString := range strings.Split(tests, "\n") {
		parseTest(testString, st, t)
	}
}

func slicesEqual[T int](a1 []T, a2 []T) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func TestPlus(t *testing.T) {
	st := initializeState()

	tests := `T{        0  5 + ->          5 }T
T{        5  0 + ->          5 }T
T{        0 -5 + ->         -5 }T
T{       -5  0 + ->         -5 }T
T{        1  2 + ->          3 }T
T{        1 -2 + ->         -1 }T
T{       -1  2 + ->          1 }T
T{       -1 -2 + ->         -3 }T
T{       -1  1 + ->          0 }T`

	parseTests(tests, st, t)
}

func TestMinus(t *testing.T) {
	st := initializeState()

	tests := `T{          0  5 - ->       -5 }T
T{          5  0 - ->        5 }T
T{          0 -5 - ->        5 }T
T{         -5  0 - ->       -5 }T
T{          1  2 - ->       -1 }T
T{          1 -2 - ->        3 }T
T{         -1  2 - ->       -3 }T
T{         -1 -2 - ->        1 }T
T{          0  1 - ->       -1 }T`

	parseTests(tests, st, t)
}

func TestMult(t *testing.T) {
	st := initializeState()

	tests := `T{  0  0 * ->  0 }T
T{  0  1 * ->  0 }T
T{  1  0 * ->  0 }T
T{  1  2 * ->  2 }T
T{  2  1 * ->  2 }T
T{  3  3 * ->  9 }T
T{ -3  3 * -> -9 }T
T{  3 -3 * -> -9 }T
T{ -3 -3 * ->  9 }T`

	parseTests(tests, st, t)
}

/*func TestDiv(t *testing.T) {
	st := initializeState()

	tests := `T{       0       1 / ->       0 }T
T{       1       1 / ->       1 }T
T{       2       1 / ->       2 }T
T{      -1       1 / ->      -1 }T
T{      -2       1 / ->      -2 }T
T{       0      -1 / ->       0 }T
T{       1      -1 / ->      -1 }T
T{       2      -1 / ->      -2 }T
T{      -1      -1 / ->       1 }T
T{      -2      -1 / ->       2 }T
T{       2       2 / ->       1 }T
T{      -1      -1 / ->       1 }T
T{      -2      -2 / ->       1 }T
T{       7       3 / ->       2 }T
T{       7      -3 / ->      -3 }T
T{      -7       3 / ->      -3 }T
T{      -7      -3 / ->       2 }T`

	parseTests(tests, st, t) // WHY THE FUCK DOES FORTH NOT SPECIFY IF YOU HAVE TO ROUND TO -∞ OR 0???
}*/
