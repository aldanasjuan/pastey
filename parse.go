package main

import (
	"aldanasjuan/pastey/codes"
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"strings"
)

var modifiers = map[string]int{
	"alt":   ModAlt,
	"ctrl":  ModCtrl,
	"shift": ModShift,
	"win":   ModWin,
}

var nums = map[string]string{
	"0": "zero",
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
}

func ParseFile(path string) ([]Hotkey, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	l := 0
	keys := []Hotkey{{-1, ModAlt, codes.GetKey("backspace"), []byte{}, "alt+backspace"}}
	for scanner.Scan() {
		l++
		s := bytes.TrimSpace(scanner.Bytes())
		if len(s) < 1 || s[0] == '#' {
			continue
		}

		split := strings.SplitN(string(s), "=", 2)
		if len(split) != 2 {
			fmt.Printf("error at line %v: %v %+v\n", l, "Bad format. Expecting alt+ctrl+a.", split)
			continue
		}
		mods := strings.Split(split[0], "+")
		if len(mods) < 2 {
			fmt.Printf("error at line %v: %v\n", l, "key should have at least one modifier. options are [ctrl, alt, shift, win]. ")
			continue
		}
		code := strings.ToLower(mods[len(mods)-1])
		if s, ok := nums[code]; ok {
			code = s
		}
		mods = mods[:len(mods)-1]
		codeV := codes.GetKey(code)
		if codeV == -1 {
			fmt.Printf("error at line %v: %v %v\n", l, "unkown key code", code)
			continue
		}
		key := Hotkey{}
		key.KeyCode = codeV
		key.Id = l
		key.Value = []byte(split[1])
		key.Text = split[0]
		for _, m := range mods {
			if v, ok := modifiers[m]; ok {
				key.Modifiers += v
			}
		}
		if key.Modifiers > 0 {
			keys = append(keys, key)
		}
	}

	return keys, nil
}
