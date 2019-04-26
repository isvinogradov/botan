package botan

import (
	"errors"
	"fmt"
	"github.com/isvinogradov/botan/entities"
	"math/rand"
	"time"
)

type Bot struct {
	config    *config
	urls      *BotUrlContainer
	callbacks *BotCallbacksContainer
}

func NewBot(conf *config, callbacks *BotCallbacksContainer) (*Bot, error) {
	if conf == nil {
		return nil, errors.New("nil config pointer")
	}
	if callbacks == nil {
		return nil, errors.New("nil callback container pointer")
	}

	bot := Bot{config: conf, callbacks: callbacks}
	bot.callbacks.checkAndInit()
	generateUrlsForBot(&bot)
	rand.Seed(time.Now().Unix()) // rand seed for GetRandomQuestion
	return &bot, nil
}

// Get bulk of updates for bot
func newUpdatesResponse(url string) (*entities.UpdatesResponse, error) {
	var uResp entities.UpdatesResponse
	if parseError := getJson(url, &uResp); parseError != nil {
		return nil, parseError
	}
	if !uResp.OK {
		return nil, errors.New("failed to validate updates response")
	}
	return &uResp, nil
}

// GetUpdates long polling loop
func (bot *Bot) GetUpdates() {
	fmt.Println("started getUpdates loop")

	// get updateID (offset) from last run; if this is a first ever call to Redis, offset == 0
	offset := bot.callbacks.OnGetOffset()

	var url string
	for {
		if offset > 0 {
			// get first update after current offset
			url = fmt.Sprintf("%s&offset=%d", bot.urls.getUpdates, offset+1)
		} else {
			// URL with no offset filter
			url = bot.urls.getUpdates
		}

		// fetch response
		response, updRespErr := newUpdatesResponse(url)
		if updRespErr != nil {
			fmt.Println("got error in getUpdates; scheduling GetUpdates timeout...")
			fmt.Println(updRespErr)
			time.Sleep(time.Duration(bot.config.getUpdatesFailCooldownSeconds) * time.Second)
		} else {
			for _, update := range response.Updates {
				fmt.Printf("Processing update %d\n", update.UpdateId)
				var cbErr error
				// At most one of (message, edited_message, channel_post, edited_channel_post, inline_query,
				// chosen_inline_result, callback_query, shipping_query, pre_checkout_query) can be present
				// in any given update.

				// According to BotCallbacksContainer configuration, in Update, we will never receive an entity
				// with to respective callback set. So, for instance, if we received a CallbackQuery in Update,
				// we can be sure that OnCallbackQuery callback != nil in BotCallbacksContainer.

				if update.Message.Id != 0 {
					cbErr = bot.callbacks.OnMessage(bot, &update.Message)
				} else if update.CallbackQuery.Id != "" { // todo test pointer objects
					cbErr = bot.callbacks.OnCallbackQuery(bot, &update.CallbackQuery)
				} else if update.InlineQuery.Id != "" {
					cbErr = bot.callbacks.OnInlineQuery(bot, &update.InlineQuery)
				} else if update.EditedMessage.Id != 0 {
					cbErr = bot.callbacks.OnEditedMessage(bot, &update.EditedMessage)
				} else if update.ChannelPost.Id != 0 {
					cbErr = bot.callbacks.OnChannelPost(bot, &update.ChannelPost)
				} else if update.EditedChannelPost.Id != 0 {
					cbErr = bot.callbacks.OnEditedChannelPost(bot, &update.EditedChannelPost)
				} else if update.ChosenInlineResult.ResultId != "" {
					cbErr = bot.callbacks.OnChosenInlineResult(bot, &update.ChosenInlineResult)
				} else if update.Poll.Id != "" {
					cbErr = bot.callbacks.OnPoll(bot, &update.Poll)
				}

				// handle callback error
				if cbErr != nil {
					bot.callbacks.OnError(cbErr)
				}

				// everything OK
				// set new offset, even if no callback is triggered
				offset = update.UpdateId
				// store new offset to external storage if provided
				bot.callbacks.OnSetNewOffset(offset)
			}
		}
	}
}