package presentations

import (
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
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

// AddSlide adds a new slide to a presentation.
func AddSlide(slidesService *slides.Service, presentationId string) (*slides.BatchUpdatePresentationResponse, error) {
	slideId := "MyNewSlide_001"
	requests := []*slides.Request{{
		CreateSlide: &slides.CreateSlideRequest{
			ObjectId:       slideId,
			InsertionIndex: 1,
			SlideLayoutReference: &slides.LayoutReference{
				PredefinedLayout: "TITLE_AND_TWO_COLUMNS",
			},
		},
	}}

	// If you wish to populate the slide with elements,
	// using the slide ID specified above.

	// Execute the request.
	body := &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	}
	response, err := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	if err != nil {
		log.Printf("Unable to create slide. %v", err)
		return nil, err
	}
	fmt.Printf("Created slide with ID: %s\n", response.Replies[0].CreateSlide.ObjectId)
	return response, nil
}

// FindPresentationByTitle searches for a presentation by title.
func FindPresentationByTitle(driveService *drive.Service, title string) (string, error) {
	q := fmt.Sprintf("name='%s' and mimeType='application/vnd.google-apps.presentation' and trashed=false", title)
	files, err := driveService.Files.List().Q(q).Fields("files(id)").Do()
	if err != nil {
		log.Printf("Unable to search for presentation: %v", err)
		return "", err
	}

	if len(files.Files) > 0 {
		return files.Files[0].Id, nil
	}

	return "", nil // Not found
}
