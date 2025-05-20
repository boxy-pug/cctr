package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type config struct {
	input      io.Reader
	subst      map[string]string
	deleteFlag bool
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Printf("couldn't load config %v", err)
		os.Exit(1)
	}

	processLines(cfg)

	// fmt.Println(res)
}

func loadConfig() (config, error) {
	var err error

	cfg := config{
		subst: make(map[string]string),
		input: os.Stdin,
	}
	flag.BoolVar(&cfg.deleteFlag, "d", false, "delete chosen chars")

	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		return cfg, fmt.Errorf("please provide chars to translate and chars to translate into: %v", args)
	}

	cfg.subst, err = loadSubstitution(args[0], args[1])
	if err != nil {
		return cfg, fmt.Errorf("error parsing subst args %q and %q", args[0], args[1])
	}

	return cfg, nil
}

func processLines(cfg config) string {
	scanner := bufio.NewScanner(cfg.input)
	scanner.Split(bufio.ScanLines)

	res := ""

	for scanner.Scan() {
		processedLine := processRunes(scanner.Text(), cfg)
		fmt.Println(processedLine)
		res += processedLine + "\n"
	}

	return strings.TrimSuffix(res, "\n")
}

func processRunes(line string, cfg config) string {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanRunes)
	res := ""

	for scanner.Scan() {
		currentRune := scanner.Text()
		val, exists := cfg.subst[currentRune]
		if exists {
			res += val
			continue
		}
		res += currentRune

	}
	return res
}

func loadSubstitution(target, translation string) (map[string]string, error) {
	res := make(map[string]string)

	for i := range len(target) {
		if i < len(translation) {
			res[string(target[i])] = string(translation[i])
		} else {
			res[string(target[i])] = string(translation[len(translation)-1])
		}
	}
	return res, nil
}

func maxLength(a, b int) int {
	if a > b {
		return a
	}
	return b
}
