/*
Copyright © 2025 Marco Souza <marco@tremtec.com>
*/
package cmd

import (
	"fmt"
	"marco-souza/dj-sandbox/internal/shared"
	"marco-souza/dj-sandbox/internal/youtube"
	"os"

	"github.com/spf13/cobra"
)

var ext string
var start string
var end string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Short: "A CLI toolbox for DJing",
	Use:   "dj-sandbox <youtube url>",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		youtubeURL := args[0]
		fmt.Println("Downloading from URL:", youtubeURL)

		tr := shared.NewTimeRange(start, end)

		youtube.DownloadAudio(youtubeURL, ext, tr)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&ext, "ext", "x", "", "Specify audio extension (default: flac)")
	rootCmd.Flags().StringVarP(&start, "start", "s", "", "Specify start time (format: HH:MM:SS.ms)")
	rootCmd.Flags().StringVarP(&end, "end", "e", "", "Specify end time (format: HH:MM:SS.ms)")
}
