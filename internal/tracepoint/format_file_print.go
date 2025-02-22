// internal/tracepoint/format_file_print.go
package tracepoint

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Obtain the format data of a specific tracepoint path.
// It just reads the raw string data from the format file, by itself.
func GetFormatData(tracepoint string) (string, error) {
	rootDir := "/sys/kernel/debug/tracing/events"
	filePath := filepath.Join(rootDir, tracepoint, "format")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read format file: %v", err)
	}
	return string(data), nil
}
