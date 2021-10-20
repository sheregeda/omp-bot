package equipmentSeller

import (
	"errors"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/business/equipmentSeller"
)

type Service interface {
	Count() uint64
	Describe(sellerID uuid.UUID) (service.EquipmentSeller, error)
	List(offset uint64, limit uint64) ([]service.EquipmentSeller, error)
	Create(seller service.EquipmentSeller) (uuid.UUID, error)
	Remove(sellerID uuid.UUID) (bool, error)
	Update(seller service.EquipmentSeller) error
}

type Commander struct {
	bot                    *tgbotapi.BotAPI
	equipmentSellerService Service
}

func NewEquipmentSellerCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot:                    bot,
		equipmentSellerService: service.NewDummyEquipmentSellerService(service.GetInitEquipmentSellers(20)),
	}
}

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var err error
	switch callbackPath.CallbackName {
	case "list":
		err = c.CallbackList(callback, callbackPath)
	default:
		err = errors.New(fmt.Sprintf("EquipmentSellerCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName))
	}
	if err != nil {
		log.Printf("callback %s exec fail, err is %v", callbackPath.CallbackName, err)
	}
}

func (c *Commander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	var err error
	switch commandPath.CommandName {
	case "help":
		err = c.Help(msg)
	case "list":
		err = c.List(msg)
	case "get":
		err = c.Get(msg)
	case "new":
		err = c.New(msg)
	case "delete":
		err = c.Delete(msg)
	case "edit":
		err = c.Edit(msg)
	default:
		err = c.Default(msg)
	}
	if err != nil {
		log.Printf("command %s failed with error: %v", commandPath.CommandName, err)
	}
}
