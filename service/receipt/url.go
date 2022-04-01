package receipt

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/readreceipt/api/config"
	"github.com/readreceipt/api/model"
	"github.com/readreceipt/api/service/signature"
)

func BuildReadURL(c echo.Context, u model.User, r model.Receipt) string {
	s := signature.Create(fmt.Sprint(r.Timestamp), r.ID, u.Email)
	// path := c.Echo().Reverse(handler.ReadHandlerName)
	path := "/public/read.png"
	ts := fmt.Sprint(r.Timestamp)

	return fmt.Sprintf(
		"%s%s?ts=%s&id=%s&email=%s&signature=%s",
		config.URL(), path, ts, r.ID, u.Email, s,
	)
}
