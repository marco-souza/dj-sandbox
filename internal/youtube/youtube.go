package youtube

import (
	"context"

	yt "github.com/lrstanley/go-ytdlp"
)

func DownloadVideo(url string, output string) {
	dl := yt.New().
		FormatSort("res,ext:mp4:m4a").
		RecodeVideo("mp4")

	if len(output) == 0 {
		dl.Output("%(extractor)s - %(title)s.%(ext)s")
	}

	_, err := dl.Run(context.TODO(), url)
	if err != nil {
		panic(err)
	}
}

func init() {
	yt.MustInstall(context.TODO(), &yt.InstallOptions{})
}
