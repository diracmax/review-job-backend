package usecase

import "github.com/diracmax/review-job-backend/internal/usecase/repository"

type UserUsecase interface {
	Register(name string, rawPassword string) error
}

var _ UserUsecase = &userUsecase{}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repository: repo,
	}
}

func (u *userUsecase) Register(name string, rawPassword string) error {
	err := u.repository.Save(name, rawPassword)
	return err
}
