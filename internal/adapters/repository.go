package adapters

import (
	"context"
	"github.com/Nav1Cr0ss/s-user-storage/internal/domain/user"
)

type Repository interface {
	GetUser(ctx context.Context, ID int) *user.User
}
