package model

type IdeaID struct {
	id string
}

func NewIdeaID(id string, IDProvider IDProvider) *IdeaID {
	if id == "" {
		id = IDProvider.Generate()
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
