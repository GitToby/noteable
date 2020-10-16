package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var defaultConfDir string

const defaultConfDirName string = ".noteable"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "noteable",
	Version: "v0.1.0",
	Short:   "Take, manage and share notes powerfully.",
	Long:    "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defaultConfDir = filepath.Join(home, defaultConfDirName)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", filepath.Join(defaultConfDir, "config.yml"), "config file path to use")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// This will always ensure the config folder path exists on the host
	os.MkdirAll(defaultConfDir, os.ModePerm)

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetDefault("database_path", filepath.Join(defaultConfDir, "noteable.db"))

	// If a config file is found, read it in, else create one with defaults already set.
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No conf found, writing default config to", cfgFile)
		viper.WriteConfigAs(cfgFile)
	}
}
