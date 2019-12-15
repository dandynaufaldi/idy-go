package model

type RatingRepository interface {
	Save(rating *Rating, ideaID *IdeaID) error
}
