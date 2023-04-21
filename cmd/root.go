/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use: "simple-server",

	Short: "Simple server written in Go",
	Long: `Simple server written in Go using Cobra and Viper.
This server is meant to be used as a template for future projects.

Example:
	simple-server start
	simple-server start --port 8080
	simple-server start --port 8080 --verbose
	simple-server start --port 8080 --verbose --debug
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Show help if no arguments are passed
	if err := rootCmd.Execute(); err != nil {
		err := rootCmd.Help()
		if err != nil {
			os.Exit(1)
		}
	}
}

//nolint:gochecknoinits
func init() {
	var verbose, debug bool
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.simple-server.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
