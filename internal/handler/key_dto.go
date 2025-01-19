package handler

import "github.com/c4miloarriagada/keys-be/internal/domain"

type keyDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Pass        string `json:"pass"`
	Alias       string `json:"alias"`
}

func (k *keyDTO) toDomain() domain.Key {
	return domain.Key{
		Name:        k.Name,
		Description: k.Description,
		Pass:        k.Pass,
		Alias:       k.Alias,
	}
}

func NewResponseDTO(k domain.Key) keyDTO {
	return keyDTO{
		Name:        k.Name,
		Description: k.Description,
		Pass:        k.Pass,
		Alias:       k.Alias,
	}
}
