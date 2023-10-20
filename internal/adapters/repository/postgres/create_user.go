package postgres

import (
	"context"
	"errors"
	"github.com/Nav1Cr0ss/s-user-storage/internal/domain/user"
)

func (r *Repository) GetUser(ctx context.Context, ID int) *user.User {

	r.log.Info().Msg("kjadbfkjbsd")
	r.log.Error().Err(errors.New("nf")).Msg("failed to send email ")

	return nil
}
