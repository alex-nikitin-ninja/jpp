// go run hello-world.go
// go build hello-world.go

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
		fmt.Println("err:")
		fmt.Print(err)
		// return 1
	} else {
		fmt.Print(rawJson)
	}

}
