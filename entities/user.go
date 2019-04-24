package entities

// This object represents a Telegram user or bot.
type User struct {
	Id           int    `json:"id"`                      // Unique identifier for this user or bot
	IsBot        bool   `json:"is_bot"`                  // True, if this user is a bot
	FirstName    string `json:"first_name"`              // User‘s or bot’s first name
	LastName     string `json:"last_name,omitempty"`     // Optional. User‘s or bot’s last name
	Username     string `json:"username,omitempty"`      // Optional. User‘s or bot’s username
	LanguageCode string `json:"language_code,omitempty"` // Optional. IETF language tag of the user's language
}

// This object represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"` // Total number of profile pictures the target user has
	Photos     [][]PhotoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}
