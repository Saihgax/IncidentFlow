// Defines incident struct
package models

import "time"

// Event represents a single event in an incident
type Event struct {
	Timestamp time.Time `json:"timestamp"`
	Actor     string    `json:"actor"`
	Action    string    `json:"action"`
	Details   string    `json:"details"`
}

// Incident represents the overall incident structure
type Incident struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	OccurredAt  time.Time `json:"occurred_at"`
	Events      []Event `json:"events"`
}
