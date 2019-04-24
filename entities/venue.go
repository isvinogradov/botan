package entities

// This object represents a venue.
type Venue struct {
	Location     Location `json:"location"`                // Venue location
	Title        string   `json:"title"`                   // Name of the venue
	Address      string   `json:"address"`                 // Address of the venue
	FoursquareID string   `json:"foursquare_id,omitempty"` // Optional. Foursquare identifier of the venue
}
