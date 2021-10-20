package equipmentSeller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__business__equipmentSeller - help\n"+
			"/list__business__equipmentSeller - equipment sellers list\n"+
			"/get__business__equipmentSeller {uuid} - get equipment seller by id\n"+
			"/delete__business__equipmentSeller {uuid} - delete equipment seller by id\n"+
			"/new__business__equipmentSeller {\"Name\": \"foobar\"} - add new equipment seller\n"+
			"/edit__business__equipmentSeller {\"ID\": \"UUID\", \"Name\": \"foobar\"} - edit equipment seller\n",
	)
	_, err := c.bot.Send(msg)
	return err
}
