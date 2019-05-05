package entities

// This object represents a sticker set.
type StickerSet struct {
	Name          string     `json:"name"`           // Sticker set name
	Title         string     `json:"title"`          // Sticker set title
	ContainsMasks bool       `json:"contains_masks"` // True, if the sticker set contains masks
	Stickers      []*Sticker `json:"stickers"`       // List of all set stickers
}
