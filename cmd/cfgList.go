package cmd

import (
	"fmt"
	"github.com/spf13/viper"
"sort"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		t := table.NewWriter()
		t.AppendHeader(table.Row{"Key", "Value"})
		keys := viper.AllKeys()
		sort.Strings(keys)
		for _, key := range keys {
			t.AppendRow(table.Row{key, viper.Get(key)})
		}
		fmt.Println(t.Render())
	},
}

func init() {
	configCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
