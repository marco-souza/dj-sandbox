package shared_test

import (
	s "marco-souza/dj-sandbox/internal/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeRangeString(t *testing.T) {
	tests := []struct {
		name     string
		tr       *s.TimeRange
		expected string
	}{
		{
			name:     "Valid time range",
			tr:       &s.TimeRange{Start: "00:32:53.2", End: "00:32:55.5"},
			expected: "*00:32:53.2-00:32:55.5",
		},
		{
			name:     "Start and end are the same",
			tr:       &s.TimeRange{Start: "01:00:00.0", End: "01:00:00.0"},
			expected: "*01:00:00.0-01:00:00.0",
		},
		{
			name:     "Different start and end",
			tr:       &s.TimeRange{Start: "00:00:00.0", End: "00:01:00.0"},
			expected: "*00:00:00.0-00:01:00.0",
		},
		{
			name:     "Empty start and end",
			tr:       s.NewTimeRange("", ""),
			expected: "*00:00:00.0-inf",
		},
		{
			name:     "Empty start",
			tr:       s.NewTimeRange("", "00:10:00.0"),
			expected: "*00:00:00.0-00:10:00.0",
		},
		{
			name:     "Empty end",
			tr:       s.NewTimeRange("00:05:00.0", ""),
			expected: "*00:05:00.0-inf",
		},
		{
			name:     "Without miliseconds",
			tr:       &s.TimeRange{Start: "01:00", End: "02:00"},
			expected: "*01:00-02:00",
		},
		{
			name:     "Without hours",
			tr:       &s.TimeRange{Start: "00:30.0", End: "00:45.0"},
			expected: "*00:30.0-00:45.0",
		},
	}

	for _, tt := range tests {
		tr := tt.tr

		t.Run(tt.name, func(t *testing.T) {
			got := tr.String()
			expected := tt.expected

			assert.Equal(t, expected, got)
		})
	}
}
