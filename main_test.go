package main

import (
	//"reflect"
	"bytes"
	"strings"
	"testing"
)

func TestProcessLines(t *testing.T) {
	t.Run("Sub small and capital c-letters", func(t *testing.T) {
		var buf bytes.Buffer
		cfg := config{
			input:       strings.NewReader("Coding Challenges"),
			target:      "C",
			translation: "c",
			output:      &buf,
		}

		cfg.translateCmd()

		got := buf.String()
		want := "coding challenges"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Sub various letters and numbers, multiline", func(t *testing.T) {
		var buf bytes.Buffer
		cfg := config{
			input:       strings.NewReader("Coding Challenges\nhello123"),
			target:      "lo12",
			translation: "bo34",
			output:      &buf,
		}

		cfg.translateCmd()

		got := buf.String()
		want := "Coding Chabbenges\nhebbo343"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("emoji rune subst", func(t *testing.T) {
		var buf bytes.Buffer
		cfg := config{
			input:       strings.NewReader("hello😊"),
			target:      "😊",
			translation: "👀",
			output:      &buf,
		}

		cfg.translateCmd()

		got := buf.String()
		want := "hello👀"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("range expression", func(t *testing.T) {
		var buf bytes.Buffer
		cfg := config{
			input:       strings.NewReader("Coding Challenges\nhello123æøå"),
			target:      "a-d",
			translation: "e-h",
			output:      &buf,
		}

		cfg.translateCmd()

		got := buf.String()
		want := "Cohing Chellenges\nhello123æøå"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("special chars", func(t *testing.T) {
		var buf bytes.Buffer
		cfg := config{
			input:       strings.NewReader("Coding =%098*\nhello123æøå"),
			target:      "æ%",
			translation: "å=.",
			output:      &buf,
		}

		cfg.translateCmd()

		got := buf.String()
		want := "Coding ==098*\nhello123åøå"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("class specifier lower to upper", func(t *testing.T) {
		var buf bytes.Buffer
		cfg := config{
			input:       strings.NewReader("Coding Challenge"),
			target:      "[:lower:]",
			translation: "[:upper:]",
			output:      &buf,
		}

		cfg.translateCmd()

		got := buf.String()
		want := "CODING CHALLENGE"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("class specifier alpha to digit", func(t *testing.T) {
		var buf bytes.Buffer
		cfg := config{
			input:       strings.NewReader("Coding Challenge123%.?"),
			target:      "[:alpha:]",
			translation: "[:digit:]",
			output:      &buf,
		}

		// tr output: 299999 2999999991239999_9999.99.?
		// don't understand logic behind that, implemented my own mapping

		cfg.translateCmd()

		got := buf.String()
		want := "293896 270994964123%.?"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
