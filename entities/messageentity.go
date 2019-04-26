package entities

// This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type          string `json:"type"`           // Type of the entity. Can be mention (@username), hashtag, cashtag, bot_command, url, email, phone_number, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
	Offset        int    `json:"offset"`         // Offset in UTF-16 code units to the start of the entity
	Length        int    `json:"length"`         // Length of the entity in UTF-16 code units
	Url           string `json:"url,omitempty"`  // Optional. For “text_link” only, url that will be opened after user taps on the text
	MentionedUser *User  `json:"user,omitempty"` // Optional. For “text_mention” only, the mentioned user
}


type MessageEntityType string

const (
	MessageEntityTypeMention     MessageEntityType = "mention"
	MessageEntityTypeHashtag                       = "hashtag"
	MessageEntityTypeCashtag                       = "cashtag"
	MessageEntityTypeBotCommand                    = "bot_command"
	MessageEntityTypeUrl                           = "url"
	MessageEntityTypeEmail                         = "email"
	MessageEntityTypePhoneNumber                   = "phone_number"
	MessageEntityTypeBold                          = "bold"
	MessageEntityTypeItalic                        = "italic"
	MessageEntityTypeCode                          = "code"
	MessageEntityTypePre                           = "pre"
	MessageEntityTypeTextLink                      = "text_link"
	MessageEntityTypeTextMention                   = "text_mention"
)
