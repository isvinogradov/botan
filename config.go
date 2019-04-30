package botan

// telegram bot options and properties container
type Config struct {
	Token                         string // telegram bot Token obtained from BotFather
	PostJsonTimeoutSeconds        int    // timeout for all bot methods (sendMessage etc.)
	LongPollTimeoutSeconds        int    // long polling timeout for getUpdates method
	GetUpdatesFailCooldownSeconds int    // sleep duration scheduled when getUpdates request fails
	Socks5ConnectionString        string // connections string if SOCKS5 proxy is used
}

//func NewConfig(
//	token string,
//	postJsonTimeoutSec,
//	longPollTimeoutSec,
//	getUpdatesFailCooldownSec int,
//	socks5ConnectionString string,
//) (*Config, error) {
//	if token == "" {
//		return nil, errors.New("bot token not specified")
//	}
//	if postJsonTimeoutSec < 1 {
//		postJsonTimeoutSec = 5
//	}
//	if longPollTimeoutSec < 1 {
//		longPollTimeoutSec = 300
//	}
//	if getUpdatesFailCooldownSec < 1 {
//		getUpdatesFailCooldownSec = 10
//	}
//
//	return &Config{
//		Token:                         token,
//		PostJsonTimeoutSeconds:        postJsonTimeoutSec,
//		LongPollTimeoutSeconds:        longPollTimeoutSec,
//		GetUpdatesFailCooldownSeconds: getUpdatesFailCooldownSec,
//		Socks5ConnectionString:        socks5ConnectionString,
//	}, nil
//}
