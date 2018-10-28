package nfa


import (
  "sync"
)

// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// An symbol in the NFA is a single rune, i.e. a character.
type symbol rune

// Given the current state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, sym symbol) []state

// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.

var wg sync.WaitGroup

func Reachable(transitions TransitionFunction, start, final state, input []symbol) bool {
/*
	for i := 0; i < len(input); i++ {
		s := transitions(start, input[i])
		if len(s) == 0 {
		    return false
		} else {
		    for j := 0; j < len(s); j++ {
		 //       wg.Add(1)
			    go Reachable(transitions, s[j], final, input[i:])
		//	    wg.Done()
		    }
		}
	//	wg.Done()
	}
//	return true
*/	
/*
    if len(input) != 1{ //&& start == final {
	    return false
	} 
*/
	s := transitions(start, input[0])
		if len(s) == 0 {
		    return false
		} 
		for j := 0; j < len(s); j++ {
		        wg.Add(1)
		   	    go Reachable(transitions, s[j], final, input[0:])
	    	    //wg.Done()
	    }
	    wg.Done()
		
	/*
	if len(input) == 0 && start == final {
	    return true
	} else {
	    return false
	}*/
	return true
}

