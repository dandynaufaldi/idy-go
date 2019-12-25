package repository

import (
	"github.com/dandynaufaldi/idy-go/domain/entity"
	"github.com/dandynaufaldi/idy-go/domain/value"
)

type IdeaRepository interface {
	ByID(ideaID *value.IdeaID) (*entity.Idea, error)
	Save(idea *entity.Idea) error
	All() ([]*entity.Idea, error)
}
