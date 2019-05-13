package entities

import (
	"strconv"
	"time"
)

// New incoming message of any kind — text, photo, sticker, etc.
//type Message1 struct {
//	Id                int             `json:"message_id"`                        // Unique message identifier inside this chat
//	Text              string          `json:"text,omitempty"`                    // Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
//	Sender            *User           `json:"from,omitempty"`                    // Optional. Sender, empty for messages sent to channels
//	ForwardFrom       *User           `json:"forward_from,omitempty"`            // Optional. For forwarded messages, sender of the original message
//	ForwardFromMsgId  int             `json:"forward_from_message_id,omitempty"` // Optional. For messages forwarded from channels, identifier of the original message in the channel
//	ReplyToMsg        *Message        `json:"reply_to_message,omitempty"`        // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
//	PinnedMsg         *Message        `json:"pinned_message,omitempty"`          // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
//	Dt                *JsonUnixTime   `json:"date"`                              // Date the message was sent in Unix time
//	Poll              *Poll           `json:"poll,omitempty"`                    // Optional. Message is a native poll, information about the poll
//	Animation         *Animation      `json:"entities,omitempty"`                // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
//	Entities          []MessageEntity `json:"entities,omitempty"`                // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
//	CaptionEntities   []MessageEntity `json:"caption_entities,omitempty"`        // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
//	Caption           string          `json:"caption,omitempty"`                 // Optional. Caption for the audio, document, photo, video or voice, 0-200 characters
//	Photo             []PhotoSize     `json:"photo,omitempty"`                   // Optional. Message is a photo, available sizes of the photo
//	Document          *Document       `json:"document,omitempty"`                // Optional. Message is a general file, information about the file
//	Sticker           *Sticker        `json:"sticker,omitempty"`                 // Optional. Message is a sticker, information about the sticker
//	Location          *Location       `json:"location,omitempty"`                // Optional. Message is a shared location, information about the location
//	Venue             *Venue          `json:"venue,omitempty"`                   // Optional. Message is a venue, information about the venue
//	EditDt            *JsonUnixTime   `json:"edit_date,omitempty"`               // Optional. Date the message was last edited in Unix time
//	ConnectedWebsite  string          `json:"connected_website,omitempty"`       // Optional. The domain name of the website on which the user has logged in
//	Game              *Game           `json:"game,omitempty"`                    // Optional. Message is a game
//	NewChatTitle      string          `json:"new_chat_title,omitempty"`          // Optional. A chat title was changed to this value
//	NewChatPhoto      []PhotoSize     `json:"new_chat_photo,omitempty"`          // Optional. A chat photo was change to this value
//	PassportData      *PassportData   `json:"passport_data,omitempty"`           // Optional. Telegram Passport data
//	LeftChatMember    *User           `json:"left_chat_member,omitempty"`        // Optional. A member was removed from the group, information about them (this member may be the bot itself)
//	NewChatMembers    []User          `json:"new_chat_members,omitempty"`        // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
//	MediaGroupId      string          `json:"media_group_id,omitempty"`          // Optional. The unique identifier of a media message group this message belongs to
//	AuthorSignature   string          `json:"author_signature,omitempty"`        // Optional. Signature of the post author for messages in channels
//	MigrateToChatId   int64           `json:"migrate_to_chat_id,omitempty"`      // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
//	MigrateFromChatId int64           `json:"migrate_from_chat_id,omitempty"`    // Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
//	Voice             *Voice          `json:"voice,omitempty"`                   // Optional. Message is a voice message, information about the file
//	VideoNote         *VideoNote      `json:"video_note,omitempty"`              // Optional. Message is a video note, information about the video message
//}

// This object represents a message.
type Message struct {
	MessageId             int                `json:"message_id"`                        // Unique message identifier inside this chat
	From                  *User              `json:"from,omitempty"`                    // Optional. Sender, empty for messages sent to channels
	Date                  *JsonUnixTime      `json:"date"`                              // Date the message was sent in Unix time
	Chat                  *Chat              `json:"chat"`                              // Conversation the message belongs to
	ForwardFrom           *User              `json:"forward_from,omitempty"`            // Optional. For forwarded messages, sender of the original message
	ForwardFromChat       *Chat              `json:"forward_from_chat,omitempty"`       // Optional. For messages forwarded from channels, information about the original channel
	ForwardFromMessageId  int                `json:"forward_from_message_id,omitempty"` // Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardSignature      string             `json:"forward_signature,omitempty"`       // Optional. For messages forwarded from channels, signature of the post author if present
	ForwardSenderName     string             `json:"forward_sender_name,omitempty"`     // Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
	ForwardDate           int                `json:"forward_date,omitempty"`            // Optional. For forwarded messages, date the original message was sent in Unix time
	ReplyToMessage        *Message           `json:"reply_to_message,omitempty"`        // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	EditDate              *JsonUnixTime      `json:"edit_date,omitempty"`               // Optional. Date the message was last edited in Unix time
	MediaGroupId          string             `json:"media_group_id,omitempty"`          // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature       string             `json:"author_signature,omitempty"`        // Optional. Signature of the post author for messages in channels
	Text                  string             `json:"text,omitempty"`                    // Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Entities              []MessageEntity    `json:"entities,omitempty"`                // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	CaptionEntities       []MessageEntity    `json:"caption_entities,omitempty"`        // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	Audio                 *Audio             `json:"audio,omitempty"`                   // Optional. Message is an audio file, information about the file
	Document              *Document          `json:"document,omitempty"`                // Optional. Message is a general file, information about the file
	Animation             *Animation         `json:"animation,omitempty"`               // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Game                  *Game              `json:"game,omitempty"`                    // Optional. Message is a game, information about the game. More about games »
	Photo                 []PhotoSize        `json:"photo,omitempty"`                   // Optional. Message is a photo, available sizes of the photo
	Sticker               *Sticker           `json:"sticker,omitempty"`                 // Optional. Message is a sticker, information about the sticker
	Video                 *Video             `json:"video,omitempty"`                   // Optional. Message is a video, information about the video
	Voice                 *Voice             `json:"voice,omitempty"`                   // Optional. Message is a voice message, information about the file
	VideoNote             *VideoNote         `json:"video_note,omitempty"`              // Optional. Message is a video note, information about the video message
	Caption               string             `json:"caption,omitempty"`                 // Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	Contact               *Contact           `json:"contact,omitempty"`                 // Optional. Message is a shared contact, information about the contact
	Location              *Location          `json:"location,omitempty"`                // Optional. Message is a shared location, information about the location
	Venue                 *Venue             `json:"venue,omitempty"`                   // Optional. Message is a venue, information about the venue
	Poll                  *Poll              `json:"poll,omitempty"`                    // Optional. Message is a native poll, information about the poll
	NewChatMembers        []User             `json:"new_chat_members,omitempty"`        // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember        *User              `json:"left_chat_member,omitempty"`        // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle          string             `json:"new_chat_title,omitempty"`          // Optional. A chat title was changed to this value
	NewChatPhoto          []PhotoSize        `json:"new_chat_photo,omitempty"`          // Optional. A chat photo was change to this value
	DeleteChatPhoto       bool               `json:"delete_chat_photo,omitempty"`       // Optional. Service message: the chat photo was deleted
	GroupChatCreated      bool               `json:"group_chat_created,omitempty"`      // Optional. Service message: the group has been created
	SupergroupChatCreated bool               `json:"supergroup_chat_created,omitempty"` // Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated    bool               `json:"channel_chat_created,omitempty"`    // Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MigrateToChatId       int                `json:"migrate_to_chat_id,omitempty"`      // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit int or double-precision float type are safe for storing this identifier.
	MigrateFromChatId     int                `json:"migrate_from_chat_id,omitempty"`    // Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit int or double-precision float type are safe for storing this identifier.
	PinnedMessage         *Message           `json:"pinned_message,omitempty"`          // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	Invoice               *Invoice           `json:"invoice,omitempty"`                 // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment,omitempty"`      // Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	ConnectedWebsite      string             `json:"connected_website,omitempty"`       // Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	PassportData          *PassportData      `json:"passport_data,omitempty"`           // Optional. Telegram Passport data
}

// tools for unmarshaling unix time from int
type JsonUnixTime time.Time

func (t *JsonUnixTime) UnmarshalJSON(b []byte) (err error) {
	ts, convErr := strconv.Atoi(string(b))
	if convErr != nil {
		return convErr
	}
	*t = JsonUnixTime(time.Unix(int64(ts), 0))
	return
}
