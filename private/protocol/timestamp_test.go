// +build go1.7

package protocol

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	cases := map[string]struct {
		formatName     string
		expectedOutput string
		input          time.Time
	}{
		"UnixTest1": {
			formatName:     UnixTimeFormatName,
			expectedOutput: "946845296",
			input:          time.Date(2000, time.January, 2, 20, 34, 56, .123e9, time.UTC),
		},
		"ISO8601Test1": {
			formatName:     ISO8601TimeFormatName,
			expectedOutput: "2000-01-02T20:34:56Z",
			input:          time.Date(2000, time.January, 2, 20, 34, 56, .123e9, time.UTC),
		},
		"RFC822Test1": {
			formatName:     RFC822TimeFormatName,
			expectedOutput: "Sun, 02 Jan 2000 20:34:56 GMT",
			input:          time.Date(2000, time.January, 2, 20, 34, 56, 0, time.UTC),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if v, _ := FormatTime(c.formatName, c.input); v != c.expectedOutput {
				t.Errorf("input %v \n and output: %v, \n don't match for %s format ", c.input, c.expectedOutput, c.formatName)
			}
		})
	}
}

func TestParseTime(t *testing.T) {

	// Assert equals for input and output works for a precision upto three decimal places
	cases := map[string]struct {
		formatName, input string
		expectedOutput    time.Time
	}{
		"UnixTest1": {
			formatName:     UnixTimeFormatName,
			input:          "946845296.123",
			expectedOutput: time.Date(2000, time.January, 2, 20, 34, 56, .123e9, time.UTC),
		},
		"UnixTest2": {
			formatName:     UnixTimeFormatName,
			input:          "946845296.12344",
			expectedOutput: time.Date(2000, time.January, 2, 20, 34, 56, .123e9, time.UTC),
		},
		"UnixTest3": {
			formatName:     UnixTimeFormatName,
			input:          "946845296.1229999",
			expectedOutput: time.Date(2000, time.January, 2, 20, 34, 56, .123e9, time.UTC),
		},
		"ISO8601Test1": {
			formatName:     ISO8601TimeFormatName,
			input:          "2000-01-02T20:34:56.123Z",
			expectedOutput: time.Date(2000, time.January, 2, 20, 34, 56, .123e9, time.UTC),
		},
		"ISO8601Test2": {
			formatName:     ISO8601TimeFormatName,
			input:          "2000-01-02T20:34:56.123456789Z",
			expectedOutput: time.Date(2000, time.January, 2, 20, 34, 56, .123456789e9, time.UTC),
		},
		"RFC822Test1": {
			formatName:     RFC822TimeFormatName,
			input:          "Sun, 2 Jan 2000 20:34:56 GMT",
			expectedOutput: time.Date(2000, time.January, 2, 20, 34, 56, 0, time.UTC),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			timeVal, err := ParseTime(c.formatName, c.input)
			if err != nil {
				t.Errorf("unable to parse time, %v", err)
			}
			if timeVal.UTC() != c.expectedOutput {
				t.Errorf("input: %v \n and output time %v,\n don't match for %s format ", c.input, c.expectedOutput, c.formatName)
			}
		})
	}
}
