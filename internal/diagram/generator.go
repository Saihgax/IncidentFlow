// Converts JSON to Mermaid diagram
package diagram

import (
	"fmt"
	"incidentflow/internal/models"
	"strings"
)

// GenerateMermaid takes an Incident and returns Mermaid.js sequence diagram
func GenerateMermaid(incident *models.Incident) string {
	var sb strings.Builder

	// Start the Mermaid sequence diagram
	sb.WriteString("sequenceDiagram\n")
	sb.WriteString(fmt.Sprintf("    participant System as %s\n", incident.Title))

	// Track unique actors to define them
	actors := make(map[string]bool)

	// Process each event and build the diagram
	for _, event := range incident.Events {
		actors[event.Actor] = true
		sb.WriteString(fmt.Sprintf("    %s->>System: %s\n", event.Actor, event.Action))
	}

	// Define all participants
	sb.WriteString("\n    %% Define Participants\n")
	for actor := range actors {
		sb.WriteString(fmt.Sprintf("    participant %s\n", actor))
	}

	return sb.String()
}
