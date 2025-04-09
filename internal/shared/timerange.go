package shared

import (
	"fmt"
)

type TimeRange struct {
	Start string
	End   string
}

func NewTimeRange(start, end string) *TimeRange {
	if start == "" {
		start = "00:00:00.0"
	}

	if end == "" {
		end = "inf"
	}

	return &TimeRange{
		Start: start,
		End:   end,
	}
}

func (tr *TimeRange) String() string {
	return fmt.Sprintf("*%s-%s", tr.Start, tr.End)
}
