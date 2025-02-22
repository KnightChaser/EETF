// cmd/format.go
package cmd

import (
	"fmt"
	"log"

	"eetf/internal/tracepoint"
	"eetf/internal/tui"

	"github.com/spf13/cobra"
)

// ./eetf format
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Print out the format of a selected tracepoint",
	Long: `Print out the format of a selected tracepoint. This command will
		   first ask you to select a tracepoint from the list of available.
		   Then, it will print out the format of the selected tracepoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		tracepoints, err := tracepoint.ListTracepoints()
		if err != nil {
			log.Fatalf("Error listing tracepoints: %v", err)
		}

		if len(tracepoints) == 0 {
			fmt.Println("No tracepoints found.")
			return
		}

		// Print out the format of the selected tracepoint
		result := tui.InteractiveSelect(tracepoints)
		formatData, err := tracepoint.GetFormatData(result)
		if err != nil {
			log.Fatalf("Error getting format data: %v", err)
		}

		fmt.Println(formatData)
	},
}
