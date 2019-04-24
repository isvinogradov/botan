package telegrambot

import (
	en "car-quiz-v3/entities"
	"fmt"
	"strings"
)

type BotCallbacksContainer struct {
	// Handlers for ordinary events
	OnMessage            func(bot *Bot, msg *en.Message) error            // Message received
	OnEditedMessage      func(bot *Bot, msg *en.Message) error            // Message edit received
	OnCallbackQuery      func(bot *Bot, cbq *en.CallbackQuery) error      // Callback query received
	OnInlineQuery        func(bot *Bot, iq *en.InlineQuery) error         // Inline query received
	OnChannelPost        func(bot *Bot, msg *en.Message) error            // New channel post received
	OnEditedChannelPost  func(bot *Bot, msg *en.Message) error            // New channel post edit received
	OnChosenInlineResult func(bot *Bot, cir *en.ChosenInlineResult) error // New result for inline query received
	OnPoll               func(bot *Bot, poll *en.Poll) error              // Poll vote received

	// Handlers for storing offset in an external source like database or file
	OnGetOffset    func() int            // Get offset from external source
	OnSetNewOffset func(newUpdateId int) // Dump offset to external source

	// Error handling
	OnError func(err error) // Pass control and error instance to this callback if an error occurred during ordinary callback execution. Panic with error if you want the bot to stop immediately.
}

// if no callbacks for getting and storing offset were provided, then generate empty functions
func (cbCont *BotCallbacksContainer) checkAndInit() {
	if cbCont.OnGetOffset == nil {
		fmt.Println("GetOffset callback missing")
		cbCont.OnGetOffset = func() int { return 0 }
	}
	if cbCont.OnSetNewOffset == nil {
		fmt.Println("SetOffset callback missing")
		cbCont.OnSetNewOffset = func(newUpdateId int) {}
	}
	if cbCont.OnError == nil {
		fmt.Println("OnError callback missing")
		cbCont.OnError = func(err error) {} // don't stop on errors
	}
}

// returns a string which represents types of updates bot can receive
// in form of: "[\"message\", \"callback_query\", ...]"
func (cbCont *BotCallbacksContainer) generateAllowedUpdates() string {
	var availableCallbacks []string

	if cbCont.OnMessage != nil {
		availableCallbacks = append(availableCallbacks, "message")
	}
	if cbCont.OnEditedMessage != nil {
		availableCallbacks = append(availableCallbacks, "edited_message")
	}
	if cbCont.OnCallbackQuery != nil {
		availableCallbacks = append(availableCallbacks, "callback_query")
	}
	if cbCont.OnInlineQuery != nil {
		availableCallbacks = append(availableCallbacks, "inline_query")
	}
	if cbCont.OnChannelPost != nil {
		availableCallbacks = append(availableCallbacks, "channel_post")
	}
	if cbCont.OnEditedChannelPost != nil {
		availableCallbacks = append(availableCallbacks, "edited_channel_post")
	}
	if cbCont.OnChosenInlineResult != nil {
		availableCallbacks = append(availableCallbacks, "chosen_inline_result")
	}
	if cbCont.OnPoll != nil {
		availableCallbacks = append(availableCallbacks, "poll")
	}

	for i := range availableCallbacks {
		availableCallbacks[i] = fmt.Sprintf("\"%s\"", availableCallbacks[i])
	}

	types := strings.Join(availableCallbacks, ", ")
	return fmt.Sprintf("[%s]", types)
}
