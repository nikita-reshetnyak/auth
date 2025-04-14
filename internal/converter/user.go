package converter

import (
	v1 "github.com/nikita-reshetnyak/auth/gen/auth_v1"
	"github.com/nikita-reshetnyak/auth/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromProtoRole(role v1.Role) model.Role {
	switch role {
	case v1.Role_ROLE_USER:
		return model.RoleUser
	case v1.Role_ROLE_ADMIN:
		return model.RoleAdmin
	default:
		return ""
	}
}
func ToUserFromService(user *model.User) *v1.GetResponse {
	var updated_at *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updated_at = timestamppb.New(user.UpdatedAt.Time)
	}
	return &v1.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      ToProtoRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updated_at,
	}
}
func ToProtoRole(role model.Role) v1.Role {
	switch role {
	case model.RoleUser:
		return v1.Role_ROLE_USER
	case model.RoleAdmin:
		return v1.Role_ROLE_ADMIN
	default:
		return v1.Role_ROLE_ADMIN
	}
}
