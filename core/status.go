package core

type UserStatus string

const (
	ACTIVE  UserStatus = "ACTIVE"
	DELETE  UserStatus = "DELETE"
	LOCKED  UserStatus = "LOCKED"
	REJECT  UserStatus = "REJECT"
	UNLOCK  UserStatus = "UNLOCK"
	PENDING UserStatus = "PENDING"
)

func (s UserStatus) IsValid() bool {
	switch s {
	case ACTIVE, DELETE, LOCKED, REJECT, UNLOCK, PENDING:
		return true
	default:
		return false
	}
}
