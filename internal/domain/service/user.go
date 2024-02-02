package service

import (
	"errors"
	"residential-manager/internal/domain/entities"
	"residential-manager/internal/domain/ports/notification"
	repo "residential-manager/internal/domain/ports/repository"
)

const (
	ErrInvalidRol  = "Rol provided not exists"
	ErrSendingMail = "There was a problem sending the mail"
)

type EmailConfig struct {
	Mail     string //The email to send mails
	Password string //The password of the mail
	SmptHost string //Host ip
	SmptPort string //Host port
}

type UserService struct {
	repo               repo.UserRepo
	email              EmailConfig
	verifyUserTemplate string
	notificator        notification.MailNotification
}

func NewUserService(repo repo.UserRepo, email EmailConfig, template string) *UserService {
	return &UserService{
		repo:               repo,
		email:              email,
		verifyUserTemplate: template,
	}
}

func (serv *UserService) CreateUser(rol, mail, username string, permissions entities.Permission) (*entities.User, error) {
	//check permissions (tokens)

	//getRoles
	roles, err := serv.repo.GetAllRoles()
	if err != nil {
		return nil, err
	}

	//checkRoles
	if !roles.ArePermissionsPresent(permissions) {
		return nil, errors.New(ErrInvalidRol)
	}

	//check user info
	user, err := entities.CreateUser(rol, mail, username, permissions)
	if err != nil {
		return nil, err
	}

	//saveUser
	err = serv.repo.SaveUser(user)
	if err != nil {
		return nil, err
	}

	//send email
	to := []string{user.Mail}
	err = serv.notificator.SendMail(serv.email.Mail, serv.email.Password, serv.email.SmptHost, serv.email.SmptPort, to)
	if err != nil {
		//update the error with the library that is going to be use to wrap errors
		//errors.New(ErrSendingMail)
		return nil, err
	}

	return nil, nil
}

// verify the email of a new user
func (serv *UserService) VerifyEmail() {

}
