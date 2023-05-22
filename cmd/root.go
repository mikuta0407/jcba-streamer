/*
Copyright © 2023 mikuta0407
*/
package cmd

import (
	"log"
	"os"

	"github.com/mikuta0407/jcba-streamer/internal/jcba"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jcba-streamer",
	Short: "Stream jcba radio to stdout",
	Long: `Streaming ogg stream from jcba to stdout.
	example usage: jcba-streamer -s fmfukuro -d 3600 | ffmpeg -i pipe: -c libvorbis -acodec copy rec.opus`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting streamer...")
		station, _ := cmd.Flags().GetString("station")
		durationSec, _ := cmd.Flags().GetInt("duration")
		jcba.Main(station, durationSec)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("station", "s", "fmfukuro", "JCBA Station Name ex:'fmfukuro")
	rootCmd.Flags().IntP("duration", "d", 60, "Play duration time")
}
