package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// TODO: all these should be modifiable via input args
	indentSymbol := "    "
	indentCurrent := 0
	textOpen := false
	prevChar := ""

	reader := bufio.NewReader(os.Stdin)

	indentKeys := map[string][]string{
		"indent":        []string{"{", "["},
		"unindent":      []string{"}", "]"},
		"newLine":       []string{","},
		"space":         []string{":"},
		"textEnclosure": []string{"\""},
	}

	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Print(err)
			}
		} else {

			if r, _ := in_array(string(c), indentKeys["textEnclosure"]); r && !textOpen {
				textOpen = true
			} else if r, _ := in_array(string(c), indentKeys["textEnclosure"]); r && textOpen && prevChar != "\\" {
				textOpen = false
			}

			if !textOpen {
				if r, _ := in_array(string(c), indentKeys["indent"]); r {
					indentCurrent++
					fmt.Print(string(c) + "\n" + strings.Repeat(indentSymbol, indentCurrent))
				} else if r, _ := in_array(string(c), indentKeys["unindent"]); r {
					indentCurrent--
					if indentCurrent < 0 {
						indentCurrent = 0
					}
					fmt.Print("\n" + strings.Repeat(indentSymbol, indentCurrent) + string(c))
				} else if r, _ := in_array(string(c), indentKeys["newLine"]); r {
					fmt.Print(string(c) + "\n" + strings.Repeat(indentSymbol, indentCurrent))
				} else if r, _ := in_array(string(c), indentKeys["space"]); r {
					fmt.Print(string(c) + " ")
				} else {
					fmt.Print(string(c))
				}
			} else {
				fmt.Print(string(c))
			}
			prevChar = string(c)
		}
	}
}

func in_array(val string, inpArray []string) (ok bool, i int) {
	for i = range inpArray {
		if ok = inpArray[i] == val; ok {
			return
		}
	}
	return
}
