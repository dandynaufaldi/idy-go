package model

type RatingID struct {
	id string
}

func NewRatingID(id string, IDProvider IDProvider) *RatingID {
	if id == "" {
		id = IDProvider.Generate()
	}
	return &RatingID{
		id: id,
	}
}

func (ratingID *RatingID) ID() string {
	return ratingID.id
}

func (ratingID *RatingID) Equals(other *RatingID) bool {
	return ratingID.ID() == other.ID()
}
