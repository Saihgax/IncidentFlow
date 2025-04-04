// Renders Mermaid.js diagram as an image
package diagram

import (
	"fmt"
	"os"
	"os/exec"
)

// RenderMermaid takes Mermaid.js content and generates an image using mmdc
func RenderMermaid(mermaidContent, outputFile string) error {
	// Create a temporary Mermaid file
	tmpFile, err := os.CreateTemp("", "diagram-*.mmd")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write Mermaid.js content to the temp file
	_, err = tmpFile.WriteString(mermaidContent)
	if err != nil {
		return fmt.Errorf("failed to write Mermaid content: %v", err)
	}
	tmpFile.Close()

	// Run mmdc to generate an image
	cmd := exec.Command("mmdc", "-i", tmpFile.Name(), "-o", outputFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("mmdc error: %v, output: %s", err, string(output))
	}

	fmt.Println("Diagram generated:", outputFile)
	return nil
}
