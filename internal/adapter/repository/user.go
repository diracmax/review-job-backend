package repository

import (
	"fmt"
	"github.com/diracmax/review-job-backend/internal/usecase/repository"
)

var _ repository.UserRepository = &userRepository{}

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (u *userRepository) Save(name string, rawPassword string) error {
	fmt.Println(name, rawPassword)
	return nil
}
