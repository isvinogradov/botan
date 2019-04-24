package entities

import (
	"encoding/json"
	"errors"
)

// This object represents one result of an inline query.
// Telegram clients currently support results of the following 20 types:
// - InlineQueryResultCachedAudio
// - InlineQueryResultCachedDocument
// - InlineQueryResultCachedGif
// - InlineQueryResultCachedMpeg4Gif
// - InlineQueryResultCachedPhoto
// - InlineQueryResultCachedSticker
// - InlineQueryResultCachedVideo
// - InlineQueryResultCachedVoice
// - InlineQueryResultArticle
// - InlineQueryResultAudio
// - InlineQueryResultContact
// - InlineQueryResultGame
// - InlineQueryResultDocument
// - InlineQueryResultGif
// - InlineQueryResultLocation
// - InlineQueryResultMpeg4Gif
// - InlineQueryResultPhoto
// - InlineQueryResultVenue
// - InlineQueryResultVideo
// - InlineQueryResultVoice
type InlineQueryResult struct {
	ExactType interface{} // put one of the appropriate types above here
}

// Custom Marshaler allows to substitute InlineQueryResult with exact type and validate it
func (iqr *InlineQueryResult) MarshalJSON() ([]byte, error) {
	switch t := iqr.ExactType.(type) {
	case
		InlineQueryResultCachedAudio,
		*InlineQueryResultCachedAudio,
		InlineQueryResultCachedDocument,
		*InlineQueryResultCachedDocument,
		InlineQueryResultCachedGif,
		*InlineQueryResultCachedGif,
		InlineQueryResultCachedMpeg4Gif,
		*InlineQueryResultCachedMpeg4Gif,
		InlineQueryResultCachedPhoto,
		*InlineQueryResultCachedPhoto,
		InlineQueryResultCachedSticker,
		*InlineQueryResultCachedSticker,
		InlineQueryResultCachedVideo,
		*InlineQueryResultCachedVideo,
		InlineQueryResultCachedVoice,
		*InlineQueryResultCachedVoice,
		InlineQueryResultArticle,
		*InlineQueryResultArticle,
		InlineQueryResultAudio,
		*InlineQueryResultAudio,
		InlineQueryResultContact,
		*InlineQueryResultContact,
		InlineQueryResultGame,
		*InlineQueryResultGame,
		InlineQueryResultDocument,
		*InlineQueryResultDocument,
		InlineQueryResultGif,
		*InlineQueryResultGif,
		InlineQueryResultLocation,
		*InlineQueryResultLocation,
		InlineQueryResultMpeg4Gif,
		*InlineQueryResultMpeg4Gif,
		InlineQueryResultPhoto,
		*InlineQueryResultPhoto,
		InlineQueryResultVenue,
		*InlineQueryResultVenue,
		InlineQueryResultVideo,
		*InlineQueryResultVideo,
		InlineQueryResultVoice,
		*InlineQueryResultVoice:
		return json.Marshal(t)
	}
	panic(errors.New("invalid InlineQueryResult value"))
}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`                   // Type of the result, must be article
	Id                  string                `json:"id"`                     // Unique identifier for this result, 1-64 Bytes
	Title               string                `json:"title"`                  // Title of the result
	InputMessageContent InputMessageContent   `json:"input_message_content"`  // Content of the message to be sent
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional. Inline keyboard attached to the message
	Url                 string                `json:"url,omitempty"`          // Optional. URL of the result
	HideUrl             bool                  `json:"hide_url,omitempty"`     // Optional. Pass True, if you don't want the URL to be shown in the message
	Description         string                `json:"description,omitempty"`  // Optional. Short description of the result
	ThumbUrl            string                `json:"thumb_url,omitempty"`    // Optional. Url of the thumbnail for the result
	ThumbWidth          int                   `json:"thumb_width,omitempty"`  // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height,omitempty"` // Optional. Thumbnail height
}

// Represents a link to a photo. By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`                            // Type of the result, must be photo
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	PhotoUrl            string                `json:"photo_url"`                       // A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
	ThumbUrl            string                `json:"thumb_url"`                       // URL of the thumbnail for the photo
	PhotoWidth          int                   `json:"photo_width,omitempty"`           // Optional. Width of the photo
	PhotoHeight         int                   `json:"photo_height,omitempty"`          // Optional. Height of the photo
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the photo to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the photo
}

// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user
// with optional caption. Alternatively, you can use input_message_content to send a message with
// the specified content instead of the animation.
type InlineQueryResultGif struct {
	Type                string                `json:"type"`                            // Type of the result, must be gif
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	GifUrl              string                `json:"gif_url"`                         // A valid URL for the GIF file. File size must not exceed 1MB
	GifWidth            int                   `json:"gif_width,omitempty"`             // Optional. Width of the GIF
	GifHeight           int                   `json:"gif_height,omitempty"`            // Optional. Height of the GIF
	GifDuration         int                   `json:"gif_duration,omitempty"`          // Optional. Duration of the GIF
	ThumbUrl            string                `json:"thumb_url"`                       // URL of the static thumbnail for the result (jpeg or gif)
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the GIF file to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the GIF animation
}

// Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file
// will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send
// a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	Type                string                `json:"type"`                            // Type of the result, must be gif
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	GifFileId           string                `json:"gif_file_id"`                     // A valid file identifier for the GIF file
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the GIF file to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the GIF animation
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can
// use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	Type                string                `json:"type"`                            // Type of the result, must be mpeg4_gif
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Mpeg4FileId         string                `json:"mpeg4_file_id"`                   // A valid file identifier for the MP4 file
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video animation
}

// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`                            // Type of the result, must be sticker
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	StickerFileId       string                `json:"sticker_file_id"`                 // A valid file identifier of the sticker
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the sticker
}

// Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with
// an optional caption. Alternatively, you can use input_message_content to send a message with the specified content
// instead of the file.
type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`                            // Type of the result, must be document
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                           // Title for the result
	DocumentFileId      string                `json:"document_file_id"`                // A valid file identifier for the file
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the document to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the file
}

// Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by
// the user with an optional caption. Alternatively, you can use input_message_content to send a message with
// the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	Type                string                `json:"type"`                            // Type of the result, must be video
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VideoFileId         string                `json:"video_file_id"`                   // A valid file identifier for the video file
	Title               string                `json:"title"`                           // Title for the result
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the video to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video
}

// Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent
// by the user. Alternatively, you can use input_message_content to send a message with the specified content instead
// of the voice message.
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`                            // Type of the result, must be voice
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VoiceFileId         string                `json:"voice_file_id"`                   // A valid file identifier for the voice message
	Title               string                `json:"title"`                           // Voice message title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the voice message
}

// Represents a link to an mp3 audio file stored on the Telegram servers. By default, this audio file will be sent
// by the user. Alternatively, you can use input_message_content to send a message with the specified content instead
// of the audio.
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`                            // Type of the result, must be audio
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	AudioFileId         string                `json:"audio_file_id"`                   // A valid file identifier for the audio file
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
}

// Represents a Game.
type InlineQueryResultGame struct {
	Type          string                `json:"type"`                   // Type of the result, must be game
	Id            string                `json:"id"`                     // Unique identifier for this result, 1-64 bytes
	GameShortName string                `json:"game_short_name"`        // Short name of the game
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional. Inline keyboard attached to the message
}

// Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user
// with an optional caption. Alternatively, you can use input_message_content to send a message with the specified
// content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	Type                string                `json:"type"`                            // Type of the result, must be photo
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	PhotoFileId         string                `json:"photo_file_id"`                   // A valid file identifier of the photo
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the photo to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the photo
}

// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content
// to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`                            // Type of the result, must be venue
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
	Latitude            float32               `json:"latitude"`                        // Latitude of the venue location in degrees
	Longitude           float32               `json:"longitude"`                       // Longitude of the venue location in degrees
	Title               string                `json:"title"`                           // Title of the venue
	Address             string                `json:"address"`                         // Address of the venue
	FoursquareId        string                `json:"foursquare_id,omitempty"`         // Optional. Foursquare identifier of the venue if known
	FoursquareType      string                `json:"foursquare_type,omitempty"`       // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the venue
	ThumbUrl            string                `json:"thumb_url,omitempty"`             // Optional. Url of the thumbnail for the result
	ThumbWidth          int                   `json:"thumb_width,omitempty"`           // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height,omitempty"`          // Optional. Thumbnail height
}

// Represents a link to an mp3 audio file. By default, this audio file will be sent by the user. Alternatively,
// you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`                            // Type of the result, must be audio
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	AudioUrl            string                `json:"audio_url"`                       // A valid URL for the audio file
	Title               string                `json:"title"`                           // Title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Performer           string                `json:"performer,omitempty"`             // Optional. Performer
	AudioDuration       int                   `json:"audio_duration,omitempty"`        // Optional. Audio duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the audio
}

// Represents a link to a voice recording in an .ogg container encoded with OPUS. By default, this voice recording
// will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified
// content instead of the the voice message.
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`                            // Type of the result, must be voice
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VoiceUrl            string                `json:"voice_url"`                       // A valid URL for the voice recording
	Title               string                `json:"title"`                           // Recording title
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	VoiceDuration       int                   `json:"voice_duration,omitempty"`        // Optional. Recording duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the voice recording
}

// Represents a link to a file. By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`                            // Type of the result, must be document
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                           // Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the document to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DocumentUrl         string                `json:"document_url"`                    // A valid URL for the file
	MimeType            string                `json:"mime_type"`                       // Mime type of the content of the file, either “application/pdf” or “application/zip”
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the file
	ThumbUrl            string                `json:"thumb_url,omitempty"`             // Optional. URL of the thumbnail (jpeg only) for the file
	ThumbWidth          int                   `json:"thumb_width,omitempty"`           // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height,omitempty"`          // Optional. Thumbnail height
}

// Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	Type                string                `json:"type"`                            // Type of the result, must be location
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
	Latitude            float32               `json:"latitude"`                        // Location latitude in degrees
	Longitude           float32               `json:"longitude"`                       // Location longitude in degrees
	Title               string                `json:"title"`                           // Location title
	LivePeriod          int                   `json:"live_period,omitempty"`           // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the location
	ThumbUrl            string                `json:"thumb_url,omitempty"`             // Optional. Url of the thumbnail for the result
	ThumbWidth          int                   `json:"thumb_width,omitempty"`           // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height,omitempty"`          // Optional. Thumbnail height
}

// Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively,
// you can use input_message_content to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	Type                string                `json:"type"`                            // Type of the result, must be contact
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 Bytes
	PhoneNumber         string                `json:"phone_number"`                    // Contact's phone number
	FirstName           string                `json:"first_name"`                      // Contact's first name
	LastName            string                `json:"last_name,omitempty"`             // Optional. Contact's last name
	Vcard               string                `json:"vcard,omitempty"`                 // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the contact
	ThumbUrl            string                `json:"thumb_url,omitempty"`             // Optional. Url of the thumbnail for the result
	ThumbWidth          int                   `json:"thumb_width,omitempty"`           // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height,omitempty"`          // Optional. Thumbnail height
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated
// MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content
// to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	Type                string                `json:"type"`                            // Type of the result, must be mpeg4_gif
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	Mpeg4Url            string                `json:"mpeg4_url"`                       // A valid URL for the MP4 file. File size must not exceed 1MB
	Mpeg4Width          int                   `json:"mpeg4_width,omitempty"`           // Optional. Video width
	Mpeg4Height         int                   `json:"mpeg4_height,omitempty"`          // Optional. Video height
	Mpeg4Duration       int                   `json:"mpeg4_duration,omitempty"`        // Optional. Video duration
	ThumbUrl            string                `json:"thumb_url"`                       // URL of the static thumbnail (jpeg or gif) for the result
	Title               string                `json:"title,omitempty"`                 // Optional. Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video animation
}

// Represents a link to a page containing an embedded video player or a video file. By default, this video file
// will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send
// a message with the specified content instead of the video. If an InlineQueryResultVideo message contains
// an embedded video (e.g., YouTube), you must replace its content using input_message_content.
type InlineQueryResultVideo struct {
	Type                string                `json:"type"`                            // Type of the result, must be video
	Id                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	VideoUrl            string                `json:"video_url"`                       // A valid URL for the embedded video player or video file
	MimeType            string                `json:"mime_type"`                       // Mime type of the content of video url, “text/html” or “video/mp4”
	ThumbUrl            string                `json:"thumb_url"`                       // URL of the thumbnail (jpeg only) for the video
	Title               string                `json:"title"`                           // Title for the result
	Caption             string                `json:"caption,omitempty"`               // Optional. Caption of the video to be sent, 0-1024 characters
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	VideoWidth          int                   `json:"video_width,omitempty"`           // Optional. Video width
	VideoHeight         int                   `json:"video_height,omitempty"`          // Optional. Video height
	VideoDuration       int                   `json:"video_duration,omitempty"`        // Optional. Video duration in seconds
	Description         string                `json:"description,omitempty"`           // Optional. Short description of the result
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"` // Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
}
