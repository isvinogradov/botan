package entities

import (
	"encoding/json"
	"errors"
)

// This object represents the content of a message to be sent as a result of an inline query.
// Telegram clients currently support the following 4 types:
// - InputTextMessageContent
// - InputLocationMessageContent
// - InputVenueMessageContent
// - InputContactMessageContent
type InputMessageContent struct {
	ExactType interface{} // put one of the appropriate types above here
}

// Custom Marshaler allows to substitute InputMessageContent with exact type and validate it
func (imc *InputMessageContent) MarshalJSON() ([]byte, error) {
	switch t := imc.ExactType.(type) {
	case
		InputTextMessageContent,
		*InputTextMessageContent,
		InputLocationMessageContent,
		*InputLocationMessageContent,
		InputVenueMessageContent,
		*InputVenueMessageContent,
		InputContactMessageContent,
		*InputContactMessageContent:
		return json.Marshal(t)
	}
	panic(errors.New("invalid InputMessageContent value"))
}

// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	MessageText           string `json:"message_text"`                       // Text of the message to be sent, 1-4096 characters
	ParseMode             string `json:"parse_mode,omitempty"`               // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"` // Optional. Disables link previews for links in the sent message
}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude   float32 `json:"latitude"`              // Latitude of the location in degrees
	Longitude  float32 `json:"longitude"`             // Longitude of the location in degrees
	LivePeriod int     `json:"live_period,omitempty"` // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	Latitude       float32 `json:"latitude"`                  // Latitude of the venue in degrees
	Longitude      float32 `json:"longitude"`                 // Longitude of the venue in degrees
	Title          string  `json:"title"`                     // Name of the venue
	Address        string  `json:"address"`                   // Address of the venue
	FoursquareId   string  `json:"foursquare_id,omitempty"`   // Optional. Foursquare identifier of the venue, if known
	FoursquareType string  `json:"foursquare_type,omitempty"` // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
}

// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`        // Contact's phone number
	FirstName   string `json:"first_name"`          // Contact's first name
	LastName    string `json:"last_name,omitempty"` // Optional. Contact's last name
	Vcard       string `json:"vcard,omitempty"`     // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
}
