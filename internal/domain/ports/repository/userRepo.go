package ports

import "residential-manager/internal/domain/entities"

type UserRepo interface {
	GetAllRoles() (entities.Permission, error)
	SaveUser(*entities.User) error
}
