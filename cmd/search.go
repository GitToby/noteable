package cmd

import (
	"fmt"

	"example.com/noteable/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var queryString string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search all your notes for a string",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		if queryString == "" {
			queryString = internal.TakeInput("Search query:")
		}
		notes := internal.SearchNotes(queryString, 10)
		if len(notes) > 0 {
			for _, note := range notes {
				fmt.Println(note.CreatedAt.Format(viper.GetString("date_format")), "-", note.Content)
			}
		} else {
			fmt.Printf("No results for search \"%s\"", queryString)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&queryString, "phrase", "p", "", "Pass the search phrase directly to the command line rather than interactively")
}
