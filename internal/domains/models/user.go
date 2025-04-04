package models

import "google.golang.org/protobuf/types/known/timestamppb"

type UserRole int32

type User struct {
	ID        int64
	Name      string
	Email     string
	Role      UserRole
	CreatedAt *timestamppb.Timestamp
	UpdatedAt *timestamppb.Timestamp
}

const (
	RoleUser  UserRole = 0
	RoleAdmin UserRole = 1
)

func (r UserRole) String() string {
	switch r {
	case RoleAdmin:
		return "admin"
	case RoleUser:
		return "user"
	default:
		return "unknown"
	}
}
func ParseUserRole(s string) UserRole {
	switch s {
	case "admin":
		return RoleAdmin
	case "user":
		return RoleUser
	default:
		return RoleUser
	}
}
