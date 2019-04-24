package entities

type Sticker struct {
	FileId       string `json:"file_id"`                 // Unique identifier for this file
	Width        int    `json:"width"`                   // Sticker width
	Height       int    `json:"height"`                  // Sticker height
	Thumb        string `json:"thumb,omitempty"`         // Optional. Sticker thumbnail in the .webp or .jpg format
	Emoji        string `json:"emoji,omitempty"`         // Optional. Emoji associated with the sticker
	SetName      string `json:"set_name,omitempty"`      // Optional. Name of the sticker set to which the sticker belongs
	MaskPosition string `json:"mask_position,omitempty"` // Optional. For mask stickers, the position where the mask should be placed
	FileSize     int    `json:"file_size,omitempty"`     // Optional. File size
}
