package youtube

import (
	"context"
	"fmt"

	yt "github.com/lrstanley/go-ytdlp"
)

func DownloadAudio(url string, ext string) error {
	if len(ext) == 0 {
		ext = "flac"
	}

	// download builder
	dl := yt.New().
		ExtractAudio().
		AudioFormat(ext).
		AudioQuality("0").
		Output("%(playlist)s/%(title)s")

	// execute
	proc, err := dl.Run(context.TODO(), url)

	fmt.Println(proc.Stdout)

	if len(proc.Stderr) > 0 {
		fmt.Errorf("%s", proc.Stderr)
	}

	return err
}

func init() {
	yt.MustInstall(context.TODO(), &yt.InstallOptions{})
}
