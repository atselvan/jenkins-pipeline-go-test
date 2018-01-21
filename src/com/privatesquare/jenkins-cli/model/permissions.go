package model

type RolePermissions struct {
	Description        string   `json:"description"`
	Filterable         bool     `json:"filterable"`
	GrantedPermissions []string `json:"grantedPermissions"`
	ID                 string   `json:"id"`
}
