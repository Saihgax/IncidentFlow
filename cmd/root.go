// Main CLI entry point
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "incident-diagram",
	Short: "Incident Diagram Generator CLI",
	Long:  `A command-line tool to generate Mermaid.js diagrams from incident JSON data.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'incident-diagram generate -f incident.json' to generate a diagram")
	},
}

// Execute runs the CLI tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
