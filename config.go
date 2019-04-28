package botan

import (
	"errors"
)

// telegram bot options and properties container
type config struct {
	token                         string // telegram bot token obtained from BotFather
	postJsonTimeoutSeconds        int    // timeout for all bot methods (sendMessage etc.)
	longPollTimeoutSeconds        int    // long polling timeout for getUpdates method
	getUpdatesFailCooldownSeconds int    // sleep duration scheduled when getUpdates request fails
	socks5ConnectionString        string // connections string if SOCKS5 proxy is used
}

func NewConfig(
	token string,
	postJsonTimeoutSec,
	longPollTimeoutSec,
	getUpdatesFailCooldownSec int,
	socks5ConnectionString string,
) (*config, error) {
	if token == "" {
		return nil, errors.New("bot token not specified")
	}
	if postJsonTimeoutSec < 1 {
		postJsonTimeoutSec = 5
	}
	if longPollTimeoutSec < 1 {
		longPollTimeoutSec = 300
	}
	if getUpdatesFailCooldownSec < 1 {
		getUpdatesFailCooldownSec = 10
	}

	return &config{
		token:                         token,
		postJsonTimeoutSeconds:        postJsonTimeoutSec,
		longPollTimeoutSeconds:        longPollTimeoutSec,
		getUpdatesFailCooldownSeconds: getUpdatesFailCooldownSec,
		socks5ConnectionString:        socks5ConnectionString,
	}, nil
}
