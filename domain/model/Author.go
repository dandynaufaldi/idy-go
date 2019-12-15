package model

type Author struct {
	name  string
	email string
}

func NewAuthor(name, email string) *Author {
	return &Author{
		name:  name,
		email: email,
	}
}

func (author *Author) Name() string {
	return author.name
}

func (author *Author) Email() string {
	return author.email
}
