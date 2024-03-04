package data

type PermissionsResponse struct {
	Name      string `json:"name"`
	Operation int    `json:"operation"`
}

func MapToPermissions(m map[string]int) (per []PermissionsResponse) {
	for key, value := range m {
		per = append(per, PermissionsResponse{Name: key, Operation: value})
	}
	return per
}
