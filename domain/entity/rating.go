package entity

import (
	"github.com/dandynaufaldi/idy-go/domain/value"
)

type Rating struct {
	user  string
	value *value.RatingValue
}

func NewRating(user string, val int) (*Rating, error) {
	ratingValue, err := value.NewRatingValue(val)
	if err != nil {
		return nil, err
	}
	return &Rating{
		user:  user,
		value: ratingValue,
	}, nil
}

func (rating *Rating) User() string {
	return rating.user
}

func (rating *Rating) Value() *value.RatingValue {
	return rating.value
}

func (rating *Rating) Equals(other *Rating) bool {
	return rating.Value().Equals(other.Value()) &&
		rating.User() == other.User()
}
