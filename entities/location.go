package entities

// This object represents a point on the map.
type Location struct {
	Longitude float32 `json:"longitude"` // Longitude as defined by sender
	Latitude  float32 `json:"latitude"`  // Latitude as defined by sender
}
