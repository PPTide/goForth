package main

import "strings"

// https://forth-standard.org/

// ---------------- Stacks ----------------

// C:
//
// * last-in, first-out
//
// * elements define the permissible matchings of control-flow words and the restrictions imposed on
// data-stack usage during the compilation of control structures
//
// * The elements of the control-flow stack are system-compilation data types
// https://forth-standard.org/standard/usage#subsubsection.3.2.3.2
type controlFlowStack []int

// S: (When there is no confusion, the data-stack stack-id may be omitted.)
//
// Objects on the data stack shall be one cell wide.
// https://forth-standard.org/standard/usage#subsubsection.3.2.3.1
type dataStack []int

func (ds *dataStack) pop() int {
	x := (*ds)[len(*ds)-1]
	*ds = (*ds)[0 : len(*ds)-1]
	return x
}

// R:
//
// * A program shall not access values on the return stack (using R@, R>, 2R@, 2R> or NR>) that it did not place there using >R, 2>R or N>R;
//
// * A program shall not access from within a do-loop values placed on the return stack before the loop was entered;
//
// * All values placed on the return stack within a do-loop shall be removed before I, J, LOOP, +LOOP, UNLOOP, or LEAVE is executed;
//
// * All values placed on the return stack within a definition shall be removed before the definition is terminated or before EXIT is executed.
//
// https://forth-standard.org/standard/usage#subsubsection.3.2.3.3
type returnStack []int

type forthDictionary map[forthNameSpace]forthDictionaryEntry
type forthNameSpace string
type forthDictionaryEntry struct {
	codeSpace forthCodeSpace
	dataSpace forthDataSpace
}
type forthCodeSpace func(st *state)
type forthDataSpace struct {
}

type state struct {
	input      *strings.Reader
	dictionary forthDictionary

	dataStack dataStack
}
