package cmd

import (
	"fmt"
	"time"

	"example.com/noteable/internal"
	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Get the history of notes you've taken",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		notes := internal.GetNotesSince(now.AddDate(0, 0, -7))
		for _, note := range notes {
			fmt.Println(note.CreatedAt.Format("Mon 02 Jan 2006 15:04:05"), "-", note.Content)
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags()
}
