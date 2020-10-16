package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TakeInput from the cli with a set of prompts
func TakeInput(prompts ...string) string {
	for _, prompt := range prompts {
		fmt.Println(prompt)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	// replace newline chars for all systems
	strings.Replace(userInput, "\n", "", -1)
	strings.Replace(userInput, "\r", "", -1)

	return userInput
}
