package youtube

import (
	"context"

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
		Output("%(extractor)s - %(title)s")

	// execute
	_, err := dl.Run(context.TODO(), url)

	return err
}

func init() {
	yt.MustInstall(context.TODO(), &yt.InstallOptions{})
}
