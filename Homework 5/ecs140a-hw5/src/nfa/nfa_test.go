package nfa

import (
	"testing"
	"time"
)

func blahTransitions(st state, sym symbol) []state {
	/*
	 * 0 -a-> 0
	 * 0 -a-> 1
	 * 1 -b-> 1
	 */
	return map[state]map[symbol][]state{
		0: map[symbol][]state{
			'a': []state{0, 1},
		},
		1: map[symbol][]state{
			'b': []state{1},
		},
	}[st][sym]
}

func expTransitions(st state, sym symbol) []state {
	/*
	 * 0 -a-> 1
	 * 0 -a-> 2
	 * 1 -a-> 0
	 * 1 -a-> 2
	 * 2 -a-> 0
	 * 2 -a-> 1
	 */
	return map[state]map[symbol][]state{
		0: map[symbol][]state{
			'a': []state{1, 2},
		},
		1: map[symbol][]state{
			'a': []state{0, 2},
		},
		2: map[symbol][]state{
			'a': []state{0, 1},
		},
		3: map[symbol][]state{},
	}[st][sym]
}

func langTransitions(st state, sym symbol) []state {
	/*
	 * Matches the regular language /(ab)*[ab]/
	 */
	return map[state]map[symbol][]state{
		0: map[symbol][]state{
			'a': []state{1, 2},
			'b': []state{2},
		},
		1: map[symbol][]state{
			'b': []state{0},
		},
		2: map[symbol][]state{},
	}[st][sym]
}

func TestReachable(t *testing.T) {
	graphs := map[string]TransitionFunction{
		"blahTransitions": blahTransitions,
		"expTransitions":  expTransitions,
		"langTransitions": langTransitions,
	}

	tests := []struct {
		graph string
		start state
		end   state
		str   string
		want  bool
	}{
		{"blahTransitions", 0, 1, "aaab", true},
		{"blahTransitions", 0, 1, "aaa", true},
		{"blahTransitions", 0, 1, "aaaba", false},
		{"expTransitions", 0, 2, "aaaa", true},
		{"expTransitions", 0, 1, "aaaa", true},
		{"expTransitions", 0, 3, "aaaaaaaaaa", false},
		{"langTransitions", 0, 2, "ababb", true},
		{"langTransitions", 0, 2, "aba", true},
		{"langTransitions", 0, 0, "abab", true},
		{"langTransitions", 0, 0, "aba", false},
	}

	for _, test := range tests {
		done := make(chan bool)
		go func() {
			done <- Reachable(graphs[test.graph], test.start, test.end, ([]symbol)(test.str))
		}()

		select {
		case <-time.NewTimer(5 * time.Millisecond).C:
			t.Errorf("nfa timed out (too slow) on (%s, %d, %d, %#v)",
				test.graph, test.start, test.end, test.str)
		case got := <-done:
			if test.want != got {
				t.Errorf("nfa failed on (%s, %d, %d, %#v); want %t, got %t",
					test.graph, test.start, test.end, test.str,
					test.want, got)
			}
		}
	}
}

func BenchmarkReachable(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reachable(expTransitions, 0, 2, ([]symbol)("aaaaaaaaaaaaaa"))
	}
}
