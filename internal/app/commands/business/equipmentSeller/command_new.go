package equipmentSeller

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	service "github.com/ozonmp/omp-bot/internal/service/business/equipmentSeller"
)

func (c *Commander) New(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "")

	args := inputMessage.CommandArguments()

	var data service.EquipmentSeller
	err := json.Unmarshal([]byte(args), &data)
	if err != nil {
		msg.Text = `invalid input data, it must be format: {"Name": "foobar"}`
		_, err = c.bot.Send(msg)
		return err
	}

	id, err := c.equipmentSellerService.Create(data)
	if err != nil {
		return err
	}

	msg.Text = fmt.Sprintf("equipment seller created with id: %s", id.String())
	_, err = c.bot.Send(msg)
	return err
}
