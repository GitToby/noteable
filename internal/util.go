package internal

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os/exec"
	"runtime"
)

// TakeInput from the cli with a set of prompts
func TakeInput(promptLabel string, isConfirm bool) string {
	res := promptui.Prompt{
		Label:     promptLabel,
		IsConfirm: isConfirm,
	}
	result, err := res.Run()
	if err != nil {
		return ""
	}
	return result
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
