package model

import (
	"fmt"
)

type RatingValue struct {
	value int
}

func NewRatingValue(value int) (*RatingValue, error) {
	if value < 0 || value > 5 {
		return nil, &ErrInvalidRatingValue{value: value}
	}
	return &RatingValue{value: value}, nil
}

func (r *RatingValue) Value() int {
	return r.value
}

func (r *RatingValue) Equals(other *RatingValue) bool {
	return r.Value() == other.Value()
}

type ErrInvalidRatingValue struct {
	value int
}

func (e *ErrInvalidRatingValue) Error() string {
	return fmt.Sprintf("rating is %d,  value must be between 0 to 5", e.value)
}
