package botan

import (
	"fmt"
)

// preformatted URLs for all API methods
type BotUrlContainer struct {
	getUpdates              string
	sendMessage             string
	sendPhoto               string
	answerCallback          string
	updateMessageMarkup     string
	getMe                   string
	sendChatAction          string
	forwardMessage          string
	sendAudio               string
	answerInlineQuery       string
	sendPoll                string
	stopPoll                string
	sendSticker             string
	getChat                 string
	getUserProfilePhotos    string
	setChatTitle            string
	sendAnimation           string
	sendVoice               string
	getFile                 string
	sendLocation            string
	sendDocument            string
	sendVideo               string
	sendVideoNote           string
	sendMediaGroup          string
	editMessageLiveLocation string
	stopMessageLiveLocation string
	sendVenue               string
	sendContact             string
	kickChatMember          string
	unbanChatMember         string
	restrictChatMember      string
	promoteChatMember       string
	exportChatInviteLink    string
	setChatPhoto            string
	deleteChatPhoto         string
	setChatDescription      string
	pinChatMessage          string
	unpinChatMessage        string
	leaveChat               string
	getChatAdministrators   string
	getChatMembersCount     string
	getChatMember           string
	setChatStickerSet       string
	deleteChatStickerSet    string
	editMessageText         string
	editMessageCaption      string
	editMessageMedia        string
	deleteMessage           string
	getStickerSet           string
	uploadStickerFile       string
	createNewStickerSet     string
	addStickerToSet         string
	setStickerPositionInSet string
	deleteStickerFromSet    string
}

// pre-generated urls for all supported bot methods
func generateUrlsForBot(bot *Bot) {
	urlPrefix := fmt.Sprintf(
		"%s/bot%s/",
		TelegramApiHost,
		bot.config.Token,
	)

	getUpdatesFullUrl := fmt.Sprintf(
		"%s%s?timeout=%d&allowed_updates=%s",
		urlPrefix,
		MethodGetUpdates,
		bot.config.LongPollTimeoutSeconds,
		bot.callbacks.generateAllowedUpdates(),
	)

	bot.urls = &BotUrlContainer{
		getUpdates:              getUpdatesFullUrl,
		sendMessage:             fmt.Sprintf("%s%s", urlPrefix, MethodSendMessage),
		sendPhoto:               fmt.Sprintf("%s%s", urlPrefix, MethodSendPhoto),
		answerCallback:          fmt.Sprintf("%s%s", urlPrefix, MethodAnswerCallbackQuery),
		updateMessageMarkup:     fmt.Sprintf("%s%s", urlPrefix, MethodEditReplyMarkup),
		getMe:                   fmt.Sprintf("%s%s", urlPrefix, MethodGetMe),
		sendChatAction:          fmt.Sprintf("%s%s", urlPrefix, MethodSendChatAction),
		forwardMessage:          fmt.Sprintf("%s%s", urlPrefix, MethodForwardMessage),
		sendAudio:               fmt.Sprintf("%s%s", urlPrefix, MethodSendAudio),
		answerInlineQuery:       fmt.Sprintf("%s%s", urlPrefix, MethodAnswerInlineQuery),
		sendPoll:                fmt.Sprintf("%s%s", urlPrefix, MethodSendPoll),
		stopPoll:                fmt.Sprintf("%s%s", urlPrefix, MethodStopPoll),
		sendSticker:             fmt.Sprintf("%s%s", urlPrefix, MethodSendSticker),
		getChat:                 fmt.Sprintf("%s%s", urlPrefix, MethodGetChat),
		getUserProfilePhotos:    fmt.Sprintf("%s%s", urlPrefix, MethodGetUserProfilePhotos),
		setChatTitle:            fmt.Sprintf("%s%s", urlPrefix, MethodSetChatTitle),
		sendAnimation:           fmt.Sprintf("%s%s", urlPrefix, MethodSendAnimation),
		sendVoice:               fmt.Sprintf("%s%s", urlPrefix, MethodSendVoice),
		getFile:                 fmt.Sprintf("%s%s", urlPrefix, MethodGetFile),
		sendLocation:            fmt.Sprintf("%s%s", urlPrefix, MethodSendLocation),
		sendDocument:            fmt.Sprintf("%s%s", urlPrefix, MethodSendDocument),
		sendVideo:               fmt.Sprintf("%s%s", urlPrefix, MethodSendVideo),
		sendVideoNote:           fmt.Sprintf("%s%s", urlPrefix, MethodSendVideoNote),
		sendMediaGroup:          fmt.Sprintf("%s%s", urlPrefix, MethodSendMediaGroup),
		editMessageLiveLocation: fmt.Sprintf("%s%s", urlPrefix, MethodEditMessageLiveLocation),
		stopMessageLiveLocation: fmt.Sprintf("%s%s", urlPrefix, MethodStopMessageLiveLocation),
		sendVenue:               fmt.Sprintf("%s%s", urlPrefix, MethodSendVenue),
		sendContact:             fmt.Sprintf("%s%s", urlPrefix, MethodSendContact),
		kickChatMember:          fmt.Sprintf("%s%s", urlPrefix, MethodKickChatMember),
		unbanChatMember:         fmt.Sprintf("%s%s", urlPrefix, MethodUnbanChatMember),
		restrictChatMember:      fmt.Sprintf("%s%s", urlPrefix, MethodRestrictChatMember),
		promoteChatMember:       fmt.Sprintf("%s%s", urlPrefix, MethodPromoteChatMember),
		exportChatInviteLink:    fmt.Sprintf("%s%s", urlPrefix, MethodExportChatInviteLink),
		setChatPhoto:            fmt.Sprintf("%s%s", urlPrefix, MethodSetChatPhoto),
		deleteChatPhoto:         fmt.Sprintf("%s%s", urlPrefix, MethodDeleteChatPhoto),
		setChatDescription:      fmt.Sprintf("%s%s", urlPrefix, MethodSetChatDescription),
		pinChatMessage:          fmt.Sprintf("%s%s", urlPrefix, MethodPinChatMessage),
		unpinChatMessage:        fmt.Sprintf("%s%s", urlPrefix, MethodUnpinChatMessage),
		leaveChat:               fmt.Sprintf("%s%s", urlPrefix, MethodLeaveChat),
		getChatAdministrators:   fmt.Sprintf("%s%s", urlPrefix, MethodGetChatAdministrators),
		getChatMembersCount:     fmt.Sprintf("%s%s", urlPrefix, MethodGetChatMembersCount),
		getChatMember:           fmt.Sprintf("%s%s", urlPrefix, MethodGetChatMember),
		setChatStickerSet:       fmt.Sprintf("%s%s", urlPrefix, MethodSetChatStickerSet),
		deleteChatStickerSet:    fmt.Sprintf("%s%s", urlPrefix, MethodDeleteChatStickerSet),
		editMessageText:         fmt.Sprintf("%s%s", urlPrefix, MethodEditMessageText),
		editMessageCaption:      fmt.Sprintf("%s%s", urlPrefix, MethodEditMessageCaption),
		editMessageMedia:        fmt.Sprintf("%s%s", urlPrefix, MethodEditMessageMedia),
		deleteMessage:           fmt.Sprintf("%s%s", urlPrefix, MethodDeleteMessage),
		getStickerSet:           fmt.Sprintf("%s%s", urlPrefix, MethodGetStickerSet),
		uploadStickerFile:       fmt.Sprintf("%s%s", urlPrefix, MethodUploadStickerFile),
		createNewStickerSet:     fmt.Sprintf("%s%s", urlPrefix, MethodCreateNewStickerSet),
		addStickerToSet:         fmt.Sprintf("%s%s", urlPrefix, MethodAddStickerToSet),
		setStickerPositionInSet: fmt.Sprintf("%s%s", urlPrefix, MethodSetStickerPositionInSet),
		deleteStickerFromSet:    fmt.Sprintf("%s%s", urlPrefix, MethodDeleteStickerFromSet),
	}
}
