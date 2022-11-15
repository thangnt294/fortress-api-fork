package role

import (
	"gorm.io/gorm"

	"github.com/dwarvesf/fortress-api/pkg/model"
)

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB) IStore {
	return &store{
		db: db,
	}
}

// One get all positions
func (s *store) All() ([]*model.Role, error) {
	var roles []*model.Role
	return roles, s.db.Find(&roles).Error
}