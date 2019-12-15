package model

// Idea represents Idea in domain model
type Idea struct {
	id            *IdeaID
	title         string
	description   string
	author        *Author
	totalVotes    int
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
