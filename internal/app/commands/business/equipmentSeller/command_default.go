package equipmentSeller

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("You wrote: %s", inputMessage.Text))
	_, err := c.bot.Send(msg)
	return err
}
