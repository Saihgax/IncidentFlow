package server

import (
	"fmt"
	"encoding/json"
	"incidentflow/internal/diagram"
	"incidentflow/internal/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// IncidentRequest represents the input JSON structure
type IncidentRequest struct {
	IncidentData map[string]interface{} `json:"incident_data"`
	OutputFormat string 				`json: "output_format"` // either mermaid or image
}

/*

Encodes req.IncidentData into JSON (json.Marshal)

Decodes it into models.Incident (json.Unmarshal)

Passes the structured incident object to diagram.GenerateMermaid

*/

// StartServer initializes the API server
func StartServer() {
	router := gin.Default()

	router.POST("/generate", func(c *gin.Context) {
		var req IncidentRequest
		if err := c.ShouldBindJSON(&req); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return 
		}

		// Convert map to Incident struct
		incidentJSON, err := json.Marshal(req.IncidentData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode incident data"})
			return
		}

		var incident models.Incident
		if err := json.Unmarshal(incidentJSON, &incident); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid incident data format"})
			return
		}

		// Generate Mermaid diagram
		mermaidCode := diagram.GenerateMermaid(&incident)

		if req.OutputFormat == "mermaid" {
			c.JSON(http.StatusOK, gin.H{"mermaid": mermaidCode})
			return
		} else if req.OutputFormat == "image" {
			outputFile := "incident.png"
			err := diagram.RenderMermaid(mermaidCode, outputFile)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate image"})
				return
			}

			// Serve the generated image
			c.File(outputFile)

			// Clean up the generated file after serving
			defer os.Remove(outputFile)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid output format"})
	})

	fmt.Println("Server running on port 8080...")
	router.Run(":8080")
}