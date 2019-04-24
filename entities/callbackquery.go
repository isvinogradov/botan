package entities

// This object represents an incoming callback query from a callback button in an inline keyboard. If the button that
// originated the query was attached to a message sent by the bot, the field message will be present. If the button
// was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present.
// Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	Id           string  `json:"id"`                          // Unique identifier for this query
	Sender       User    `json:"from"`                        // Sender
	ChatInstance string  `json:"chat_instance"`               // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent.
	Message      Message `json:"message,omitempty"`           // Optional. Message with the callback button that originated the query
	CallbackData string  `json:"data,omitempty"`              // Optional. Data associated with the callback button. Bad client can send arbitrary data in this field.
	InlineMsgId  string  `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
}
