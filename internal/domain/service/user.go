package service

import (
	"residential-manager/internal/common/response"
	request "residential-manager/internal/common/srvRequest"
	"residential-manager/internal/common/validator"
	"residential-manager/internal/domain/entities"
	"residential-manager/internal/domain/ports/notification"
	repo "residential-manager/internal/domain/ports/repository"
	"time"

	"github.com/google/uuid"
)

const (
	ErrSendingMail    = "There was a problem sending the mail"
	ErrDuplicatedMail = "The user already exists"
	ErrNotValidToken  = "The token providen is not valid"
)

type UserSrv interface {
	GetRoles(req request.Request) (response response.Response[[]string])
	GetAllPermissions(req request.Request) (response response.Response[map[string]int])

	GetUsers(req request.Request) (response response.Response[[]entities.User])
	GetUserById(req request.Request, id uuid.UUID) (response response.Response[entities.User])
	CreateUser(req request.Request, rol, mail, block, number string) (response response.Response[entities.User])
	UpdateUser(req request.Request, user entities.User) (response response.Response[entities.User])
	DeleteUser(req request.Request, id uuid.UUID) (response response.Response[bool])

	GetVerifyToken(token string) (response response.Response[entities.User])
	VerifyToken(token string, password string) (response response.Response[bool])
	//VerifyUser(req request.Request, id uuid.UUID, password string) (response response.Response[string])
}

type MailData struct {
	Path     string
	LogoURL  string
	TokenURL string
}

type UserService struct {
	repo                repo.UserRepo
	template            MailData
	checkPermissionFunc CheckPermissionFunc
	notificator         notification.MailNotification
}

func NewUserService(repo repo.UserRepo, notificator notification.MailNotification, checkPermissionFunc CheckPermissionFunc, mailData MailData) *UserService {
	return &UserService{
		repo:                repo,
		notificator:         notificator,
		checkPermissionFunc: checkPermissionFunc,
		template:            mailData,
	}
}

func (srv *UserService) GetUserById(req request.Request, id uuid.UUID) (response response.Response[entities.User]) {
	allow, err := srv.checkPermissionFunc(req.Mail, entities.QueryUsers)
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if !allow {
		return response.ResponseUnauthorized(entities.NotAllowedQueryUsers)
	}

	user, err := srv.repo.GetUserById(id.String())
	if err != nil {
		return response.ResponseError(err.Error())
	}

	if user == nil {
		return response.ResponseNotFound()
	}

	return response.ResponseSuccess(*user)
}

func (srv *UserService) GetUsers(req request.Request) (response response.Response[[]entities.User]) {
	//check permissions
	allow, err := srv.checkPermissionFunc(req.Mail, entities.QueryUsers)
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if !allow {
		return response.ResponseUnauthorized(entities.NotAllowedQueryUsers)
	}

	users, err := srv.repo.GetUsers()
	if err != nil {
		return response.ResponseError(err.Error())
	}

	return response.ResponseSuccess(users)
}

func (srv *UserService) CreateUser(req request.Request, rol, mail, block, number string) (response response.Response[entities.User]) {
	allow, notAllowMessage, err := srv.checkRolPermit(req.Mail, rol)
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if !allow {
		return response.ResponseUnauthorized(notAllowMessage)
	}

	//check user info
	user, validations := entities.CreateUser(rol, mail, block, number)
	if validations != nil {
		return response.ResponseFail(validations)
	}

	//saveUser
	verifyUser, err := srv.repo.GetUserByMail(user.Mail)
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if verifyUser != nil {
		return response.ResponseFail(map[string]string{"mail": ErrDuplicatedMail})
	}

	err = srv.repo.CreateUser(user)
	if err != nil {
		return response.ResponseError(err.Error())
	}

	//send email
	to := []string{user.Mail}
	err = srv.notificator.SendHTMLMail(to, "New Residential Account", map[string]any{
		"logo":    srv.template.LogoURL,
		"baseURL": srv.template.TokenURL,
		"token":   user.Token,
	}, srv.template.Path)
	if err != nil {
		return response.ResponseError(ErrSendingMail)
	}

	return response.ResponseSuccess(*user)
}

func (srv *UserService) GetAllPermissions(req request.Request) (response response.Response[map[string]int]) {
	//check permissions
	allow, err := srv.checkPermissionFunc(req.Mail, entities.QueryPermissions)
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if !allow {
		return response.ResponseUnauthorized(entities.NotAllowedQueryPermissions)
	}

	permissions := entities.GetAllPermissions()
	return response.ResponseSuccess(permissions)
}

func (srv *UserService) GetRoles(req request.Request) (response response.Response[[]string]) {
	//check permissions
	allow, err := srv.checkPermissionFunc(req.Mail, entities.QueryPermissions)
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if !allow {
		return response.ResponseUnauthorized(entities.NotAllowedQueryRoles)
	}

	roles := entities.GetAllRoles()
	return response.ResponseSuccess(roles)
}

func (srv *UserService) UpdateUser(req request.Request, user entities.User) (response response.Response[entities.User]) {
	//check permissions
	allow, err := srv.checkPermissionFunc(req.Mail, entities.UpdateUsers)
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if !allow {
		return response.ResponseUnauthorized(entities.NotAllowedQueryUsers)
	}

	oldUser, err := srv.repo.GetUserById(user.Id.String())
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if oldUser == nil {
		return response.ResponseNotFound()
	}

	user.Mail = oldUser.Mail
	//update just values that are not updated
	if user.Apartment == nil {
		user.Apartment = oldUser.Apartment
	}
	if user.Rol == nil {
		user.Rol = oldUser.Rol
	}

	v := validator.New()
	user.ValidateUserInfo(v)
	if !v.Valid() {
		return response.ResponseFail(v.Errors)
	}

	allow, notAllowMessage, err := srv.checkRolPermit(req.Mail, *user.Rol)
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if !allow {
		return response.ResponseUnauthorized(notAllowMessage)
	}

	err = srv.repo.UpdateUser(&user)
	if err != nil {
		return response.ResponseError(err.Error())
	}

	return response.ResponseSuccess(user)
}

func (srv *UserService) DeleteUser(req request.Request, id uuid.UUID) (response response.Response[bool]) {
	//check permissions
	allow, err := srv.checkPermissionFunc(req.Mail, entities.UpdateUsers)
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if !allow {
		return response.ResponseUnauthorized(entities.NotAllowedQueryUsers)
	}

	user, err := srv.repo.GetUserById(id.String())
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if user == nil {
		return response.ResponseNotFound()
	}

	err = srv.repo.DeleteUser(id)
	if err != nil {
		return response.ResponseError(err.Error())
	}

	return response.ResponseSuccess(true)
}

func (srv *UserService) checkRolPermit(creatorMail, rol string) (allow bool, message string, err error) {
	switch rol {
	case entities.RolAdministrador:
		allow, err = srv.checkPermissionFunc(creatorMail, entities.CreateAdmin)
		message = entities.NotAllowedCreateAdmin
	case entities.RolEncargadoApartamento:
		allow, err = srv.checkPermissionFunc(creatorMail, entities.CreateApartmentAdmin)
		message = entities.NotAllowedCreateApartmentAdmin
	case entities.RolResidente:
		allow, err = srv.checkPermissionFunc(creatorMail, entities.CreateResident)
		message = entities.NotAllowedCreateResident
	case entities.RolVigilante:
		allow, err = srv.checkPermissionFunc(creatorMail, entities.CreateWatchman)
		message = entities.NotAllowedCreateWatchman
	default:
		allow = false
		err = nil
		message = entities.ErrInvalidRol
	}

	return allow, message, err
}

func (srv *UserService) GetVerifyToken(token string) (response response.Response[entities.User]) {
	users, err := srv.repo.GetUserByToken(token)
	if err != nil {
		return response.ResponseError(err.Error())
	}
	if users == nil {
		return response.ResponseFailMessage(ErrNotValidToken)
	}

	return response.ResponseSuccess(*users)
}

func (srv *UserService) VerifyToken(token, password string) (response response.Response[bool]) {
	user, err := srv.repo.GetUserByToken(token)
	if err != nil {
		return response.ResponseError(err.Error())
	}

	if user == nil {
		return response.ResponseFailMessage(ErrNotValidToken)
	}

	if user.Token.Token == "" || user.Token.Expire.Before(time.Now()) {
		return response.ResponseFailMessage(ErrNotValidToken)
	}

	v := validator.New()
	entities.ValidatePassword(v, password)
	if !v.Valid() {
		return response.ResponseFail(v.Errors)
	}

	newPassword := user.HashPassword(password)
	err = srv.repo.UpdatePassword(user.Id, newPassword)
	if err != nil {
		return response.ResponseError(err.Error())
	}

	return response.ResponseSuccess(true)
}
