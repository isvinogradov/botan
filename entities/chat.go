package entities

import (
	"encoding/json"
	"errors"
)

// This object represents a chat.
type Chat struct {
	Id                          int       `json:"id"`                                       // Unique identifier for this chat. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	Type                        string    `json:"type"`                                     // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Title                       string    `json:"title,omitempty"`                          // Optional. Title, for supergroups, channels and group chats
	Username                    string    `json:"username,omitempty"`                       // Optional. Username, for private chats, supergroups and channels if available
	FirstName                   string    `json:"first_name,omitempty"`                     // Optional. First name of the other party in a private chat
	LastName                    string    `json:"last_name,omitempty"`                      // Optional. Last name of the other party in a private chat
	AllMembersAreAdministrators bool      `json:"all_members_are_administrators,omitempty"` // Optional. True if a group has ‘All Members Are Admins’ enabled.
	Photo                       ChatPhoto `json:"photo,omitempty"`                          // Optional. Chat photo. Returned only in getChat.
	Description                 string    `json:"description,omitempty"`                    // Optional. Description, for supergroups and channel chats. Returned only in getChat.
	InviteLink                  string    `json:"invite_link,omitempty"`                    // Optional. Chat invite link, for supergroups and channel chats. Each administrator in a chat generates their own invite links, so the bot must first generate the link using exportChatInviteLink. Returned only in getChat.
	PinnedMessage               Message   `json:"pinned_message,omitempty"`                 // Optional. Pinned message, for groups, supergroups and channels. Returned only in getChat.
	StickerSetName              string    `json:"sticker_set_name,omitempty"`               // Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetStickerSet            bool      `json:"can_set_sticker_set,omitempty"`            // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
}

// This object represents a chat photo.
type ChatPhoto struct {
	SmallFileId string `json:"small_file_id"` // Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download.
	BigFileId   string `json:"big_file_id"`   // Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download.
}

// this struct represents chat_id because it can be int or string
type ChatId struct {
	ExactType interface{}
}

func (cid *ChatId) MarshalJSON() ([]byte, error) {
	switch t := cid.ExactType.(type) {
	case int, string:
		return json.Marshal(t)
	}
	panic(errors.New("invalid ChatId value (must be int or string)"))
}
