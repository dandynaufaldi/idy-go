package model

type Rating struct {
	user  string
	value *RatingValue
}

func NewRating(user string, value int) (*Rating, error) {
	ratingValue, err := NewRatingValue(value)
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

func (rating *Rating) Value() *RatingValue {
	return rating.value
}

func (rating *Rating) Equals(other *Rating) bool {
	return rating.Value().Equals(other.Value()) &&
		rating.User() == other.User()
}
