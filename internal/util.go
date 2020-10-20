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
func TakeInput(promptLabel string, isConfirm bool) string {
	reader := bufio.NewReader(os.Stdin)
	if isConfirm {
		for {
			fmt.Print(promptLabel, " [y/n] ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			text = strings.ToLower(text)
			strings.ReplaceAll(text, "yes", "y")
			strings.ReplaceAll(text, "no", "n")
			if text == "y" || text == "n" {
				return text
			}
		}
	} else {
		fmt.Print(promptLabel, " ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		return text
	}
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
