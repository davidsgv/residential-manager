package service_test

// import (
// 	"residential-manager/internal/common/response"
// 	"residential-manager/internal/domain/entities"
// 	"residential-manager/internal/domain/service"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// type mockUserRepo struct {
// 	GetUsersFunc      func() ([]entities.User, error)
// 	CreateUserFunc    func(user *entities.User) error
// 	GetUserByMailFunc func(mail string) (*entities.User, error)
// }

// func (m mockUserRepo) GetUsers() ([]entities.User, error) {
// 	return m.GetUsersFunc()
// }
// func (m mockUserRepo) GetUserByMail(mail string) (*entities.User, error) {
// 	return m.GetUserByMailFunc(mail)
// }
// func (m mockUserRepo) CreateUser(user *entities.User) error {
// 	return m.CreateUserFunc(user)
// }

// type testGetUsers struct {
// 	description         string
// 	checkPermissionFunc service.CheckPermissionFunc
// 	getUserFunc         func() ([]entities.User, error)
// 	output              response.Response[[]entities.User]
// }

// var checkPermissionFunc service.CheckPermissionFunc = func(mail string, permissions ...int) (bool, error) {
// 	return true, nil
// }

// func TestCreateUser(t *testing.T) {
// 	testCases := []testGetUsers{
// 		{
// 			description:         "nil users",
// 			checkPermissionFunc: checkPermissionFunc,
// 			getUserFunc: func() ([]entities.User, error) {
// 				return []entities.User{}, nil
// 			},
// 			output: response.Response[[]entities.User]{
// 				Status:      response.StatusSuccess,
// 				Data:        &[]entities.User{},
// 				Message:     "",
// 				Validations: nil,
// 			},
// 		},
// 		{
// 			description:         "return users",
// 			checkPermissionFunc: checkPermissionFunc,
// 			getUserFunc: func() ([]entities.User, error) {
// 				return []entities.User{
// 					{Rol: entities.RolAdministrador, Mail: "example@gmail.com", Password: "", MailVerified: false},
// 					{Rol: entities.RolEncargadoApartamento, Mail: "example2@gmail.com", Password: "", MailVerified: false},
// 				}, nil
// 			},
// 			output: response.Response[[]entities.User]{
// 				Status: response.StatusSuccess,
// 				Data: &[]entities.User{
// 					{Rol: entities.RolAdministrador, Mail: "example@gmail.com", Password: "", MailVerified: false},
// 					{Rol: entities.RolEncargadoApartamento, Mail: "example2@gmail.com", Password: "", MailVerified: false},
// 				},
// 				Message:     "",
// 				Validations: nil,
// 			},
// 		},
// 		{
// 			description:         "Not logged in",
// 			checkPermissionFunc: checkPermissionFunc,
// 			getUserFunc: func() ([]entities.User, error) {
// 				return nil, nil
// 			},
// 			output: response.Response[[]entities.User]{
// 				Status:      response.StatusUnauthorized,
// 				Data:        nil,
// 				Message:     service.ErrNotLogedIn,
// 				Validations: nil,
// 			},
// 		},
// 	}

// 	for _, scenario := range testCases {
// 		t.Run(scenario.description, func(t *testing.T) {
// 			service := service.NewUserService(mockUserRepo{GetUsersFunc: scenario.getUserFunc}, nil, scenario.loginFunc, "")
// 			res := service.GetUsers()

// 			if eq := assert.Equal(t, res, scenario.output); !eq {
// 				t.FailNow()
// 			}
// 		})
// 	}
// }
