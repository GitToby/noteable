package cmd

import (
	"fmt"

	"example.com/noteable/internal"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search all your notes for a string",
	Run: func(cmd *cobra.Command, args []string) {
		queryString := internal.TakeInput("Search query:")
		notes := internal.SearchNotes(queryString, 10)
		if len(notes) > 0 {
			for _, note := range notes {
				fmt.Println(note.CreatedAt.Format("Mon 02 Jan 2006 15:04:05"), "-", note.Content)
			}
		} else {
			fmt.Printf("No results for search \"%s\"", queryString)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
