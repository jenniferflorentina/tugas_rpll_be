package model

import "errors"

type Role int

const (
	ADMIN Role = iota
	CASHIER
	OTHERS
)

func (r Role) String() string {
	return [...]string{"Cashier", "Admin"}[r]
}

func (r Role) EnumIndex() int {
	return int(r)
}

func RoleFromString(s string) (Role, error) {
	switch s {
	case "Cashier":
		return CASHIER, nil
	case "Admin":
		return ADMIN, nil
	default:
		return OTHERS, errors.New("Role not supported")
	}
}
