package entities

// This object contains information about a poll.
type Poll struct {
	Id       string       `json:"id"`        // Unique poll identifier
	Question string       `json:"question"`  // Poll question, 1-255 characters
	Options  []PollOption `json:"options"`   // List of poll options
	IsClosed bool         `json:"is_closed"` // True, if the poll is closed
}

// This object contains information about one answer option in a poll.
type PollOption struct {
	Text       string `json:"text"`        // Option text, 1-100 characters
	VoterCount int    `json:"voter_count"` // Number of users that voted for this option
}
