package goForth

import (
	"fmt"
	"strconv"
	"strings"
)

type word interface {
	getType() string
}

type forthInt int

type forthExecutionToken func(st state)

func (i forthInt) getType() string {
	return "forthInt"
}

func (xt forthExecutionToken) getType() string {
	return "xt"
}

type state struct {
	reader *strings.Reader
	buf    []rune
	stack  []word
	output string
	ops    map[string]func(st *state)
}

var ops = map[string]func(st *state){
	"+": func(state *state) {
		stack := &(*state).stack
		x2 := pop(stack).(forthInt)
		x1 := pop(stack).(forthInt)

		*stack = append(*stack, x1+x2)
	},
	"-": func(state *state) {
		stack := &(*state).stack
		x2 := pop(stack).(forthInt)
		x1 := pop(stack).(forthInt)

		*stack = append(*stack, x1-x2)
	},
	"*": func(state *state) {
		stack := &(*state).stack
		x2 := pop(stack).(forthInt)
		x1 := pop(stack).(forthInt)

		*stack = append(*stack, x1*x2)
	},
	"/": func(state *state) {
		stack := &(*state).stack
		x2 := pop(stack).(forthInt)
		x1 := pop(stack).(forthInt)

		*stack = append(*stack, x1/x2)
	},
	".": func(state *state) {
		stack := &(*state).stack
		output := &(*state).output
		x := pop(stack)
		if x.getType() == "forthInt" {
			*output += fmt.Sprintf("%d ", x)
			return
		}
		*output += fmt.Sprintf("%s", x)
	},
	"EMIT": func(st *state) {
		(*st).output += fmt.Sprintf("%c", rune(pop(&(*st).stack).(forthInt)))
	},
	"CR": func(state *state) {
		output := &(*state).output
		*output += "\n"
	},
	"SWAP": func(state *state) {
		stack := &(*state).stack
		x1 := pop(stack)
		x2 := pop(stack)

		*stack = append(*stack, x1, x2)
	},
	"DUP": func(state *state) {
		stack := &(*state).stack
		x := pop(stack)

		*stack = append(*stack, x, x)
	},
	"WORDS": func(st *state) {
		for s := range (*st).ops {
			(*st).output += fmt.Sprintf("%s ", s)
		}
	},
	":": func(st *state) { // Word definition
		//TODO: implement
		panic("not implemented")
	},
	"EXECUTE": func(st *state) {

	},
}

func Parse(input string) (string, error) {
	st := &state{
		reader: strings.NewReader(input),
		buf:    make([]rune, 0),
		stack:  make([]word, 0),
		output: "",
		ops:    ops,
	}

	interpret(st)

	return st.output, nil
}

func interpret(st *state) { // https://forth-standard.org/standard/usage#section.3.4
a: // a. Skip leading spaces and parse a name (see 3.4.1); //TODO: skip leading spaces
	name := parseWord(st)

	if len(name) <= 0 {
		// error?
		return
	}

	// b. Search the dictionary name space (see 3.4.2). If a definition name matching the string is found:
	xt, ok := (*st).ops[name]
	if ok {
		// 1. if interpreting, perform the interpretation semantics of the definition (see 3.4.3.2), and continue at a).
		xt(st)
		goto a
	}

	// c. If a definition name matching the string is not found, attempt to convert the string to a number (see 3.4.1.3). If successful:
	//TODO: implement https://forth-standard.org/standard/usage#usage:numbers
	num, err := strconv.Atoi(name)
	if err == nil {
		// 1. if interpreting, place the number on the data stack, and continue at a);
		(*st).stack = append((*st).stack, forthInt(num))
		goto a
	}

	// d. If unsuccessful, an ambiguous condition exists (see 3.4.4).
	panic("word not found")
}

func parseWord(st *state) string {
	reader := (*st).reader
	for reader.Len() >= 1 {
		r, _, _ := reader.ReadRune()
		if r == ' ' {
			break
		}
		(*st).buf = append((*st).buf, r)
	}
	w := string((*st).buf)
	(*st).buf = make([]rune, 0)
	return w
}

func pop[T any](stack *[]T) T {
	x := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return x
}
