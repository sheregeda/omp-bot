package equipmentSeller

import (
	"encoding/json"
	"errors"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	service "github.com/ozonmp/omp-bot/internal/service/business/equipmentSeller"
)

func (c *Commander) Edit(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "")

	args := inputMessage.CommandArguments()

	var data service.EquipmentSeller
	err := json.Unmarshal([]byte(args), &data)
	if err != nil {
		msg.Text = `invalid input data, it must be format: {"Name": "foobar"}`
		_, err = c.bot.Send(msg)
		return err
	}

	err = c.equipmentSellerService.Update(data)
	if err != nil {
		if errors.Is(err, service.ErrSellerNotFound) {
			msg.Text = fmt.Sprintf("seller with ID %q not found", data.ID.String())
		} else {
			msg.Text = "internal error"
		}
		_, err = c.bot.Send(msg)
		return err
	}

	msg.Text = "edit successful"
	_, err = c.bot.Send(msg)
	return err
}
