// Renders Mermaid.js diagram as an image
package diagram

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

// renderMaid takes Mermaid.js syntax and renders it as an image
func renderMaid(mermaidCode string, outputFile string) error {
	
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Mermaid HTML template
	htmlContent := fmt.Sprintf(`
	<html>
	<head>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/mermaid/10.1.0/mermaid.min.js"></script>
	</head>
	<body>
		<div class="mermaid">%s</div>
		<script>mermaid.init();</script>
	</body>
	</html>`, mermaidCode)

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate("data:text/html,"+htmlContent),
		chromedp.Sleep(2*time.Second), // Allow time for Mermaid.js to render
		chromedp.CaptureScreenshot(&buf),
	)
	if err != nil {
		return err
	}

	// Save screenshot
	return os.WriteFile(outputFile, buf, 0644)
}

func main() {
	incidentDiagram := `sequenceDiagram
	participant A
	participant B
	A->>B: Hello
	B-->>A: Hi`

	outputFile := "incident.png"
	err := renderMaid(incidentDiagram, outputFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Incident diagram saved as", outputFile)
}
