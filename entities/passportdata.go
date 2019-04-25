package entities

// Contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`        // Array with information about documents and other Telegram Passport elements that was shared with the bot
	Credentials *EncryptedCredentials      `json:"credentials"` // Encrypted credentials required to decrypt the data
}
