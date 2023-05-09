/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/shubydo/simple-server/internal/server"
)

var port int

// startCmd represents the start command.
var startCmd = &cobra.Command{
	Use:       "start",
	ValidArgs: []string{"start"},

	Example: "simple-server start",
	Short:   "Start the server",
	Long: `Start the server and listen for requests.

Example:
	simple-server start
	simple-server start --port 8080
	simple-server start --port 8080 --verbose
	simple-server start --port 8080 --verbose --debug
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// parse flags
		s := server.New(
			server.WithPort(port),
		)

		addr := fmt.Sprintf(":%d", port)
		log.Printf("Starting server on %s", addr)
		log.Fatalf("Server failed: %v", http.ListenAndServe(addr, s))
	},
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// parse flags
	// 	cmd.Flags().Parse(args)

	// 	// Check if port is passed in
	// 	// If not, use default port
	// 	// If yes, use port passed in

	// 	// Pass in the port
	// 	s := server.New(
	// 		server.WithPort(port),
	// 	)
	// 	log.Printf("Starting server on port: %s", "8080")
	// 	log.Fatalf("Server failed: %v", http.ListenAndServe(fmt.Sprintf(":%s", "8080"), s))
	// },
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.
	// Write flag for port

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.PersistentFlags().IntVarP(&port, "port", "p", server.DefaultPort, "The port to listen on")
}
