package user

import (
	"context"
	"github.com/Nav1Cr0ss/s-user-storage/internal/domain/user"
)

type UserService interface {
	Get(ctx context.Context, userID string) (*user.User, error)
}
