package services

import (
	"github.com/google/wire"
	"haolinju.xianhaohan.com/internal/app/models"
)

var Provider = wire.NewSet(New)

// Service service.
type Service struct {
	model *models.Model
}

// New new a service and return.
func New(m *models.Model) (s *Service, err error) {
	s = &Service{
		model: m,
	}

	return
}
