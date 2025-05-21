package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type config struct {
	input           io.Reader
	subst           map[rune]rune
	deleteFlag      bool
	checkFunc       checkFunc
	translateFunc   translateFunc
	target          string
	translation     string
	targetType      expressionType
	translationType expressionType
	output          io.Writer
}

type expressionType string

const (
	Regular  expressionType = "regular"
	Range    expressionType = "range"
	Function expressionType = "function"
)

type (
	checkFunc     func(rune) bool
	translateFunc func(rune) rune
)

type substFuncs struct {
	check     checkFunc
	translate translateFunc
}

var specifierFuncMap = map[string]substFuncs{
	"alpha": {unicode.IsLetter, nil},
	"upper": {unicode.IsUpper, unicode.ToUpper},
	"lower": {unicode.IsLower, unicode.ToLower},
	"digit": {unicode.IsDigit, ToDigit},
	"print": {unicode.IsPrint, nil},
	"punct": {unicode.IsPunct, nil},
	"space": {unicode.IsSpace, nil},
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Printf("couldn't load config %v", err)
		os.Exit(1)
	}

	cfg.translateCmd()

	// fmt.Println(res)
}

func loadConfig() (config, error) {
	cfg := config{
		subst:       make(map[rune]rune),
		input:       os.Stdin,
		target:      "",
		translation: "",
		output:      os.Stdout,
	}
	flag.BoolVar(&cfg.deleteFlag, "d", false, "delete chosen chars")

	flag.Parse()
	args := flag.Args()

	switch {
	case len(args) == 1 && cfg.deleteFlag:
		cfg.target = args[0]
	case len(args) < 2:
		return cfg, fmt.Errorf("please provide chars to translate and chars to translate into: %v", args)
	case len(args) == 2:
		cfg.target = args[0]
		cfg.translation = args[1]
	default:
		return cfg, fmt.Errorf("please provide cmd <target> <translation>: %v", args)
	}

	return cfg, nil
}

func (cfg *config) translateCmd() {
	cfg.checkAndLoadExpression()

	scanner := bufio.NewScanner(cfg.input)
	scanner.Split(bufio.ScanLines)

	firstLine := true

	for scanner.Scan() {
		if !firstLine {
			fmt.Fprintln(cfg.output)
		} else {
			firstLine = false
		}

		line := scanner.Text()
		processedLine := processRunes(line, cfg.subst)
		fmt.Fprint(cfg.output, processedLine)
	}
}

func processRunes(line string, trMap map[rune]rune) string {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanRunes)

	var res strings.Builder

	for scanner.Scan() {
		currentRune := []rune(scanner.Text())[0]
		val, exists := trMap[currentRune]
		if exists {
			res.WriteRune(val)
			continue
		}
		res.WriteRune(currentRune)

	}
	return res.String()
}

func loadSubstitution(target, translation string) map[rune]rune {
	res := make(map[rune]rune)

	targetRunes := []rune(target)
	translationRunes := []rune(translation)

	for i, r := range targetRunes {
		if i < len(translationRunes) {
			res[r] = translationRunes[i]
		} else {
			res[r] = translationRunes[len(translationRunes)-1]
		}
	}
	return res
}

func expandRange(s string) string {
	var res []byte
	idx := strings.Index(s, "-")

	startTarget := s[idx-1]
	endTarget := s[idx+1]
	i := startTarget

	for i <= endTarget {
		res = append(res, byte(i))
		i++
	}
	return string(res)
}

func validRangeSubstitution(s string) bool {
	idx := strings.Index(s, "-")
	return idx != -1 && len(s) >= 3 && idx > 0 && idx < len(s)-1
}

func (cfg *config) checkAndLoadExpression() {
	cfg.targetType = checkExpression(cfg.target)
	cfg.translationType = checkExpression(cfg.translation)

	if cfg.targetType == Range {
		cfg.target = expandRange(cfg.target)
	}

	if cfg.translationType == Range {
		cfg.translation = expandRange(cfg.translation)
	}

	cfg.subst = loadSubstitution(cfg.target, cfg.translation)
}

func checkExpression(s string) expressionType {
	classPattern := `^\[:[a-z]+\:]$`
	re := regexp.MustCompile(classPattern)

	switch {
	case re.MatchString(s):
		return Function
	case validRangeSubstitution(s):
		return Range
	default:
		return Regular
	}
}
