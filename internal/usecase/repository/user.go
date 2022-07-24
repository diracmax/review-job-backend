package repository

type UserRepository interface {
	Save(name string, rawPassword string) error
}
