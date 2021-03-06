package cmd

import (
	"fmt"
	"os"

	"example.com/noteable/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var force bool

// resetCmd represents the remove command
var resetCmd = &cobra.Command{
	Use:       "reset",
	Short:     "A brief description of your command",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: viper.AllKeys(),
	Run: func(cmd *cobra.Command, args []string) {
		if !force {
			res := internal.TakeInput("Are you sure?", true)
			if res != "y" {
				return
			}
		}
		err := os.Remove(CfgFile)
		if err != nil {
			fmt.Println("Cannot remove original notable config file, continuing to write to", CfgFile)
		} else {
			fmt.Println("Removed config file, run \"noteable init\" to write a fresh configuration file")
		}
	},
}

func init() {
	configCmd.AddCommand(resetCmd)
	resetCmd.Flags().BoolVarP(&force, "force", "f", false, "Force reset the config to defaults ")
}
