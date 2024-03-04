package entities_test

import (
	"residential-manager/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	description string
	input       input
	output      output
}

type input struct {
	mail   string
	rol    string
	block  string
	number string
}

type output struct {
	user        *entities.User
	validations map[string]string
}

func TestCreateUser(t *testing.T) {
	testCases := []test{
		{
			description: "Bad Mail",
			input:       input{"no es un correo", entities.RolAdministrador, "A", "101"},
			output: output{nil, map[string]string{
				"mail": entities.ErrInvalidMail,
			}}},
		{
			description: "Bad rol",
			input:       input{"email@example.com", "SysAdmin", "A", "101"},
			output: output{nil, map[string]string{
				"rol": entities.ErrInvalidRol,
			}},
		},
		{
			description: "Bad rol and mail",
			input:       input{"no es un correo", "SysAdmin", "A", "101"},
			output: output{nil, map[string]string{
				"rol":  entities.ErrInvalidRol,
				"mail": entities.ErrInvalidMail,
			}},
		},
		{
			description: "Good Path",
			input:       input{"email@example.com", entities.RolResidente, "A", "101"},
			output: output{&entities.User{
				Rol:          rolToPointer(entities.RolResidente),
				Mail:         "email@example.com",
				Password:     "",
				MailVerified: false,
				Apartment: &entities.Apartment{
					Block:  "A",
					Number: "101",
				},
			}, nil},
		},
	}

	for _, scenario := range testCases {
		t.Run(scenario.description, func(t *testing.T) {
			user, val := entities.CreateUser(scenario.input.rol, scenario.input.mail, scenario.input.block, scenario.input.number)

			// t.FailNow()
			if eq := assert.Equal(t, val, scenario.output.validations); !eq {
				t.FailNow()
			}

			if user == nil {
				if user == nil && scenario.output.user != nil {
					t.FailNow()
				}
				return
			}

			if user.Mail != scenario.output.user.Mail || user.Password != "" || user.Rol != scenario.output.user.Rol {
				t.FailNow()
			}

			if user.Apartment.Block != scenario.output.user.Apartment.Block || user.Apartment.Number != scenario.output.user.Apartment.Number {
				t.FailNow()
			}

		})
	}
}

func rolToPointer(s string) *string {
	var rol *string
	rol = &s
	return rol
}
