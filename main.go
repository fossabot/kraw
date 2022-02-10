package main

import (
	"fmt"
	"runtime"

	"github.com/aw83/kraw/commands"
)

var (
	version = "snapshot"
	date    = ""
)

func main() {
	commands.Execute(makeVersion(version, date))
}

func makeVersion(version, date string) string {
	v := version
	if date != "" {
		v = fmt.Sprintf("%s\nbuilt at: %s", v, date)
	}
	v = fmt.Sprintf("%s\ngoos: %s\ngoarch: %s\ncompiler: %s\ngo version: %s",
		v, runtime.GOOS, runtime.GOARCH, runtime.Compiler, runtime.Version())
	return v
}
