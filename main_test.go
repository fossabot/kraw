package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeVersion(t *testing.T) {
	os := fmt.Sprintf("\ngoos: %s\ngoarch: %s\ncompiler: %s\ngo version: %s",
		runtime.GOOS, runtime.GOARCH, runtime.Compiler, runtime.Version())
	tests := []struct {
		name, version, date, result string
	}{
		{
			name:   "all empty",
			result: os,
		},
		{
			name:    "only version",
			version: "0.1",
			result:  "0.1" + os,
		},
		{
			name:    "version and date",
			version: "0.2",
			date:    "2020-1-1",
			result:  "0.2\nbuilt at: 2020-1-1" + os,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver := makeVersion(tt.version, tt.date)
			assert.Equal(t, tt.result, ver)
		})
	}
}
