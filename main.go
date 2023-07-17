package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type word interface {
	getType() string
	int() forthInt
}

type forthInt int

func (i forthInt) int() forthInt {
	return forthInt(i)
}

func (i forthInt) getType() string {
	return "forthInt"
}

var ops = map[string]func(stack *[]word, output *string){
	"+": func(stack *[]word, _ *string) {
		x2 := pop(stack).int()
		x1 := pop(stack).int()

		*stack = append(*stack, x1+x2)
	},
	"-": func(stack *[]word, _ *string) {
		x2 := pop(stack).int()
		x1 := pop(stack).int()

		*stack = append(*stack, x1-x2)
	},
	"*": func(stack *[]word, _ *string) {
		x2 := pop(stack).int()
		x1 := pop(stack).int()

		*stack = append(*stack, x1*x2)
	},
	"/": func(stack *[]word, _ *string) {
		x2 := pop(stack).int()
		x1 := pop(stack).int()

		*stack = append(*stack, x1/x2)
	},
	".": func(stack *[]word, output *string) {
		x := pop(stack)
		if x.getType() == "forthInt" {
			*output += fmt.Sprintf("%d ", x)
			return
		}
		*output += fmt.Sprintf("%s", x)
	},
	"EMIT": func(stack *[]word, output *string) {
		*output += fmt.Sprintf("%c", rune(pop(stack).int()))
	},
	"CR": func(_ *[]word, output *string) {
		*output += "\n"
	},
	"SWAP": func(stack *[]word, _ *string) {
		x1 := pop(stack)
		x2 := pop(stack)

		*stack = append(*stack, x1, x2)
	},
	"DUP": func(stack *[]word, _ *string) {
		x := pop(stack)

		*stack = append(*stack, x, x)
	},
}

func main() {}

func parse(input string) (string, error) {
	reader := strings.NewReader(input)
	output := ""
	stack := make([]word, 0)

	buf := make([]rune, 0)
	for true {
		if reader.Len() <= 0 {
			err := act(&buf, &stack, &output)
			if err != nil {
				return output, err
			}
			break
		}
		r, _, _ := reader.ReadRune()

		if r == ' ' {
			err := act(&buf, &stack, &output)
			if err != nil {
				return output, err
			}
			continue
		}

		buf = append(buf, r)
	}
	return output, nil
}

func act(buf *[]rune, stack *[]word, output *string) error {
	x := string(*buf)
	*buf = make([]rune, 0)

	i, err := strconv.Atoi(x)
	if err == nil {
		*stack = append(*stack, forthInt(i))
		return nil
	}

	if val, ok := ops[x]; ok {
		val(stack, output)
		return nil
	}

	return errors.New("not implemented")
}

func pop[T any](stack *[]T) T {
	x := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return x
}
