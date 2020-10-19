package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
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

// OpenBrowser opens the browser of the given system
func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}