// internal/tracepoint/format_file_parser.go
package tracepoint

import (
	"fmt"
	"strconv"
	"strings"
)

type Field struct {
	// Field represents a single field from the tracepoint format.
	Type   string // Field type
	Name   string // Field name
	Offset string // Offset within the tracepoint event structure
	Size   string // Size of the field
	Signed string // Whether the field is signed
}

// TracepointFormatData holds the complete parsed data from a tracepoint format file.
type TracepointFormatData struct {
	Name        string  // Tracepoint name
	Id          uint32  // Tracepoint ID
	Fields      []Field // Parsed fields from the format section
	PrintFormat string  // The print format string from the file
}

// ParseTracepointFormat parses the raw tracepoint format data into a TracepointFormatData struct.
// Example raw data:
//
//	name: block_bio_frontmerge
//	ID: 1252
//	format:
//	   field:unsigned short common_type;	offset:0;	size:2;	signed:0;
//	   field:unsigned char common_flags;	offset:2;	size:1;	signed:0;
//	   ...
//	print fmt: "%d,%d %s %llu + %u [%s]", ((unsigned int) ((REC->dev) >> 20)), ...
func ParseTracepointFormat(raw string) (TracepointFormatData, error) {
	var data TracepointFormatData
	var fields []Field

	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Parse the tracepoint name.
		if strings.HasPrefix(line, "name:") {
			data.Name = strings.TrimSpace(strings.TrimPrefix(line, "name:"))
			continue
		}

		// Parse the tracepoint ID.
		if strings.HasPrefix(line, "ID:") {
			idStr := strings.TrimSpace(strings.TrimPrefix(line, "ID:"))
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err != nil {
				return data, fmt.Errorf("failed to parse ID: %v", err)
			}
			data.Id = uint32(id)
			continue
		}

		// Parse the print format.
		if strings.HasPrefix(line, "print fmt:") {
			data.PrintFormat = strings.TrimSpace(strings.TrimPrefix(line, "print fmt:"))
			continue
		}

		// Parse field definitions.
		if strings.HasPrefix(line, "field:") {
			// Remove the "field:" prefix.
			fieldLine := strings.TrimSpace(strings.TrimPrefix(line, "field:"))
			// Split the line into parts based on semicolons.
			parts := strings.Split(fieldLine, ";")
			if len(parts) < 4 {
				// Not enough parts to be a valid field definition.
				continue
			}

			// First part should include the type and field name.
			firstPart := strings.TrimSpace(parts[0])
			typeAndName := strings.Fields(firstPart)
			if len(typeAndName) < 2 {
				continue
			}
			// The field type is all tokens except the last, which is the field name.
			fieldType := strings.Join(typeAndName[:len(typeAndName)-1], " ")
			fieldName := typeAndName[len(typeAndName)-1]

			offset := ""
			size := ""
			signed := ""

			// Parse the remaining parts.
			for _, part := range parts[1:] {
				part = strings.TrimSpace(part)
				if strings.HasPrefix(part, "offset:") {
					offset = strings.TrimSpace(strings.TrimPrefix(part, "offset:"))
				} else if strings.HasPrefix(part, "size:") {
					size = strings.TrimSpace(strings.TrimPrefix(part, "size:"))
				} else if strings.HasPrefix(part, "signed:") {
					signed = strings.TrimSpace(strings.TrimPrefix(part, "signed:"))
				}
			}

			field := Field{
				Type:   fieldType,
				Name:   fieldName,
				Offset: offset,
				Size:   size,
				Signed: signed,
			}
			fields = append(fields, field)
		}
	}
	data.Fields = fields
	return data, nil
}
