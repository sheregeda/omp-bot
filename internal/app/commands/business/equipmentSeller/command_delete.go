package equipmentSeller

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
)

func (c *Commander) Delete(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "")

	args := inputMessage.CommandArguments()
	id, err := uuid.Parse(strings.TrimSpace(args))
	if err != nil {
		msg.Text = "seller ID must be valid UUID"
		_, err = c.bot.Send(msg)
		return err
	}

	ok, err := c.equipmentSellerService.Remove(id)
	if err != nil {
		return err
	}
	if ok {
		msg.Text = "seller deleted"
	} else {
		msg.Text = fmt.Sprintf("seller with ID %q not found", id.String())
	}

	_, err = c.bot.Send(msg)
	return nil
}
