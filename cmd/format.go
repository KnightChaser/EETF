// cmd/format.go
package cmd

import (
	"fmt"
	"log"

	"eetf/internal/tracepoint"
	"eetf/internal/tui"

	"github.com/spf13/cobra"
)

// A helper function if the given string array contains the given string.
func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

var outputFormat string

// formatCmd represents the "format" command.
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Print out the format of a selected tracepoint",
	Long: `Print out the format of a selected tracepoint.
This command will present an interactive fuzzy finder for you to select the tracepoint,
so you don't have to manually insert the tracepoint's name.

Output modes:
  raw   - Print the raw format file.
  c     - Print as a C struct.
  table - Print a tabular summary of the fields.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		AllowedOutputFormat := []string{"raw", "c", "table"}
		if !contains(AllowedOutputFormat, outputFormat) {
			log.Fatalf("Unsupported output format. Please choose one of: %v", AllowedOutputFormat)
		}

		// List all available tracepoints.
		tracepoints, err := tracepoint.ListTracepoints()
		if err != nil {
			log.Fatalf("Error listing tracepoints: %v", err)
		}

		if len(tracepoints) == 0 {
			fmt.Println("No tracepoints found.")
			return
		}

		// Use interactive search to select a tracepoint.
		selected := tui.InteractiveSelect(tracepoints)
		if selected == "" {
			fmt.Println("No tracepoint selected.")
			return
		}

		// Retrieve the raw format data.
		rawFormat, err := tracepoint.GetFormatData(selected)
		if err != nil {
			log.Fatalf("Error getting format data: %v", err)
		}

		// Parse the raw data using the enhanced parser.
		tpData, err := tracepoint.ParseTracepointFormat(rawFormat)
		if err != nil {
			log.Fatalf("Error parsing tracepoint format: %v", err)
		}

		// Output based on the selected mode.
		switch outputFormat {
		case "raw":
			fmt.Println(rawFormat)
		case "c":
			cStruct := tracepoint.FormatAsCStruct(tpData.Name, tpData.Fields)
			fmt.Println(cStruct)
		case "table":
			table := tracepoint.FormatAsTable(tpData.Fields)
			fmt.Println(table)
		default:
			fmt.Println("Unsupported output format. Please choose one of: raw, c, table")
		}
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
	formatCmd.Flags().StringVarP(&outputFormat, "output", "o", "raw", "Output format: raw, c, table")
}
