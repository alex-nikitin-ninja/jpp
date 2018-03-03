package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rawJson, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error:")
		fmt.Print(err)
		os.Exit(1)
	} else {
		fmt.Print(rawJson)
	}
}
