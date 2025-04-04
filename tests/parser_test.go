package tests

import (
	"../internal/diagram"
	"testing"
)

func TestReadIncidentJSON(t *testing.T) {
	incident, err := diagram.ReadIncidentJSON("../incident.json")
	if err != nil {
		t.Fatalf("Failed to read incident JSON: %v", err)
	}

	if incident.Title == "" || len(incident.Events) == 0 {
		t.Fatalf("Parsed incident data is incomplete")
	}
}
