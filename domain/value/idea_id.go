package value

import (
	"github.com/google/uuid"
)

type IdeaID struct {
	id string
}

func NewIdeaID(id string) *IdeaID {
	if id == "" {
		id = uuid.New().String()
	}
	return &IdeaID{
		id: id,
	}
}

func (ideaID *IdeaID) ID() string {
	return ideaID.id
}

func (ideaID *IdeaID) Equals(other *IdeaID) bool {
	return ideaID.ID() == other.ID()
}
