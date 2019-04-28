package botan

import (
	"fmt"
)

// preformatted URLs for all API methods
type BotUrlContainer struct {
	getUpdates           string
	sendMessage          string
	sendPhoto            string
	answerCallback       string
	updateMessageMarkup  string
	getMe                string
	sendChatAction       string
	forwardMessage       string
	sendAudio            string
	answerInlineQuery    string
	sendPoll             string
	stopPoll             string
	sendSticker          string
	getChat              string
	getUserProfilePhotos string
	setChatTitle         string
	sendAnimation        string
	sendVoice            string
	getFile              string
}

// pre-generated urls for all supported bot methods
func generateUrlsForBot(bot *Bot) {
	urlPrefix := fmt.Sprintf(
		"%s/bot%s/",
		TelegramApiHost,
		bot.config.token,
	)

	getUpdatesFullUrl := fmt.Sprintf(
		"%s%s?timeout=%d&allowed_updates=%s",
		urlPrefix,
		MethodGetUpdates,
		bot.config.longPollTimeoutSeconds,
		bot.callbacks.generateAllowedUpdates(),
	)

	bot.urls = &BotUrlContainer{
		getUpdates:           getUpdatesFullUrl,
		sendMessage:          fmt.Sprintf("%s%s", urlPrefix, MethodSendMessage),
		sendPhoto:            fmt.Sprintf("%s%s", urlPrefix, MethodSendPhoto),
		answerCallback:       fmt.Sprintf("%s%s", urlPrefix, MethodAnswerCallbackQuery),
		updateMessageMarkup:  fmt.Sprintf("%s%s", urlPrefix, MethodEditReplyMarkup),
		getMe:                fmt.Sprintf("%s%s", urlPrefix, MethodGetMe),
		sendChatAction:       fmt.Sprintf("%s%s", urlPrefix, MethodSendChatAction),
		forwardMessage:       fmt.Sprintf("%s%s", urlPrefix, MethodForwardMessage),
		sendAudio:            fmt.Sprintf("%s%s", urlPrefix, MethodSendAudio),
		answerInlineQuery:    fmt.Sprintf("%s%s", urlPrefix, MethodAnswerInlineQuery),
		sendPoll:             fmt.Sprintf("%s%s", urlPrefix, MethodSendPoll),
		stopPoll:             fmt.Sprintf("%s%s", urlPrefix, MethodStopPoll),
		sendSticker:          fmt.Sprintf("%s%s", urlPrefix, MethodSendSticker),
		getChat:              fmt.Sprintf("%s%s", urlPrefix, MethodGetChat),
		getUserProfilePhotos: fmt.Sprintf("%s%s", urlPrefix, MethodGetUserProfilePhotos),
		setChatTitle:         fmt.Sprintf("%s%s", urlPrefix, MethodSetChatTitle),
		sendAnimation:        fmt.Sprintf("%s%s", urlPrefix, MethodSendAnimation),
		sendVoice:            fmt.Sprintf("%s%s", urlPrefix, MethodSendVoice),
		getFile:              fmt.Sprintf("%s%s", urlPrefix, MethodGetFile),
	}
}
