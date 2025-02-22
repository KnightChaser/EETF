// internal/tracepoint/tracepoint.go
package tracepoint

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// ListTracepoints scans /sys/kernel/debug/tracing/events/*/*/format
// and returns a slice of tracepoints in the form "group/event".
func ListTracepoints() ([]string, error) {
	var tracepoints []string

	rootDir := "/sys/kernel/debug/tracing/events/"
	groups, err := ioutil.ReadDir(rootDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %v", rootDir, err)
	}

	// Iterate over reach group directory (e.g., "syscalls", "alarmtimer", etc.)
	for _, group := range groups {
		if group.IsDir() {
			groupPath := filepath.Join(rootDir, group.Name())
			events, err := ioutil.ReadDir(groupPath)
			if err != nil {
				// Skip grouops that we can't read
				continue
			}

			// Iterate over each event directory under the group
			for _, event := range events {
				if event.IsDir() {
					formatPath := filepath.Join(groupPath, event.Name(), "format")
					if _, err := ioutil.ReadFile(formatPath); err == nil {
						// Found a tracepoint with a format file, add it to the list
						tracepoints = append(tracepoints, fmt.Sprintf("%s/%s", group.Name(), event.Name()))
					}
				}
			}
		}
	}

	return tracepoints, nil
}
