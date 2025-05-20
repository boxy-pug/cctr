package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	text       io.Reader
	subst      map[string]string
	deleteFlag bool
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Printf("couldn't load config %v", err)
		os.Exit(1)
	}

	res := Substitute(cfg)

	fmt.Println(res)
}

func loadConfig() (config, error) {
	var err error

	cfg := config{
		subst: make(map[string]string),
		text:  os.Stdin,
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

func Substitute(cfg config) string {
	scanner := bufio.NewScanner(cfg.text)
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

	if len(target) != len(translation) {
		return res, fmt.Errorf("unequal length")
	}

	for i := range target {
		res[string(target[i])] = string(translation[i])
	}

	return res, nil
}
