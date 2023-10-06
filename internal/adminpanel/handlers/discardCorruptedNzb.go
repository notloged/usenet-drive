package handlers

import (
	"net/http"
	"strconv"

	"github.com/javi11/usenet-drive/internal/usenet/corruptednzbsmanager"
	echo "github.com/labstack/echo/v4"
)

func DiscardCorruptedNzbHandler(cNzb corruptednzbsmanager.CorruptedNzbsManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if _, err := cNzb.Discard(c.Request().Context(), idInt); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
