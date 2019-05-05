package botan

import (
	en "github.com/isvinogradov/botan/entities"
)

// A simple method for testing your bot's auth token. Requires no parameters. Returns basic information
// about the bot in form of a User object.
func (bot *Bot) GetMe() (*en.User, error) {
	var target en.User
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId                interface{}     `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Text                  string          `json:"text"`                               // Text of the message to be sent
	ParseMode             string          `json:"parse_mode,omitempty"`               // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	ReplyMarkup           *en.ReplyMarkup `json:"reply_markup,omitempty"`             // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyToMessageId      int             `json:"reply_to_message_id,omitempty"`      // Optional. If the message is a reply, ID of the original message
	DisableNotification   bool            `json:"disable_notification,omitempty"`     // Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableWebPagePreview bool            `json:"disable_web_page_preview,omitempty"` // Optional. Disables link previews for links in this message
}

// todo connect target and url (mapping)
func (bot *Bot) SendMessage(msg *SendMessageRequest) (*en.Message, error) {
	var target en.Message
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId              interface{}     `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo               string          `json:"photo"`                          // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data
	Caption             string          `json:"caption,omitempty"`              // Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters
	ParseMode           string          `json:"parse_mode,omitempty"`           // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DisableNotification bool            `json:"disable_notification,omitempty"` // Optional. Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int             `json:"reply_to_message_id,omitempty"`  // Optional. If the message is a reply, ID of the original message
	ReplyMarkup         *en.ReplyMarkup `json:"reply_markup,omitempty"`         // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot *Bot) SendPhoto(sPhoto *SendPhotoRequest) (*en.Message, error) {
	var target en.Message
	if postErr := bot.requestGate.makePostRequest(
		bot.urls.sendPhoto,
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
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId          interface{}              `json:"chat_id,omitempty"`           // Optional  Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int                      `json:"message_id,omitempty"`        // Optional  Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId string                   `json:"inline_message_id,omitempty"` // Optional  Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *en.InlineKeyboardMarkup `json:"reply_markup,omitempty"`      // Optional  A JSON-serialized object for an inline keyboard.
}

func (bot *Bot) EditMessageReplyMarkup(editReplyMkup *EditMessageReplyMarkupRequest) (*en.Message, error) {
	var target en.Message
	if postErr := bot.requestGate.makePostRequest(
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
// todo check Results pointer
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
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Action string      `json:"action"`  // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
}

func (bot *Bot) SendChatAction(chatAction *SendChatActionRequest) (bool, error) {
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId              interface{}     `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername). A native poll can't be sent to a private chat.
	Question            string          `json:"question"`                       // Poll question, 1-255 characters
	Options             []string        `json:"options"`                        // List of answer options, 2-10 strings 1-100 characters each
	DisableNotification bool            `json:"disable_notification,omitempty"` // Optional 	Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int             `json:"reply_to_message_id,omitempty"`  // Optional 	If the message is a reply, ID of the original message
	ReplyMarkup         *en.ReplyMarkup `json:"reply_markup,omitempty"`         // Optional 	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot *Bot) SendPoll(poll *SendPollRequest) (*en.Message, error) {
	var target en.Message
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId      interface{}              `json:"chat_id"`                // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId   int                      `json:"message_id"`             // Identifier of the original message with the poll
	ReplyMarkup *en.InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional 	A JSON-serialized object for a new message inline keyboard.
}

func (bot *Bot) StopPoll(poll *StopPollRequest) (*en.Poll, error) {
	var target en.Poll
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId              interface{}     `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Sticker             string          `json:"sticker"`                        // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	DisableNotification bool            `json:"disable_notification,omitempty"` // Optional 	Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int             `json:"reply_to_message_id,omitempty"`  // Optional 	If the message is a reply, ID of the original message
	ReplyMarkup         *en.ReplyMarkup `json:"reply_markup,omitempty"`         // Optional 	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot *Bot) SendSticker(stickerReq *SendStickerRequest) (*en.Message, error) {
	var target en.Message
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (bot *Bot) GetChat(getChatReq *GetChatRequest) (*en.Chat, error) {
	var target en.Chat
	if postErr := bot.requestGate.makePostRequest(
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
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId              interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	FromChatId          interface{} `json:"from_chat_id"`                   // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	DisableNotification bool        `json:"disable_notification,omitempty"` // Optional 	Sends the message silently. Users will receive a notification with no sound.
	MessageId           int         `json:"message_id"`                     // Message identifier in the chat specified in from_chat_id
}

func (bot *Bot) ForwardMessage(fwdMsgReq *ForwardMessageRequest) (*en.Message, error) {
	var target en.Message
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string      `json:"title"`   // New chat title, 1-255 characters
}

func (bot *Bot) SetChatTitle(setChatTReq *SetChatTitleRequest) (bool, error) {
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId              interface{}     `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
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
	if postErr := bot.requestGate.makePostRequest(
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
	ChatId              interface{}     `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
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
	if postErr := bot.requestGate.makePostRequest(
		bot.urls.sendVoice,
		sendVoiceReq,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// Use this method to get basic info about a file and prepare it for downloading. For the moment, bots can download files
// of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the
// link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed
// that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
type GetFileRequest struct {
	FileId string `json:"file_id"` // File identifier to get info about
}

func (bot *Bot) GetFile(getFileReq *GetFileRequest) (*en.File, error) {
	var target en.File
	if postErr := bot.requestGate.makePostRequest(
		bot.urls.getFile,
		getFileReq,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

type SendLocationRequest struct {
	ChatId              interface{}     `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude            float32         `json:"latitude"`                       // Latitude of the location
	Longitude           float32         `json:"longitude"`                      // Longitude of the location
	LivePeriod          int             `json:"live_period,omitempty"`          // Optional 	Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	DisableNotification bool            `json:"disable_notification,omitempty"` // Optional 	Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int             `json:"reply_to_message_id,omitempty"`  // Optional 	If the message is a reply, ID of the original message
	ReplyMarkup         *en.ReplyMarkup `json:"reply_markup,omitempty"`         // Optional 	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// Use this method to send point on the map. On success, the sent Message is returned.
func (bot *Bot) SendLocation(sendLocReq *SendLocationRequest) (*en.Message, error) {
	var target en.Message
	if postErr := bot.requestGate.makePostRequest(
		bot.urls.sendLocation,
		sendLocReq,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

// todo document and thumb is an inputfile or string
type SendDocumentRequest struct {
	ChatId              interface{}     `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Document            string          `json:"document"`                       // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Thumb               string          `json:"thumb,omitempty"`                // Optional 	Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption             string          `json:"caption,omitempty"`              // Optional 	Document caption (may also be used when resending documents by file_id), 0-1024 characters
	ParseMode           string          `json:"parse_mode,omitempty"`           // Optional 	Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DisableNotification bool            `json:"disable_notification,omitempty"` // Optional 	Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int             `json:"reply_to_message_id,omitempty"`  // Optional 	If the message is a reply, ID of the original message
	ReplyMarkup         *en.ReplyMarkup `json:"reply_markup,omitempty"`         // Optional 	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files
// of any type of up to 50 MB in size, this limit may be changed in the future.
func (bot *Bot) SendDocument(sendDocReq *SendDocumentRequest) (*en.Message, error) {
	var target en.Message
	if postErr := bot.requestGate.makePostRequest(
		bot.urls.sendDocument,
		sendDocReq,
		&target,
	); postErr != nil {
		return nil, postErr
	}
	return &target, nil
}

type SendVideoRequest struct{}

// Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document).
// On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit
// may be changed in the future.
func (bot *Bot) SendVideo(svReq *SendVideoRequest) (*en.Message, error) {}

type SendVideoNoteRequest struct{}

// As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send
// video messages. On success, the sent Message is returned.
func (bot *Bot) SendVideoNote(svnReq *SendVideoNoteRequest) (*en.Message, error) {}

type SendMediaGroupRequest struct{}

// Use this method to send a group of photos or videos as an album. On success, an array of the sent Messages is returned.
func (bot *Bot) SendMediaGroup(smgReq *SendMediaGroupRequest) ([]*en.Message, error) {}

type EditMessageLiveLocationRequest struct{}

// Use this method to edit live location messages. A location can be edited until its live_period expires or editing
// is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message was sent by the bot,
// the edited Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageLiveLocation(emllReq *EditMessageLiveLocationRequest) (zzz, error) {}

type StopMessageLiveLocationRequest struct{}

// Use this method to stop updating a live location message before live_period expires. On success, if the message was
// sent by the bot, the sent Message is returned, otherwise True is returned.
func (bot *Bot) StopMessageLiveLocation(smllReq *StopMessageLiveLocationRequest) (zzz, error) {}

type SendVenueRequest struct{}

// Use this method to send information about a venue. On success, the sent Message is returned.
func (bot *Bot) SendVenue(svenReq *SendVenueRequest) (*en.Message, error) {}

type SendContactRequest struct{}

// Use this method to send phone contacts. On success, the sent Message is returned.
func (bot *Bot) SendContact(sconReq *SendContactRequest) (*en.Message, error) {}

type KickChatMemberRequest struct{}

// Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups and channels,
// the user will not be able to return to the group on their own using invite links, etc., unless unbanned first.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// Returns True on success.
func (bot *Bot) KickChatMember(kcmReq *KickChatMemberRequest) (bool, error) {}

type UnbanChatMemberRequest struct{}

// Use this method to unban a previously kicked user in a supergroup or channel. The user will not return to the group
// or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work.
// Returns True on success.
func (bot *Bot) UnbanChatMember(ucmReq *UnbanChatMemberRequest) (bool, error) {}

type RestrictChatMemberRequest struct{}

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this
// to work and must have the appropriate admin rights. Pass True for all boolean parameters to lift restrictions from
// a user. Returns True on success.
func (bot *Bot) RestrictChatMember(rcmReq *RestrictChatMemberRequest) (bool, error) {}

type PromoteChatMemberRequest struct{}

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in
// the chat for this to work and must have the appropriate admin rights. Pass False for all boolean parameters to
// demote a user. Returns True on success.
func (bot *Bot) PromoteChatMember(pcmReq *PromoteChatMemberRequest) (bool, error) {}

type ExportChatInviteLinkRequest struct{}

// Use this method to generate a new invite link for a chat; any previously generated link is revoked. The bot must
// be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the new invite
// link as String on success.
func (bot *Bot) ExportChatInviteLink(ecilReq *ExportChatInviteLinkRequest) (string, error) {}

type SetChatPhotoRequest struct{}

// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be
// an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (bot *Bot) SetChatPhoto(scpReq *SetChatPhotoRequest) (bool, error) {}

type DeleteChatPhotoRequest struct{}

// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (bot *Bot) DeleteChatPhoto(dcpReq *DeleteChatPhotoRequest) (bool, error) {}

type SetChatDescriptionRequest struct{}

// Use this method to change the description of a supergroup or a channel. The bot must be an administrator in the chat
// for this to work and must have the appropriate admin rights. Returns True on success.
func (bot *Bot) SetChatDescription(scdReq *SetChatDescriptionRequest) (bool, error) {}

type PinChatMessageRequest struct{}

// Use this method to pin a message in a group, a supergroup, or a channel. The bot must be an administrator in the chat
// for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin
// right in the channel. Returns True on success.
func (bot *Bot) PinChatMessage(picmReq *PinChatMessageRequest) (bool, error) {}

type UnpinChatMessageRequest struct{}

// Use this method to unpin a message in a group, a supergroup, or a channel. The bot must be an administrator in the
// chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’
// admin right in the channel. Returns True on success.
func (bot *Bot) UnpinChatMessage(upcmReq *UnpinChatMessageRequest) (bool, error) {}

type LeaveChatRequest struct{}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
func (bot *Bot) LeaveChat(lcmReq *LeaveChatRequest) (bool, error) {}

type GetChatAdministratorsRequest struct{}

// Use this method to get a list of administrators in a chat. On success, returns an Array of ChatMember objects that
// contains information about all chat administrators except other bots. If the chat is a group or a supergroup and
// no administrators were appointed, only the creator will be returned.
func (bot *Bot) GetChatAdministrators(gcaReq *GetChatAdministratorsRequest) ([]*en.ChatMember, error) {
}

type GetChatMembersCountRequest struct{}

// Use this method to get the number of members in a chat. Returns Int on success.
func (bot *Bot) GetChatMembersCount(gcmcReq *GetChatMembersCountRequest) (int, error) {}

type GetChatMemberRequest struct{}

// Use this method to get information about a member of a chat. Returns a ChatMember object on success.
func (bot *Bot) GetChatMember(gcmemReq *GetChatMemberRequest) (*en.ChatMember, error) {}

type SetChatStickerSetRequest struct{}

// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for
// this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in
// getChat requests to check if the bot can use this method. Returns True on success.
func (bot *Bot) SetChatStickerSet(scstReq *SetChatStickerSetRequest) (bool, error) {}

type DeleteChatStickerSetRequest struct{}

// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for
// this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in
// getChat requests to check if the bot can use this method. Returns True on success.
func (bot *Bot) DeleteChatStickerSet(dcstReq *DeleteChatStickerSetRequest) (bool, error) {}

type EditMessageTextRequest struct{}

// Use this method to edit text and game messages. On success, if edited message is sent by the bot, the edited
// Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageText(emtReq *EditMessageTextRequest) (zzz, error) {}

type EditMessageCaptionRequest struct{}

// Use this method to edit captions of messages. On success, if edited message is sent by the bot, the edited
// Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageCaption(emcReq *EditMessageCaptionRequest) (zzz, error) {}

type EditMessageMediaRequest struct{}

// Use this method to edit animation, audio, document, photo, or video messages. If a message is a part of a message
// album, then it can be edited only to a photo or a video. Otherwise, message type can be changed arbitrarily. When
// inline message is edited, new file can't be uploaded. Use previously uploaded file via its file_id or specify a URL.
// On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
func (bot *Bot) EditMessageMedia(emmReq *EditMessageMediaRequest) (zzz, error) {}

type DeleteMessageRequest struct{}

// Use this method to delete a message, including service messages, with the following limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
// Returns True on success.
func (bot *Bot) DeleteMessage(dmReq *DeleteMessageRequest) (bool, error) {}

//type GetStickerSetRequest struct{}
//
//// Use this method to get a sticker set. On success, a StickerSet object is returned.
//func (bot *Bot) GetStickerSet(gstsReq *GetStickerSetRequest) (*en.StickerSet, error) {}

//type UploadStickerFileRequest struct{}
//
//// Use this method to upload a .png file with a sticker for later use in createNewStickerSet and addStickerToSet
//// methods (can be used multiple times). Returns the uploaded File on success.
//func (bot *Bot) UploadStickerFile(ustfReq *UploadStickerFileRequest) (File, error) {}

type CreateNewStickerSetRequest struct{}

// Use this method to create new sticker set owned by a user. The bot will be able to edit the created sticker set.
// Returns True on success.
func (bot *Bot) CreateNewStickerSet(cnstsReq *CreateNewStickerSetRequest) (bool, error) {}

type AddStickerToSetRequest struct{}

// Use this method to add a new sticker to a set created by the bot. Returns True on success.
func (bot *Bot) AddStickerToSet(asttsReq *AddStickerToSetRequest) (bool, error) {}

type SetStickerPositionInSetRequest struct{}

// Use this method to move a sticker in a set created by the bot to a specific position . Returns True on success.
func (bot *Bot) SetStickerPositionInSet(sstpisReq *SetStickerPositionInSetRequest) (bool, error) {}

type DeleteStickerFromSetRequest struct{}

// Use this method to delete a sticker from a set created by the bot. Returns True on success.
func (bot *Bot) DeleteStickerFromSet(dstfsReq *DeleteStickerFromSetRequest) (bool, error) {}
