package cmd

import (
	"fmt"
	"time"

	"example.com/noteable/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "github.com/spf13/viper"
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
		force := noteContent != "" // if a param is passed in, force no input
		if !force {
			inputPrompt := fmt.Sprintf("Note for %s:", today.Format("Mon 2 Jan 2006"))
			noteContent = internal.TakeInput(inputPrompt, false)
		}

		if noteContent == "" {
			fmt.Println("Cancelled - nothing saved.")
			return
		}

		note := internal.Note{Content: noteContent}
		note.Save()
		fmt.Println("Saved.")
		shareDay := viper.Get("share.schedule")
		if int(today.Weekday()) == shareDay {
			if force {
				shareCmd.Run(cmd, args)
			} else if share := internal.TakeInput("Today is a share day, share with audience?", true); share == "y" {
				shareCmd.Run(cmd, args)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(takeCmd)
	takeCmd.Flags().StringVarP(&noteContent, "message", "m", "", "Pass a message directly in the commandline rather than interactively")
}
