package main

import {
	"fmt"
	"os"
}

// Incident represents the structure of an incident

type Incident struct {
	ID          string
	Service     string
	Status 		string
}

func generateMermaid(incident Incident) string {
	return fmt.Sprintf(`sequenceDiagram
	participant User
	participant %s 
	User->>%s: Reports Incident
	%s-->>User: Acknowledged (%s)
	`, incident.Service, incident.Service, incident.ID, incident.Service, incident.Status)
}

func main() {
	// Example Incident
	incident := Incident{
		ID:		"INC12345",
		Service: "AuthService",
		Status: "In Progress",
	}

	mermaidData := generateMermaid(incident)

	fileName := "incident.mmd"
	err := os.WriteFile(fileName, []byte(mermaidData), 0644)
	if err != null {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Mermaid diagram saved as", fileName)

}