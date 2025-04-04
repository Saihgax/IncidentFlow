// "generate" commmand for the CLI
package cmd

import (
	"fmt"
	"incidentflow/internal/diagram"
	"incidentflow/internal/models"
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the 'generate' command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Mermaid.js diagram from an incident JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		if filePath == "" {
			fmt.Println("Error: Please specify an input JSON file using --file")
			os.Exit(1)
		}

		incident, err := diagram.ReadIncidentJSON(filePath)
		if err != nil {
			fmt.Println("Failed to read incident JSON:", err)
			os.Exit(1)
		}

		mermaid := diagram.GenerateMermaid(incident)
		fmt.Println(mermaid)
	},
}

func init() {
	generateCmd.Flags().StringP("file", "f", "", "Path to incident JSON file")
	rootCmd.AddCommand(generateCmd)
}
