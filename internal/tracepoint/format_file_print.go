// internal/tracepoint/format_file_print.go
package tracepoint

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
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
// This also highlights the C code using Chroma package.
func FormatAsCStruct(structName string, fields []Field) string {
	var sb strings.Builder

	// Start the struct definition.
	sb.WriteString(fmt.Sprintf("typedef struct %s {\n", structName))

	// First, determine the maximum length of field declarations ("type name;").
	//        It will be used to align the annotations.
	maxDeclLen := 0
	decls := make([]string, len(fields))
	for i, field := range fields {
		decl := fmt.Sprintf("%s %s;", field.Type, field.Name)
		decls[i] = decl
		if len(decl) > maxDeclLen {
			maxDeclLen = len(decl)
		}
	}

	// Append each field with an annotation aligned in the same column.
	for i, field := range fields {
		decl := decls[i]
		padding := strings.Repeat(" ", maxDeclLen-len(decl))
		// The two spaces before the annotation can be adjusted for your visual preference.
		line := fmt.Sprintf("    %s%s  // offset: %v, size: %v", decl, padding, field.Offset, field.Size)
		sb.WriteString(line + "\n")
	}

	// End the struct definition.
	sb.WriteString(fmt.Sprintf("} %s;\n", structName))

	code := sb.String()

	// Use Chroma to highlight the generated C code.
	lexer := lexers.Get("c")
	if lexer == nil {
		lexer = lexers.Fallback
	}
	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}
	style := styles.Get("colorful")

	if style == nil {
		style = styles.Fallback
	}

	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		// Fallback to unhighlighted code in case of an error.
		return code
	}
	var highlightedCode strings.Builder
	err = formatter.Format(&highlightedCode, style, iterator)
	if err != nil {
		return code
	}

	return highlightedCode.String()
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
