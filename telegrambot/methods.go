package telegrambot

import (
	en "car-quiz-v3/entities"
)

// A simple method for testing your bot's auth token. Requires no parameters. Returns basic information
// about the bot in form of a User object.
func (bot *Bot) GetMe() (*en.User, error) {
	var target en.User
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.getMe,
		nil,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to send text messages. On success, the sent Message is returned.
// todo check parsemode
type SendMessageRequest struct {
	ChatId                int             `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Text                  string          `json:"text"`                               // Text of the message to be sent
	ParseMode             string          `json:"parse_mode,omitempty"`               // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	ReplyMarkup           *en.ReplyMarkup `json:"reply_markup,omitempty"`             // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyToMessageId      int             `json:"reply_to_message_id,omitempty"`      // Optional. If the message is a reply, ID of the original message
	DisableNotification   bool            `json:"disable_notification,omitempty"`     // Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableWebPagePreview bool            `json:"disable_web_page_preview,omitempty"` // Optional. Disables link previews for links in this message
}

func (bot *Bot) SendMessage(msg *SendMessageRequest) (*en.Message, error) {
	var target en.Message
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.sendMessage,
		msg,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to send photos. On success, the sent Message is returned.
// todo: Photo can be InputFile
type SendPhotoRequest struct {
	ChatId              int             `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo               string          `json:"photo"`                          // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data
	Caption             string          `json:"caption,omitempty"`              // Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters
	ParseMode           string          `json:"parse_mode,omitempty"`           // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DisableNotification bool            `json:"disable_notification,omitempty"` // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int             `json:"reply_to_message_id,omitempty"`  // Optional. If the message is a reply, ID of the original message
	ReplyMarkup         *en.ReplyMarkup `json:"reply_markup,omitempty"`         // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot *Bot) SendPhoto(sPhoto *SendPhotoRequest) (*en.Message, error) {
	var target en.Message
	if postErr := makePostRequest( // todo move bot to container as well
		bot.config.postJsonTimeoutSeconds, // TODO: move to closure?
		bot.urls.sendPhoto,                // todo map methods to struct types?
		sPhoto,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed
// to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
type AnswerCallbackQueryRequest struct {
	CallbackQueryId string `json:"callback_query_id"`    // Unique identifier for the query to be answered
	Text            string `json:"text,omitempty"`       // Optional  Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert       bool   `json:"show_alert,omitempty"` // Optional  If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	Url             string `json:"url,omitempty"`        // Optional  URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game – note that this will only work if the query comes from a callback_game button. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime       int    `json:"cache_time,omitempty"` // Optional  The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}

func (bot *Bot) AnswerCallbackQuery(answerCbQ *AnswerCallbackQueryRequest) (bool, error) {
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.answerCallback,
		answerCbQ,
		nil,
	); postErr != nil {
		return false, postErr
	}
	return true, nil
}

// Use this method to edit only the reply markup of messages sent by the bot or via the bot (for inline bots).
// On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageReplyMarkupRequest struct {
	ChatId          int                      `json:"chat_id,omitempty"`           // Optional  Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int                      `json:"message_id,omitempty"`        // Optional  Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId string                   `json:"inline_message_id,omitempty"` // Optional  Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *en.InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional  A JSON-serialized object for an inline keyboard.
}

func (bot *Bot) EditMessageReplyMarkup(editReplyMkup *EditMessageReplyMarkupRequest) (*en.Message, error) {
	var target en.Message
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.updateMessageMarkup,
		editReplyMkup,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to send answers to an inline query. On success, True is returned. No more than 50 results
// per query are allowed.
type AnswerInlineQueryRequest struct {
	InlineQueryId     string                  `json:"inline_query_id"`               // Unique identifier for the answered query
	Results           *[]en.InlineQueryResult `json:"results"`                       // A JSON-serialized array of results for the inline query
	CacheTime         int                     `json:"cache_time,omitempty"`          // Optional  The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	IsPersonal        bool                    `json:"is_personal,omitempty"`         // Optional  Pass True, if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
	NextOffset        string                  `json:"next_offset,omitempty"`         // Optional  Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes.
	SwitchPmText      string                  `json:"switch_pm_text,omitempty"`      // Optional  If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
	SwitchPmParameter string                  `json:"switch_pm_parameter,omitempty"` // Optional  Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.
}

func (bot *Bot) AnswerInlineQuery(answer *AnswerInlineQueryRequest) (bool, error) {
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.answerInlineQuery,
		answer,
		nil,
	); postErr != nil {
		return false, postErr
	}
	return true, nil
}

// Use this method when you need to tell the user that something is happening on the bot's side. We only recommend
// using this method when a response from the bot will take a noticeable amount of time to arrive. The status is set
// for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
// Returns True on success.
// Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a
// text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with
// action = upload_photo. The user will see a “sending photo” status for the bot.
// todo check action
type SendChatActionRequest struct {
	ChatId int    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Action string `json:"action"`  // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
}

func (bot *Bot) SendChatAction(chatAction *SendChatActionRequest) (bool, error) {
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.sendChatAction,
		chatAction,
		nil,
	); postErr != nil {
		return false, postErr
	}
	return true, nil
}

// Use this method to send a native poll. A native poll can't be sent to a private chat.
// On success, the sent Message is returned.
type SendPollRequest struct {
	ChatId              int             `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername). A native poll can't be sent to a private chat.
	Question            string          `json:"question"`                       // Poll question, 1-255 characters
	Options             []string        `json:"options"`                        // List of answer options, 2-10 strings 1-100 characters each
	DisableNotification bool            `json:"disable_notification,omitempty"` // Optional 	Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int             `json:"reply_to_message_id,omitempty"`  // Optional 	If the message is a reply, ID of the original message
	ReplyMarkup         *en.ReplyMarkup `json:"reply_markup,omitempty"`         // Optional 	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot *Bot) SendPoll(poll *SendPollRequest) (*en.Message, error) {
	var target en.Message
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.sendPoll,
		poll,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to stop a poll which was sent by the bot.
// On success, the stopped Poll with the final results is returned.
type StopPollRequest struct {
	ChatId      int                      `json:"chat_id"`                // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId   int                      `json:"message_id"`             // Identifier of the original message with the poll
	ReplyMarkup *en.InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional 	A JSON-serialized object for a new message inline keyboard.
}

func (bot *Bot) StopPoll(poll *StopPollRequest) (*en.Poll, error) {
	var target en.Poll
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.stopPoll,
		poll,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to send .webp stickers. On success, the sent Message is returned.
// todo sticker=InputFile
type SendStickerRequest struct {
	ChatId              int             `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Sticker             string          `json:"sticker"`                        // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	DisableNotification bool            `json:"disable_notification,omitempty"` // Optional 	Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int             `json:"reply_to_message_id,omitempty"`  // Optional 	If the message is a reply, ID of the original message
	ReplyMarkup         *en.ReplyMarkup `json:"reply_markup,omitempty"`         // Optional 	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot *Bot) SendSticker(stickerReq *SendStickerRequest) (*en.Message, error) {
	var target en.Message
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.sendSticker,
		stickerReq,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to get up to date information about the chat (current name of the user for one-on-one
// conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
type GetChatRequest struct {
	ChatId int `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (bot *Bot) GetChat(getChatReq *GetChatRequest) (*en.Chat, error) {
	var target en.Chat
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.getChat,
		getChatReq,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
type GetUserProfilePhotosRequest struct {
	UserId int `json:"user_id"`          // Unique identifier of the target user
	Offset int `json:"offset,omitempty"` // Optional 	Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit  int `json:"limit,omitempty"`  // Optional 	Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
}

func (bot *Bot) GetUserProfilePhotos(getPhotReq *GetUserProfilePhotosRequest) (*en.UserProfilePhotos, error) {
	var target en.UserProfilePhotos
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.getUserProfilePhotos,
		getPhotReq,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to forward messages of any kind. On success, the sent Message is returned.
type ForwardMessageRequest struct {
	ChatId              int  `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	FromChatId          int  `json:"from_chat_id"`                   // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	DisableNotification bool `json:"disable_notification,omitempty"` // Optional 	Sends the message silently. Users will receive a notification with no sound.
	MessageId           int  `json:"message_id"`                     // Message identifier in the chat specified in from_chat_id
}

func (bot *Bot) ForwardMessage(fwdMsgReq *ForwardMessageRequest) (*en.Message, error) {
	var target en.Message
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.forwardMessage,
		fwdMsgReq,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be
// an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
// Note: In regular groups (non-supergroups), this method will only work if the ‘All Members Are Admins’ setting
// is off in the target group.
type SetChatTitleRequest struct {
	ChatId int    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string `json:"title"`   // New chat title, 1-255 characters
}

func (bot *Bot) SetChatTitle(setChatTReq *SetChatTitleRequest) (bool, error) {
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.setChatTitle,
		setChatTReq,
		nil,
	); postErr != nil {
		return false, postErr
	}
	return true, nil
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message
// is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
// todo: thumb/animation can be InputFile
type SendAnimationRequest struct {
	ChatId              int             `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Animation           string          `json:"animation"`                      // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More info on Sending Files »
	Duration            int             `json:"duration,omitempty"`             // Optional 	Duration of sent animation in seconds
	Width               int             `json:"width,omitempty"`                // Optional 	Animation width
	Height              int             `json:"height,omitempty"`               // Optional 	Animation height
	Thumb               string          `json:"thumb,omitempty"`                // Optional 	Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption             string          `json:"caption,omitempty"`              // Optional 	Animation caption (may also be used when resending animation by file_id), 0-1024 characters
	ParseMode           string          `json:"parse_mode,omitempty"`           // Optional 	Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DisableNotification bool            `json:"disable_notification,omitempty"` // Optional 	Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int             `json:"reply_to_message_id,omitempty"`  // Optional 	If the message is a reply, ID of the original message
	ReplyMarkup         *en.ReplyMarkup `json:"reply_markup,omitempty"`         // Optional 	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot *Bot) SendAnimation(sendAnReq *SendAnimationRequest) (*en.Message, error) {
	var target en.Message
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.sendAnimation,
		sendAnReq,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as
// Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up
// to 50 MB in size, this limit may be changed in the future.
type SendVoiceRequest struct {
	ChatId              int             `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Voice               string          `json:"voice"`                          // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Caption             string          `json:"caption,omitempty"`              // Optional 	Voice message caption, 0-1024 characters
	ParseMode           string          `json:"parse_mode,omitempty"`           // Optional 	Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Duration            int             `json:"duration,omitempty"`             // Optional 	Duration of the voice message in seconds
	DisableNotification bool            `json:"disable_notification,omitempty"` // Optional 	Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int             `json:"reply_to_message_id,omitempty"`  // Optional 	If the message is a reply, ID of the original message
	ReplyMarkup         *en.ReplyMarkup `json:"reply_markup,omitempty"`         // Optional 	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot *Bot) SendVoice(sendVoiceReq *SendVoiceRequest) (*en.Message, error) {
	var target en.Message
	if postErr := makePostRequest(
		bot.config.postJsonTimeoutSeconds,
		bot.urls.sendVoice,
		sendVoiceReq,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}
