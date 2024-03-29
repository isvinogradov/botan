package botan

// GENERAL
const (
	TelegramApiHost = "https://api.telegram.org"
)

// TELEGRAM BOT API METHODS
const (
	MethodGetUpdates              = "getUpdates"
	MethodSendPhoto               = "sendPhoto"
	MethodSendMessage             = "sendMessage"
	MethodEditReplyMarkup         = "editMessageReplyMarkup"
	MethodAnswerCallbackQuery     = "answerCallbackQuery"
	MethodGetMe                   = "getMe"
	MethodSendChatAction          = "sendChatAction"
	MethodSendAudio               = "sendAudio"
	MethodAnswerInlineQuery       = "answerInlineQuery"
	MethodSendPoll                = "sendPoll"
	MethodStopPoll                = "stopPoll"
	MethodSendSticker             = "sendSticker"
	MethodGetChat                 = "getChat"
	MethodGetUserProfilePhotos    = "getUserProfilePhotos"
	MethodForwardMessage          = "forwardMessage"
	MethodSetChatTitle            = "setChatTitle"
	MethodSendAnimation           = "sendAnimation"
	MethodSendVoice               = "sendVoice"
	MethodGetFile                 = "getFile"
	MethodSendLocation            = "sendLocation"
	MethodSendDocument            = "sendDocument"
	MethodSendVideo               = "sendVideo"
	MethodSendVideoNote           = "sendVideoNote"
	MethodSendMediaGroup          = "sendMediaGroup"
	MethodEditMessageLiveLocation = "editMessageLiveLocation"
	MethodStopMessageLiveLocation = "stopMessageLiveLocation"
	MethodSendVenue               = "sendVenue"
	MethodSendContact             = "sendContact"
	MethodKickChatMember          = "kickChatMember"
	MethodUnbanChatMember         = "unbanChatMember"
	MethodRestrictChatMember      = "restrictChatMember"
	MethodPromoteChatMember       = "promoteChatMember"
	MethodExportChatInviteLink    = "exportChatInviteLink"
	MethodSetChatPhoto            = "setChatPhoto"
	MethodDeleteChatPhoto         = "deleteChatPhoto"
	MethodSetChatDescription      = "setChatDescription"
	MethodPinChatMessage          = "pinChatMessage"
	MethodUnpinChatMessage        = "unpinChatMessage"
	MethodLeaveChat               = "leaveChat"
	MethodGetChatAdministrators   = "getChatAdministrators"
	MethodGetChatMembersCount     = "getChatMembersCount"
	MethodGetChatMember           = "getChatMember"
	MethodSetChatStickerSet       = "setChatStickerSet"
	MethodDeleteChatStickerSet    = "deleteChatStickerSet"
	MethodEditMessageText         = "editMessageText"
	MethodEditMessageCaption      = "editMessageCaption"
	MethodEditMessageMedia        = "editMessageMedia"
	MethodDeleteMessage           = "deleteMessage"
	MethodGetStickerSet           = "getStickerSet"
	MethodUploadStickerFile       = "uploadStickerFile"
	MethodCreateNewStickerSet     = "createNewStickerSet"
	MethodAddStickerToSet         = "addStickerToSet"
	MethodSetStickerPositionInSet = "setStickerPositionInSet"
	MethodDeleteStickerFromSet    = "deleteStickerFromSet"
)

// TELEGRAM BOT API FORMATTING OPTIONS
type MessageFormat string

const (
	FormatHtml     MessageFormat = "HTML"
	FormatMarkdown               = "Markdown"
)

// CONSTANTS FOR SEND CHAT ACTION METHOD
type ChatAction string

const (
	Typing          ChatAction = "typing"
	UploadPhoto                = "upload_photo"
	UploadDocument             = "upload_document"
	FindLocation               = "find_location"
	UploadVideo                = "upload_video"
	UploadAudio                = "upload_audio"
	UploadVideoNote            = "upload_video_note"
)
