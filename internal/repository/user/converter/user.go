package converter

import (
	"github.com/nikita-reshetnyak/auth/internal/model"
	modelRepo "github.com/nikita-reshetnyak/auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      model.Role(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func ToModelRole(role modelRepo.Role) model.Role {
	switch role {
	case modelRepo.RoleUser:
		return model.RoleUser
	case modelRepo.RoleAdmin:
		return model.RoleAdmin
	default:
		return ""
	}
}
