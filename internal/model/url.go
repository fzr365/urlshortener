package model

type CreateURLRequest struct {
	OriginalURL string `json:"original_url" validate:"required"`
	CustomCode  string `json:"custom_code,omitempty"`
	Duration    *int    `json:"duration,omitempty"`
}

