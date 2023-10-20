package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) getUser(c echo.Context) error {
	res, err := h.user.Get(c.Request().Context(), "1")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}
