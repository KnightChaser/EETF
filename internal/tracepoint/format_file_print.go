// internal/tracepoint/format_file_print.go
package tracepoint

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
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
func PrintFormatAsTable(data TracepointFormatData) {
	// Print Metadata
	fmt.Println("Tracepoint Name:", data.Name)
	fmt.Println("Tracepoint ID:  ", data.Id)
	// fmt.Println("Print Format:   ", data.PrintFormat)		// Seems unnecessary for now. Later, maybe.

	// Create custom formatters for header and first column.
	headerFmt := color.New(color.FgCyan, color.Underline, color.Bold).SprintfFunc()
	columnFmt := color.New(color.FgHiYellow).SprintfFunc()

	// Create the table with headers.
	tbl := table.New("Field Name", "Type", "Size", "Offset", "Signed")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Add rows for each field.
	for _, field := range data.Fields {
		tbl.AddRow(field.Name, field.Type, field.Size, field.Offset, field.Signed)
	}

	// Capture the printed table output by redirecting the writer, or simply print it.
	// For simplicity, we'll print directly:
	tbl.Print()
}
