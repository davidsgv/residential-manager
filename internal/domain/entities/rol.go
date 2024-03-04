package entities

const (
	RolAdministrador        = "Administrador"
	RolVigilante            = "Vigilante"
	RolEncargadoApartamento = "Encargado"
	RolResidente            = "Residente"
)

func GetAllRoles() []string {
	return []string{
		RolAdministrador,
		RolVigilante,
		RolEncargadoApartamento,
		RolResidente,
	}
}
