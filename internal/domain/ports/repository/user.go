package ports

import (
	"residential-manager/internal/domain/entities"

	"github.com/google/uuid"
)

type UserRepo interface {
	// GetRoles() ([]string, error)
	GetUsers() ([]entities.User, error)
	GetUserWithCredentials(mail string) (*entities.User, error)
	GetUserByMail(mail string) (*entities.User, error)
	GetUserById(id string) (*entities.User, error)
	GetUserByToken(token string) (*entities.User, error)
	// GetAllRoles() (entities.Rol, error)
	GetUserPermissions(mail string) ([]int, error)
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User) error
	DeleteUser(id uuid.UUID) error
	UpdatePassword(id uuid.UUID, password string) error
}
