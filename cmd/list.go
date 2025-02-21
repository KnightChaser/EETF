// cmd/list.go
package cmd

import (
	"fmt"
	"log"

	"eetf/internal/tracepoint"
	"eetf/internal/tui"

	"github.com/spf13/cobra"
)

// ./eetf list
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available tracepoints",
	Long:  `Scan the /sys/kernel/debug/tracing/events/syscalls directory and list all tracepoints that have a format file.`,
	Run: func(cmd *cobra.Command, args []string) {
		tracepoints, err := tracepoint.ListTracepoints()
		if err != nil {
			log.Fatalf("Error listing tracepoints: %v", err)
		}

		if len(tracepoints) == 0 {
			fmt.Println("No tracepoints found.")
			return
		}

		tui.ShowInteractiveSearch(tracepoints)
	},
}
