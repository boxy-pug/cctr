package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestProcessLines(t *testing.T) {
	t.Run("Sub small and capital c-letters", func(t *testing.T) {
		input := config{
			input: strings.NewReader("Coding Challenges"),
			subst: map[string]string{
				"C": "c",
			},
		}

		got := processLines(input)
		want := "coding challenges"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Sub small and capital letters, variation", func(t *testing.T) {
		input := config{
			input: strings.NewReader("Coding Challenges\nHello GOODbye"),
			subst: map[string]string{
				"e": "E",
			},
		}

		got := processLines(input)
		want := "Coding ChallEngEs\nHEllo GOODbyE"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Multiple subst chars and numbers", func(t *testing.T) {
		input := config{
			input: strings.NewReader("Coding Challenges123\nHelLo GOODbye"),
			// cctr ab12 sd56
			subst: map[string]string{
				"a": "s",
				"b": "d",
				"1": "5",
				"2": "6",
			},
		}

		got := processLines(input)
		want := "Coding Chsllenges563\nHelLo GOODdye"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Emoji rune test", func(t *testing.T) {
		input := config{
			input: strings.NewReader("hey👋"),
			// cctr ab12 sd56
			subst: map[string]string{
				"👋": "👀",
				"h": "b",
			},
		}

		got := processLines(input)
		want := "bey👀"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestLoadSubstitution(t *testing.T) {
	t.Run("range substitution", func(t *testing.T) {
		target := "a-d"
		translation := "A-D"

		got, err := loadSubstitution(target, translation)
		want := map[string]string{
			"a": "A",
			"b": "B",
			"c": "C",
			"d": "D",
		}

		if err != nil {
			t.Fatalf("didnt expect error %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
