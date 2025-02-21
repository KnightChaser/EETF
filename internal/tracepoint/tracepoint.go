// internal/tracepoint/tracepoint.go
package tracepoint

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Scans /sys/kernel/debug/tracing/events/syscalls
// for tracepoints that have a format file.
func ListTracepoints() ([]string, error) {
	var tracepoints []string

	rootDir := "/sys/kernel/debug/tracing/events/syscalls"
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		return nil, fmt.Errorf("unable to read tracing events: %v", err)
	}

	// Find all syscall tracepoints that have a format file
	for _, file := range files {
		if file.IsDir() {
			formatFile := filepath.Join(rootDir, file.Name(), "format")
			if _, err := ioutil.ReadFile(formatFile); err == nil {
				tracepoints = append(tracepoints, file.Name())
			}
		}
	}
	return tracepoints, nil
}
