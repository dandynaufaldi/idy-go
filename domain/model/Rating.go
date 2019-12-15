package model

type Rating struct {
	user  string
	value int
}

func NewRating(user string, value int) *Rating {
	return &Rating{
		user:  user,
		value: value,
	}
}

func (rating *Rating) User() string {
	return rating.user
}

func (rating *Rating) Value() int {
	return rating.value
}

func (rating *Rating) Equals(other *Rating) bool {
	return rating.Value() == other.Value() && rating.User() == other.User()
}
