package equipmentSeller

import (
	"encoding/json"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const sellersListLimitSize = 4

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *Commander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) error {
	var data CallbackListData
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &data)
	if err != nil {
		return err
	}

	sellers, err := c.equipmentSellerService.List(data.Offset, sellersListLimitSize)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "")

	if len(sellers) == 0 {
		msg.Text = "List is empty"
		_, err = c.bot.Send(msg)
		return err
	}

	var buf strings.Builder
	for _, es := range sellers {
		buf.WriteString(fmt.Sprintf("%s\n", es.String()))
	}
	msg.Text = buf.String()

	var nextData CallbackListData
	if data.Offset+sellersListLimitSize >= c.equipmentSellerService.Count() {
		nextData = CallbackListData{Offset: data.Offset}
	} else {
		nextData = CallbackListData{Offset: data.Offset + sellersListLimitSize}
	}
	next, err := json.Marshal(&nextData)
	if err != nil {
		return err
	}

	var prevData CallbackListData
	if data.Offset < sellersListLimitSize {
		prevData = CallbackListData{Offset: 0}
	} else {
		prevData = CallbackListData{Offset: data.Offset - sellersListLimitSize}
	}
	prev, err := json.Marshal(&prevData)
	if err != nil {
		return err
	}

	nextBtn := tgbotapi.NewInlineKeyboardButtonData("Next", path.CallbackPath{
		Domain:       "business",
		Subdomain:    "equipmentSeller",
		CallbackName: "list",
		CallbackData: string(next),
	}.String())
	prevBtn := tgbotapi.NewInlineKeyboardButtonData("Prev", path.CallbackPath{
		Domain:       "business",
		Subdomain:    "equipmentSeller",
		CallbackName: "list",
		CallbackData: string(prev),
	}.String())
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(prevBtn, nextBtn))
	_, err = c.bot.Send(msg)
	return err
}
