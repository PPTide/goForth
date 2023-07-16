package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ops = map[string]func(stack *[]int, output *string){
	"+": func(stack *[]int, _ *string) {
		x2 := pop(stack)
		x1 := pop(stack)

		*stack = append(*stack, x1+x2)
	},
	"-": func(stack *[]int, _ *string) {
		x2 := pop(stack)
		x1 := pop(stack)

		*stack = append(*stack, x1-x2)
	},
	"*": func(stack *[]int, _ *string) {
		x2 := pop(stack)
		x1 := pop(stack)

		*stack = append(*stack, x1*x2)
	},
	"/": func(stack *[]int, _ *string) {
		x2 := pop(stack)
		x1 := pop(stack)

		*stack = append(*stack, x1/x2)
	},
	".": func(stack *[]int, output *string) {
		*output += fmt.Sprintf("%d\n", pop(stack))
	},
}

func main() {
	parse("4 2 + .")
}

func parse(input string) (string, error) {
	reader := strings.NewReader(input)
	output := ""
	stack := make([]int, 0)

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

func act(buf *[]rune, stack *[]int, output *string) error {
	x := string(*buf)
	*buf = make([]rune, 0)

	i, err := strconv.Atoi(x)
	if err == nil {
		*stack = append(*stack, i)
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
