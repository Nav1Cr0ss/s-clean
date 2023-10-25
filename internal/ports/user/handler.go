package user

import (
	"github.com/Nav1Cr0ss/s-user-storage/internal/service"
	"github.com/Nav1Cr0ss/s-user-storage/internal/service/user"
	logger "github.com/Nav1Cr0ss/s-user-storage/pkg/logger/zerolog"
	"github.com/Nav1Cr0ss/s-user-storage/pkg/xecho"
)

type Handler struct {
	log  logger.Logger
	user user.UserService
	//validator *validator.Validator
}

func NewHandler(service *service.Service) xecho.HandlerResult {
	h := &Handler{
		log:  logger.NewDefaultComponentLogger("account_handler"),
		user: service.User,
	}

	return xecho.HandlerResult{
		Handler: h,
	}
}
