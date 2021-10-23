package business

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/equipmentSeller"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type EquipmentSellerCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type Commander struct {
	bot                      *tgbotapi.BotAPI
	equipmentSellerCommander EquipmentSellerCommander
}

func NewBusinessCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot:                      bot,
		equipmentSellerCommander: equipmentSeller.NewEquipmentSellerCommander(bot),
	}
}

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "equipmentSeller":
		c.equipmentSellerCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BusinessCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *Commander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "equipmentSeller":
		c.equipmentSellerCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BusinessCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
