package user

import (
	"context"
	"errors"
	"github.com/Nav1Cr0ss/s-user-storage/internal/domain/user"
)

func (u *UserApp) Get(ctx context.Context, accountID string) (*user.User, error) {

	u.repo.GetUser(ctx, 12)
	return nil, errors.New("sknd")
}
