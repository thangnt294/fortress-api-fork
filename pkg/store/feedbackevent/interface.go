package feedbackevent

import (
	"gorm.io/gorm"

	"github.com/dwarvesf/fortress-api/pkg/model"
)

type IStore interface {
	IsExist(db *gorm.DB, id string) (bool, error)
	GetBySubtypeWithPagination(db *gorm.DB, subtype string, pagination model.Pagination) (events []*model.FeedbackEvent, total int64, err error)
	Create(db *gorm.DB, feedbackEvent *model.FeedbackEvent) (*model.FeedbackEvent, error)
	One(db *gorm.DB, id string) (*model.FeedbackEvent, error)
}