// internal/tracepoint/format_file_print.go
package tracepoint

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
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

// Converts and returns the fields into a C-struct style representation.
func FormatAsCStruct(structName string, fields []Field) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("struct %s {\n", structName))
	for _, f := range fields {
		sb.WriteString(fmt.Sprintf("    %s %s;  // offset:%s; size:%s; signed:%s\n",
			f.Type, f.Name, f.Offset, f.Size, f.Signed))
	}
	sb.WriteString("};")
	return sb.String()
}

// Converts and returns a tabular representation of the fields.
func FormatAsTable(fields []Field) string {
	var sb strings.Builder
	sb.WriteString("FIELD\tTYPE\tOFFSET\tSIZE\tSIGNED\n")
	for _, f := range fields {
		sb.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n",
			f.Name, f.Type, f.Offset, f.Size, f.Signed))
	}
	return sb.String()
}
