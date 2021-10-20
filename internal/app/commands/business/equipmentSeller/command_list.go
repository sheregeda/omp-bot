package equipmentSeller

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) error {
	d, err := json.Marshal(CallbackListData{Offset: 0})
	if err != nil {
		return err
	}
	p := path.CallbackPath{
		Domain:       "business",
		Subdomain:    "equipmentSeller",
		CallbackName: "list",
		CallbackData: string(d),
	}
	return c.CallbackList(&tgbotapi.CallbackQuery{Message: inputMessage}, p)
}
