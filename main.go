package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var dictionary = forthDictionary{
	"QUIT": forthDictionaryEntry{
		codeSpace: func(st *state) {
			// https://forth-standard.org/standard/core/QUIT
		},
		dataSpace: forthDataSpace{},
	},
	".": forthDictionaryEntry{
		codeSpace: func(st *state) {
			fmt.Printf("%d ", (*st).dataStack.pop())
		},
	},
	"+": forthDictionaryEntry{
		codeSpace: func(st *state) {
			x2 := (*st).dataStack.pop()
			x1 := (*st).dataStack.pop()

			(*st).dataStack = append((*st).dataStack, x1+x2)
		},
	},
	"-": forthDictionaryEntry{
		codeSpace: func(st *state) {
			x2 := (*st).dataStack.pop()
			x1 := (*st).dataStack.pop()

			(*st).dataStack = append((*st).dataStack, x1-x2)
		},
	},
	"*": forthDictionaryEntry{
		codeSpace: func(st *state) {
			x2 := (*st).dataStack.pop()
			x1 := (*st).dataStack.pop()

			(*st).dataStack = append((*st).dataStack, x1*x2)
		},
	},
	"/": forthDictionaryEntry{
		codeSpace: func(st *state) {
			x2 := (*st).dataStack.pop()
			x1 := (*st).dataStack.pop()

			(*st).dataStack = append((*st).dataStack, x1/x2)
		},
	},
}

func main() {
	st := &state{
		dictionary: dictionary,
	}
	(*st).dictionary["QUIT"].codeSpace(st)

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		fmt.Print("\033[A") // Move Cursor Up

		fmt.Printf("> %s ", input)

		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		err = interpret(input, st)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("OK\n")
	}
}

// https://forth-standard.org/standard/usage#section.3.4
func interpret(input string, st *state) error {
	reader := (*st).input
	reader = strings.NewReader(input)

	for true {
		if reader.Len() <= 0 {
			break
		}

		nameA := make([]rune, 0)
		for true { // Read Name
			if reader.Len() <= 0 {
				break
			}
			r, _, err := reader.ReadRune()
			if err != nil {
				log.Fatal(err)
			}
			if r == ' ' && len(nameA) == 0 {
				continue
			}
			if r == ' ' {
				break
			}
			nameA = append(nameA, r)
		}

		n, exist := (*st).dictionary[forthNameSpace(nameA)]
		if exist {
			// if interpreting
			n.codeSpace(st)
			continue
		}
		if n, err := convertInputNumber(string(nameA)); err == nil {
			// if interpreting
			(*st).dataStack = append((*st).dataStack, n)
			continue
		}
		// https://forth-standard.org/standard/usage#usage:ambiguous
		return errors.New("Ambiguous condition")
	}
	return nil
}

func convertInputNumber(str string) (int, error) {
	//TODO: Interpret numbers correctly https://forth-standard.org/standard/usage#subsubsection.3.4.1.3
	return strconv.Atoi(str)
}

func initializeState() *state {
	st := &state{
		dictionary: dictionary,
	}
	(*st).dictionary["QUIT"].codeSpace(st)
	return st
}
