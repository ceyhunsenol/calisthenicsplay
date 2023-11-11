package model

// RequirementContentRequest Request
type RequirementContentRequest struct {
	ContentID            string
	RequirementContentID string
}

// HelperContentRequest Request
type HelperContentRequest struct {
	ContentID       string
	HelperContentID string
}

type ContentModel struct {
	ID string
}
