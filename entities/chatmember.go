package entities

// This object contains information about one member of a chat.
type ChatMember struct {
	User                  *User  `json:"user"`                                // Information about the user
	Status                string `json:"status"`                              // The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
	UntilDate             int    `json:"until_date,omitempty"`                // Optional. Restricted and kicked only. Date when restrictions will be lifted for this user, unix time
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`             // Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`           // Optional. Administrators only. True, if the administrator can change the chat title, photo and other settings
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`         // Optional. Administrators only. True, if the administrator can post in the channel, channels only
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`         // Optional. Administrators only. True, if the administrator can edit messages of other users and can pin messages, channels only
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`       // Optional. Administrators only. True, if the administrator can delete messages of other users
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`          // Optional. Administrators only. True, if the administrator can invite new users to the chat
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`      // Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`          // Optional. Administrators only. True, if the administrator can pin messages, groups and supergroups only
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`       // Optional. Administrators only. True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	IsMember              bool   `json:"is_member,omitempty"`                 // Optional. Restricted only. True, if the user is a member of the chat at the moment of the request
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`         // Optional. Restricted only. True, if the user can send text messages, contacts, locations and venues
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`   // Optional. Restricted only. True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`   // Optional. Restricted only. True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"` // Optional. Restricted only. True, if user may add web page previews to his messages, implies can_send_media_messages
}
