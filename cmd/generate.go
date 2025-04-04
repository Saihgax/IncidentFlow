package cmd

import (
	"fmt"
	"incidentflow/internal/diagram"
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the 'generate' command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Mermaid.js diagram from an incident JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		outputFile, _ := cmd.Flags().GetString("output")

		if filePath == "" {
			fmt.Println("Error: Please specify an input JSON file using --file")
			os.Exit(1)
		}

		// Read incident data
		incident, err := diagram.ReadIncidentJSON(filePath)
		if err != nil {
			fmt.Println("Failed to read incident JSON:", err)
			os.Exit(1)
		}

		// Generate Mermaid.js
		mermaid := diagram.GenerateMermaid(incident)

		if outputFile == "" {
			fmt.Println(mermaid) // Print Mermaid.js to console
		} else {
			// Render as image
			err = diagram.RenderMermaid(mermaid, outputFile)
			if err != nil {
				fmt.Println("Failed to render diagram:", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	generateCmd.Flags().StringP("file", "f", "", "Path to incident JSON file")
	generateCmd.Flags().StringP("output", "o", "", "Path to output image file (PNG/SVG)")
	rootCmd.AddCommand(generateCmd)
}
