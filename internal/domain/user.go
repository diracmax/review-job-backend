package domain

type user struct {
	id   int64
	name string
}

func NewUser(id int64, name string) *user {
	return &user{
		id:   id,
		name: name,
	}
}
