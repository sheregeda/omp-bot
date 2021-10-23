package equipmentSeller

import (
	"fmt"

	"github.com/google/uuid"
)

type EquipmentSeller struct {
	ID   uuid.UUID
	Name string
}

func (e *EquipmentSeller) String() string {
	return fmt.Sprintf("ID: %s | Name: %s", e.ID.String(), e.Name)
}

func GetInitEquipmentSellers(count uint) []EquipmentSeller {
	initData := make([]EquipmentSeller, 0, count)
	for i := 0; i < int(count); i++ {
		initData = append(initData, EquipmentSeller{ID: uuid.New(), Name: fmt.Sprintf("seller_%d", i)})
	}
	return initData
}
