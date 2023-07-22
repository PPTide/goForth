package main

import (
	"bufio"
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
	":": forthDictionaryEntry{
		codeSpace: func(st *state) {
			(*st).interpreting = false
			(*st).compiling = true

			(*st).definitionName = forthNameSpace(readName(st))
		},
	},
}

func main() {
	st := initializeState()

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
	(*st).input = strings.NewReader(input)
	reader := (*st).input

	for true {
		if reader.Len() <= 0 {
			break
		}

		nameA := readName(st)

		if (*st).compiling && string(nameA) == ";" {
			definitionStack := (*st).definitionStack
			(*st).dictionary[(*st).definitionName] = forthDictionaryEntry{
				codeSpace: func(st *state) {
					for _, entry := range definitionStack {
						entry.codeSpace(st)
					}
				},
			}

			(*st).compiling = false
			(*st).interpreting = true
			(*st).definitionStack = make(forthDefinitionStack, 0)
			continue
		}
		n, exist := (*st).dictionary[forthNameSpace(nameA)]
		if exist {
			if (*st).interpreting {
				n.codeSpace(st)
				continue
			}
			if (*st).compiling {
				(*st).definitionStack = append((*st).definitionStack, n)
				continue
			}
		}
		if n, err := convertInputNumber(string(nameA)); err == nil {
			if (*st).interpreting {
				(*st).dataStack = append((*st).dataStack, n)
				continue
			}
			if (*st).compiling {
				(*st).definitionStack = append((*st).definitionStack, forthDictionaryEntry{codeSpace: func(st *state) {
					(*st).dataStack = append((*st).dataStack, n)
				}})
				continue
			}
		}
		// https://forth-standard.org/standard/usage#usage:ambiguous
		return fmt.Errorf("ambiguous condition: %s", string(nameA))
	}
	return nil
}

func readName(st *state) []rune {
	reader := (*st).input

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
	return nameA
}

func convertInputNumber(str string) (int, error) {
	//TODO: Interpret numbers correctly https://forth-standard.org/standard/usage#subsubsection.3.4.1.3
	return strconv.Atoi(str)
}

func initializeState() *state {
	st := &state{
		dictionary:   dictionary,
		interpreting: true,
	}
	(*st).dictionary["QUIT"].codeSpace(st)
	return st
}
