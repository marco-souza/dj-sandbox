package shared

import (
	"fmt"
)

type TimeRange struct {
	Start string
	End   string
}

// timerange to string to be used in the following command
// yt-dlp -f bestvideo+bestaudio --download-sections "*00:32:53.2-00:32:55.5" --force-keyframes-at-cuts https://www.youtube.com/watch?v=S9HdPi9Ikhk --output "%UserProfile%\Downloads\%(extractor)s/%(title)s [%(resolution)s] [%(id)s] [f%(format_id)s].%(ext)s" --ffmpeg-location %UserProfile%\Downloads\ffmpeg-5.1.1-full_build\bin\ffmpeg.exe
func (tr *TimeRange) String() string {
	return fmt.Sprintf("*%s-%s", tr.Start, tr.End)
}
