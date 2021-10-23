package equipmentSeller

import (
	"errors"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	service "github.com/ozonmp/omp-bot/internal/service/business/equipmentSeller"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "")

	args := inputMessage.CommandArguments()
	id, err := uuid.Parse(strings.TrimSpace(args))
	if err != nil {
		msg.Text = "seller ID must be valid UUID"
		_, err = c.bot.Send(msg)
		return err
	}

	es, err := c.equipmentSellerService.Describe(id)
	if err != nil {
		if errors.Is(err, service.ErrSellerNotFound) {
			msg.Text = fmt.Sprintf("seller with ID %q not found", id.String())
		} else {
			msg.Text = "internal error"
		}
		_, err = c.bot.Send(msg)
		return err
	}

	msg.Text = es.String()
	_, err = c.bot.Send(msg)
	return nil
}
