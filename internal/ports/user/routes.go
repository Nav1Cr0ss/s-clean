package user

import (
	"github.com/Nav1Cr0ss/s-user-storage/pkg/xecho"
	"net/http"
	//"github.com/labstack/echo/v4"
)

func (h *Handler) GetRoutes() []xecho.Route {
	//commonMW := []echo.MiddlewareFunc{middleware.SetStringID}

	return []xecho.Route{
		// Get user accounts
		{
			Method:        http.MethodGet,
			Path:          "/user",
			Handle:        h.getUser,
			NoCORSHandler: true,
			//Middlewares: []echo.MiddlewareFunc{
			//	middleware.SetListParams,
			//},
		},
	}
}
