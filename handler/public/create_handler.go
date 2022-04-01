package public

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/readreceipt/api/model"
	"github.com/readreceipt/api/repository"
	"github.com/readreceipt/api/service/monitoring"
	"github.com/readreceipt/api/service/receipt"
)

func CreateHandler(c echo.Context) error {
	request := &struct {
		Email string `json:"email" validate:"required,email"`
	}{}

	if err := c.Bind(request); err != nil {
		monitoring.CaptureError(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := c.Validate(request); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	u := model.User{
		Email: request.Email,
	}
	r := model.Receipt{
		Timestamp: time.Now().Unix(),
		ID:        uuid.NewString(),
	}

	// TODO: Store IP and some other identification
	err := repository.UpsertReceipt(u, r)

	if err != nil {
		monitoring.CaptureError(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"url": receipt.BuildReadURL(c, u, r),
	})
}
