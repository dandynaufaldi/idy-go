package repository

import (
	"github.com/dandynaufaldi/idy-go/domain/entity"
	"github.com/dandynaufaldi/idy-go/domain/value"
)

type RatingRepository interface {
	Save(rating *entity.Rating, ideaID *value.IdeaID) error
}
