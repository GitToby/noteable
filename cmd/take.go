package cmd

import (
	"fmt"
	"time"

	"example.com/noteable/internal"
	"github.com/spf13/cobra"
)

var noteContent string

// takeCmd represents the take command
var takeCmd = &cobra.Command{
	Use:     "take",
	Short:   "Take a note and save it",
	Aliases: []string{"t"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		today := time.Now()
		if noteContent == "" {
			inputPrompt := fmt.Sprintf("Note for %s:", today.Format("Mon 2 Jan 2006"))
			noteContent = internal.TakeInput(inputPrompt)
		}
		if noteContent == "" {
			fmt.Println("No note given - nothing saved.")
		} else {
			note := internal.Note{Content: noteContent}
			note.Save()
		}
		if today.Weekday() == 5 {
			shareCmd.Run(cmd, args)
		}
	},
}

func init() {
	rootCmd.AddCommand(takeCmd)
	takeCmd.Flags().StringVarP(&noteContent, "message", "m", "", "Pass a message directly in the commandline rather than interactively")
}
