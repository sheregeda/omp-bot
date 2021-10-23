package equipmentSeller

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var ErrSellerNotFound = errors.New("equipment seller not found")

type DummyEquipmentSellerService struct {
	data []EquipmentSeller
	mu   sync.RWMutex
}

func NewDummyEquipmentSellerService(initData []EquipmentSeller) *DummyEquipmentSellerService {
	if initData == nil {
		initData = make([]EquipmentSeller, 0)
	}
	return &DummyEquipmentSellerService{data: initData}
}

func (s *DummyEquipmentSellerService) Describe(sellerID uuid.UUID) (EquipmentSeller, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, es := range s.data {
		if es.ID == sellerID {
			return es, nil
		}
	}
	return EquipmentSeller{}, ErrSellerNotFound
}

func (s *DummyEquipmentSellerService) Count() uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return uint64(len(s.data))
}

func (s *DummyEquipmentSellerService) List(offset uint64, limit uint64) ([]EquipmentSeller, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	lenData := uint64(len(s.data))
	start := offset
	if start > lenData-1 {
		start = lenData - 1
	}
	end := start + limit
	if end > lenData {
		end = lenData
	}

	res := make([]EquipmentSeller, end-start)
	copy(res, s.data[start:end])

	return res, nil
}

func (s *DummyEquipmentSellerService) Create(es EquipmentSeller) (uuid.UUID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.New()
	s.data = append(s.data, EquipmentSeller{ID: id, Name: es.Name})
	return id, nil
}

func (s *DummyEquipmentSellerService) Remove(sellerID uuid.UUID) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, es := range s.data {
		if es.ID == sellerID {
			s.data = append(s.data[:i], s.data[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (s *DummyEquipmentSellerService) Update(seller EquipmentSeller) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, es := range s.data {
		if es.ID == seller.ID {
			s.data[i] = EquipmentSeller{ID: seller.ID, Name: seller.Name}
			return nil
		}
	}
	return ErrSellerNotFound
}
