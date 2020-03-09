package usescases

import (
	"fmt"
	"github.com/fmcarrero/bookstore_users-api/application/commands"
	"github.com/fmcarrero/bookstore_users-api/application/factory"
	"github.com/fmcarrero/bookstore_users-api/domain/model"
	"github.com/fmcarrero/bookstore_users-api/domain/ports"
)

type CreatesUseCase interface {
	Handler(userCommand commands.UserCommand) (model.User, error)
}

type UseCaseUserCreate struct {
	UserRepository ports.UsersRepository
}

func (createsUseCase *UseCaseUserCreate) Handler(userCommand commands.UserCommand) (model.User, error) {
	fmt.Println(userCommand)
	user, err := factory.CreateUser(userCommand)
	if err != nil {
		return model.User{}, err
	}
	createUserErr := createsUseCase.UserRepository.Save(&user)
	if createUserErr != nil {
		return model.User{}, createUserErr
	}
	return user, nil
}
