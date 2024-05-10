package parameters

import (
	"errors"
	"testing"

	"github.com/LiddleChild/findr/internal/errorwrapper"
)

func TestMaxDepth(t *testing.T) {
	tcs := []struct {
		name   string
		params []string
		werr   errorwrapper.ErrorWrapper
	}{
		{
			name:   "max depth neg",
			params: []string{"findr", "-mx", "-1"},
			werr:   errorwrapper.New(errorwrapper.Argument, errors.New("")),
		},
		{
			name:   "max depth zero",
			params: []string{"findr", "-mx", "0"},
			werr:   nil,
		},
		{
			name:   "max depth post",
			params: []string{"findr", "-mx", "1"},
			werr:   nil,
		},
		{
			name:   "max depth string",
			params: []string{"findr", "-mx", "abc"},
			werr:   errorwrapper.New(errorwrapper.Argument, errors.New("")),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_, _, werr := Parse(tc.params)

			if (werr != nil && tc.werr == nil) ||
				(werr == nil && tc.werr != nil) {
				t.Fatal("Error mismatched")
			} else if werr != nil && tc.werr != nil && werr.Type() != tc.werr.Type() {
				t.Fatal("Error type mismatched")
			}
		})
	}
}

func TestDirectory(t *testing.T) {
	tcs := []struct {
		name   string
		params []string
		werr   errorwrapper.ErrorWrapper
	}{
		{
			name:   "invalid path",
			params: []string{"findr", "-d", "thisfileshouldntexist"},
			werr:   errorwrapper.New(errorwrapper.Argument, errors.New("")),
		},
		{
			name:   "file",
			params: []string{"findr", "-d", "./go.mod"},
			werr:   errorwrapper.New(errorwrapper.Argument, errors.New("")),
		},
		{
			name:   "valid directory",
			params: []string{"findr", "-d", "."},
			werr:   nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_, _, werr := Parse(tc.params)

			if (werr != nil && tc.werr == nil) ||
				(werr == nil && tc.werr != nil) {
				t.Fatal("Error mismatched")
			} else if werr != nil && tc.werr != nil && werr.Type() != tc.werr.Type() {
				t.Fatal("Error type mismatched")
			}
		})
	}
}
