/*
Copyright © 2025 Marco Souza <marco@tremtec.com>
*/
package cmd

import (
	"fmt"
	"marco-souza/dj-sandbox/internal/youtube"
	"os"

	"github.com/spf13/cobra"
)

var outputName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Short: "A CLI toolbox for DJing",
	Use:   "dj-sandbox <youtube url>",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		youtubeURL := args[0]
		handleDownload(youtubeURL, outputName)
	},
}

func handleDownload(youtubeURL string, outputName string) {
	fmt.Println("Downloading from URL:", youtubeURL)

	youtube.DownloadVideo(youtubeURL, outputName)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&outputName, "output", "o", "", "Specify the output file name")
}
