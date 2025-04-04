package diagram

import (
	"encoding/json"
	"errors"
	"incidentflow/internal/models"
	"os"
)

// ReadIncidentJSON reads an incident from a JSON file
func ReadIncidentJSON(filePath string) (*models.Incident, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("failed to read file: " + err.Error())
	}

	var incident models.Incident
	err = json.Unmarshal(data, &incident)
	if err != nil {
		return nil, errors.New("failed to parse JSON: " + err.Error())
	}

	return &incident, nil
}
