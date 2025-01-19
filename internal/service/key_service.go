package service

import (
	"github.com/c4miloarriagada/keys-be/internal/domain"
)

type KeyService struct {
	keyRepository domain.KeyRepository
}

func NewKeyService(keyRepository domain.KeyRepository) *KeyService {
	return &KeyService{
		keyRepository: keyRepository,
	}
}

func (s *KeyService) Save(key *domain.Key) error {
	return s.keyRepository.Save(key)
}

func (s *KeyService) GetKeyByID(id int64) (*domain.Key, error) {
	return s.keyRepository.GetByID(id)
}
