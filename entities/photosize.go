package entities

// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileId   string `json:"file_id"`             // Unique identifier for this file
	Width    int    `json:"width"`               // Photo width
	Height   int    `json:"height"`              // Photo height
	FileSize int    `json:"file_size,omitempty"` // Optional. File size
}
