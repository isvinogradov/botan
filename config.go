package botan

// telegram bot options and properties container
type Config struct {
	Token                         string // telegram bot Token obtained from BotFather
	PostJsonTimeoutSeconds        int    // timeout for all bot methods (sendMessage etc.)
	LongPollTimeoutSeconds        int    // long polling timeout for getUpdates method
	GetUpdatesFailCooldownSeconds int    // sleep duration scheduled when getUpdates request fails
	Socks5ConnectionString        string // connections string if SOCKS5 proxy is used
}
