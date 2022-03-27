package trending_stickers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "5")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
