package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DefaultConfDir is the config dir for storing the related notable files (conf, database, ...)
var DefaultConfDir = getDefaultConfDir()

// CfgFile is the configuration file for storing config
var CfgFile string

const defaultConfDirName string = ".noteable"

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set and retrieve config for notable",
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getDefaultConfDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return filepath.Join(home, defaultConfDirName)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// This will always ensure the config folder path exists on the host
	os.MkdirAll(DefaultConfDir, os.ModePerm)

	viper.AutomaticEnv() // read in environment variables that match

	// These match the config in the cmd/config.go data
	viper.SetDefault("database_path", filepath.Join(DefaultConfDir, "noteable.db"))
	viper.SetDefault("date_format", "Mon 02 Jan 2006 15:04:05")
	viper.SetDefault("share.methods", []string{"email"})
	viper.SetDefault("share.schedule", 5)
	viper.SetDefault("share.date_format", "Mon 02 Jan 03:04 pm")
	viper.SetDefault("share.email.to", []string{""})
	viper.SetDefault("share.email.cc", []string{""})
	viper.SetDefault("share.email.subject", "Notes from noteable!")

	// If a config file is found, read it in, else create one with defaults already set.
	viper.SetConfigFile(CfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No config found, writing default config to", CfgFile)
		viper.WriteConfigAs(CfgFile)
	}
}
