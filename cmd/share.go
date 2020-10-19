package cmd

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"example.com/noteable/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var daysToLookBackShare int

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		includedNotes := internal.GetNotesSince(now.AddDate(0, 0, -1*daysToLookBack))

		if len(includedNotes) < 1 {
			fmt.Println("Looks like you have no notes to share, make some using \"noteable take\"")
			return
		}
		firstNoteDate := includedNotes[0].CreatedAt.Format("Mon 01 Jan")
		sort.Slice(includedNotes, func(i, j int) bool {
			return includedNotes[i].CreatedAt.Before(includedNotes[j].CreatedAt)
		})

		emailList := strings.Join(viper.GetStringSlice("share.email.to"), ";")
		if emailList == "" {
			fmt.Println("No emails to send to, canceling share")
			return
		}
		ccList := strings.Join(viper.GetStringSlice("share.email.cc"), ";")
		subject := viper.GetString("share.email.subject")
		body := fmt.Sprintf("Notes since %s:\n\n", firstNoteDate)

		dateFmt := viper.GetString("share.date_format")
		for _, note := range includedNotes {
			body += fmt.Sprintf("%s: %s\n", note.CreatedAt.Format(dateFmt), note.Content)
		}

		mailURL := fmt.Sprintf("mailto:?to=%s&cc=%s&subject=%s&body=%s", emailList, ccList, url.QueryEscape(subject), url.QueryEscape(body))
		println(mailURL)
		internal.OpenBrowser(mailURL)
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.Flags().IntVarP(&daysToLookBackShare, "days", "d", 7, "period of days of history to share")
}
