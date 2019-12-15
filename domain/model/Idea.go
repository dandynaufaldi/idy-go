package model

import (
	"fmt"
)

// Idea represents Idea in domain model
type Idea struct {
	id            *IdeaID
	title         string
	description   string
	author        *Author
	totalVotes    int
	ratings       []*Rating
	averageRating float32
}

// NewIdea create fresh new idea
func NewIdea(title, description, authorName, authorEmail string, IDProvider IDProvider) *Idea {
	author := NewAuthor(authorName, authorEmail)
	return &Idea{
		id:          NewIdeaID("", IDProvider),
		title:       title,
		description: description,
		author:      author,
		ratings:     []*Rating{},
	}
}

func (idea *Idea) ID() *IdeaID {
	return idea.id
}

func (idea *Idea) Title() string {
	return idea.title
}

func (idea *Idea) Author() *Author {
	return idea.author
}

func (idea *Idea) TotalVotes() int {
	return idea.totalVotes
}

func (idea *Idea) AverageRating() float32 {
	return idea.averageRating
}

func (idea *Idea) Vote() {
	idea.totalVotes++
}

func (idea *Idea) AddRating(raterEmail string, value int) error {
	newRating, err := NewRating(raterEmail, value)
	if err != nil {
		return err
	}

	exist := false
	for _, rating := range idea.ratings {
		if rating.Equals(newRating) {
			exist = true
			break
		}
	}
	if exist {
		return &ErrUserHasRated{
			raterEmail: raterEmail,
			ideaID:     idea.ID(),
		}
	}

	idea.ratings = append(idea.ratings, newRating)
	idea.updateAverageRatings()
	return nil
}

func (idea *Idea) updateAverageRatings() {
	total := 0
	for _, rating := range idea.ratings {
		total += rating.Value().Value()
	}
	average := float32(total) / float32(len(idea.ratings))
	idea.averageRating = average
}

type ErrUserHasRated struct {
	raterEmail string
	ideaID     *IdeaID
}

func (e *ErrUserHasRated) Error() string {
	return fmt.Sprintf("User %s has rated idea with ID %s", e.raterEmail, e.ideaID.ID())
}
