package main

import (
	"strings"
	"testing"
)

func TestTr(t *testing.T) {
	t.Run("Sub small and capital c-letters", func(t *testing.T) {
		input := config{
			text: strings.NewReader("Coding Challenges"),
			subst: map[string]string{
				"C": "c",
			},
		}

		got := Substitute(input)
		want := "coding challenges"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Sub small and capital letters, variation", func(t *testing.T) {
		input := config{
			text: strings.NewReader("Coding Challenges\nHello GOODbye"),
			subst: map[string]string{
				"e": "E",
			},
		}

		got := Substitute(input)
		want := "Coding ChallEngEs\nHEllo GOODbyE"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Multiple subst chars and numbers", func(t *testing.T) {
		input := config{
			text: strings.NewReader("Coding Challenges123\nHelLo GOODbye"),
			// cctr ab12 sd56
			subst: map[string]string{
				"a": "s",
				"b": "d",
				"1": "5",
				"2": "6",
			},
		}

		got := Substitute(input)
		want := "Coding Chsllenges563\nHelLo GOODdye"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
