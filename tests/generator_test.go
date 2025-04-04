package tests

import (
	"incidentflow/internal/diagram"
	"incidentflow/internal/models"
	"testing"
	"time"
	"strings"
)

func TestGenerateMermaid(t *testing.T) {
	incident := models.Incident{
		Title:       "Test Incident",
		Description: "A sample incident",
		OccurredAt:  time.Now(),
		Events: []models.Event{
			{Timestamp: time.Now(), Actor: "User", Action: "Reported", Details: "System down"},
			{Timestamp: time.Now(), Actor: "Engineer", Action: "Investigated", Details: "Checked logs"},
			{Timestamp: time.Now(), Actor: "System", Action: "Restarted", Details: "Rebooted service"},
		},
	}

	mermaid := diagram.GenerateMermaid(&incident)
	if !strings.Contains(mermaid, "sequenceDiagram") {
		t.Fatalf("Generated Mermaid.js is incorrect")
	}
}
