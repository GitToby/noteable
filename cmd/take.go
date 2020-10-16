package cmd

import (
	"fmt"
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
		inputPrompt := fmt.Sprintf("Note for %s:", today.Format("Mon 2 Jan 2006"))
		noteContent := internal.TakeInput(inputPrompt)
		if noteContent == "" {
			fmt.Println("No note given - nothing saved.")
		} else {
			note := internal.Note{Content: noteContent}
			note.Save()
		}
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
