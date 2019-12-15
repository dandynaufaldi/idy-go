package model

type IdeaRepository interface {
	ByID(ideaID *IdeaID) (*Idea, error)
	Save(idea *Idea) error
	All() ([]*Idea, error)
}
