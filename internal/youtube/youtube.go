package youtube

import (
	"context"
	"fmt"
	"marco-souza/dj-sandbox/internal/shared"

	yt "github.com/lrstanley/go-ytdlp"
)

func DownloadAudio(url string, ext string, tr *shared.TimeRange) error {
	if len(ext) == 0 {
		ext = "flac"
	}

	// download builder
	dl := yt.New().
		ExtractAudio().
		AudioFormat(ext).
		AudioQuality("0").
		Output("%(playlist)s/%(title)s")

	if tr != nil {
		dl.
			DownloadSections(tr.String()).
			ForceKeyframesAtCuts()
	}

	// execute
	proc, err := dl.Run(context.TODO(), url)
	if err != nil {
		return fmt.Errorf("%s - %s", err, proc.Stderr)
	}

	fmt.Println(proc.Stdout)
	return nil
}

func init() {
	yt.MustInstall(context.TODO(), &yt.InstallOptions{})
}
