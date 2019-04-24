package entities

import (
	"encoding/json"
	"errors"
)

// Generalization of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove or ForceReply
type ReplyMarkup struct {
	ExactType interface{} // put one of the appropriate types above here
}

// Custom Marshaler allows to substitute ReplyMarkup with exact type and validate it
func (rm *ReplyMarkup) MarshalJSON() ([]byte, error) {
	switch t := rm.ExactType.(type) {
	case
		InlineKeyboardMarkup,
		*InlineKeyboardMarkup,
		ReplyKeyboardMarkup,
		*ReplyKeyboardMarkup,
		ReplyKeyboardRemove,
		*ReplyKeyboardRemove,
		ForceReply,
		*ForceReply:
		return json.Marshal(t)
	}
	panic(errors.New("invalid ReplyMarkup value"))
}

// This object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

// This object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text                         string `json:"text"`                                       // Label text on the button
	Url                          string `json:"url,omitempty"`                              // Optional. HTTP or tg:// url to be opened when button is pressed
	CallbackData                 string `json:"callback_data,omitempty"`                    // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchInlineQuery            string `json:"switch_inline_query,omitempty"`              // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted. Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"` // Optional. If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted. This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
	Pay                          bool   `json:"pay,omitempty"`                              // Optional. Specify True, to send a Pay button. NOTE: This type of button must always be the first button in the first row.
}

// This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`                    // Array of button rows, each represented by an Array of KeyboardButton objects
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`   // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"` // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	Selective       bool               `json:"selective,omitempty"`         // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message. Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
}

// This object represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields are mutually exclusive.
type KeyboardButton struct {
	Text            string `json:"text"`                       // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestContact  bool   `json:"request_contact,omitempty"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	RequestLocation bool   `json:"request_location,omitempty"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
}

// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`     // =true. Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      bool `json:"selective,omitempty"` // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message. Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ForceReply bool `json:"force_reply"`         // =true. Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	Selective  bool `json:"selective,omitempty"` // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}
