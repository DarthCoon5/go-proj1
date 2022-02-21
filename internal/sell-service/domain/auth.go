package domain

type Role int

const (
	Admin Role = iota
	Customer
)

func (role Role) IsAdmin() bool {
	if role == Admin {
		return true
	}

	return false
}
