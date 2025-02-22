// cmd/root.go
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Define the root command
var rootCmd = &cobra.Command{
	Use:   "eetf",
	Short: "Easy eBPF Tracepoint Finder",
	Long:  "EETF is a TUI tool to interactively and easily search (and manage) eBPF tracepoints.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(formatCmd)
}
