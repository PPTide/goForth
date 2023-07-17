package goForth

import (
	"errors"
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
	reader := (*st).reader

	for true {
		if reader.Len() <= 0 {
			err := act(st)
			if err != nil {
				return st.output, err
			}
			break
		}
		r, _, _ := reader.ReadRune()

		if r == ' ' {
			err := act(st)
			if err != nil {
				return st.output, err
			}
			continue
		}

		st.buf = append(st.buf, r)
	}
	return st.output, nil
}

func act(state *state) error {
	buf := &(*state).buf
	x := string(*buf)
	*buf = make([]rune, 0)

	i, err := strconv.Atoi(x)
	if err == nil {
		(*state).stack = append((*state).stack, forthInt(i))
		return nil
	}

	if val, ok := ops[x]; ok {
		val(state)
		return nil
	}

	return errors.New("not implemented")
}

func pop[T any](stack *[]T) T {
	x := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return x
}
