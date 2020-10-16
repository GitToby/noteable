package cmd

import (
	"fmt"
	"time"

	"example.com/noteable/internal"
	"github.com/spf13/cobra"
)

// defalt to look back 7 days worth of history
var daysToLookBack int

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Get the history of notes you've taken",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		notes := internal.GetNotesSince(now.AddDate(0, 0, -1*daysToLookBack))
		if len(notes) > 0 {
			for _, note := range notes {
				fmt.Println(note.CreatedAt.Format("Mon 02 Jan 2006 15:04:05"), "-", note.Content)
			}
		} else {
			fmt.Println("Looks like you have no notes, make some using \"noteable take\"")
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().IntVarP(&daysToLookBack, "days", "d", 7, "number of days of history to return")
}
