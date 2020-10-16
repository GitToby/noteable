package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/noteable/internal"
	"github.com/spf13/cobra"
)

// takeCmd represents the take command
var takeCmd = &cobra.Command{
	Use:   "take",
	Short: "Take a note and save it",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		today := time.Now()
		fmt.Printf("Note for %s:\n", today.Format("Mon 2 Jan 2006"))
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		noteContent := scanner.Text()
		strings.Replace(noteContent, "\n", "", -1)
		strings.Replace(noteContent, "\r", "", -1)
		note := internal.Note{Content: noteContent}
		note.Save()
	},
}

func init() {
	rootCmd.AddCommand(takeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// takeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// takeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
