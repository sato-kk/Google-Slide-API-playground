package presentations

import (
	"fmt"
	"log"

	"google.golang.org/api/slides/v1"
)

// CreatePresentation creates a new Google Slides presentation.
func CreatePresentation(slidesService *slides.Service, title string) (string, error) {
	p := &slides.Presentation{
		Title: title,
	}
	presentation, err := slidesService.Presentations.Create(p).Fields(
		"presentationId",
	).Do()
	if err != nil {
		log.Printf("Unable to create presentation. %v", err)
		return "", err
	}
	fmt.Printf("Created presentation with ID: %s\n", presentation.PresentationId)
	return presentation.PresentationId, nil
}
