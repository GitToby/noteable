package cmd

import (
	"fmt"

	"example.com/noteable/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var phraseString string
var all bool
var limit int

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search all your notes for a string",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		var queryString string
		if phraseString == "" && !all {
			queryString = internal.TakeInput("Search query:", false)
		} else {
			queryString = phraseString
		}

		if queryString == "" && !all {
			fmt.Println("No search param given. To retrieve all notes pass flag --all")
			return
		}

		notes := internal.SearchNotes(queryString, limit)
		if len(notes) > 0 {
			for _, note := range notes {
				fmt.Println(note.CreatedAt.Format(viper.GetString("date_format")), "-", note.Content)
			}
		} else {
			fmt.Printf("No results for \"%s\"", queryString)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&phraseString, "phrase", "p", "", "Pass the search phrase directly to the command line rather than interactively")
	searchCmd.Flags().BoolVarP(&all, "all", "a", false, "Return all notes")
	searchCmd.Flags().IntVarP(&limit, "limit", "l", 10, "The max number of notes to return")
}
