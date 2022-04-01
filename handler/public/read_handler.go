package public

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/readreceipt/api/repository"
	"github.com/readreceipt/api/service/monitoring"
	"github.com/readreceipt/api/service/receipt"
	"github.com/readreceipt/api/service/signature"
)

func ReadHandler(c echo.Context) error {
	request := &struct {
		Timestamp string `query:"ts" validate:"required,number,len=10"`
		ID        string `query:"id" validate:"required,uuid4"`
		Email     string `query:"email" validate:"required,email"`
		Signature string `query:"signature" validate:"required,alphanum,len=64"`
	}{}
	response := func() error {
		c.Response().Header().Add("Cache-Control", "max-age=31536000")
		return c.File("/go/src/readreceipt/api/image.jpg")
	}

	if err := c.Bind(request); err != nil {
		monitoring.CaptureError(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := c.Validate(request); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if !signature.Verify(request.Signature, request.Timestamp, request.ID, request.Email) {
		return c.NoContent(http.StatusNotFound)
	}

	isRead, err := receipt.IsRead(request.Signature, request.ID)

	if err != nil {
		monitoring.CaptureError(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if isRead {
		return response()
	}

	if err := repository.UpdateSetReceiptRead(request.ID); err != nil {
		monitoring.CaptureError(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return response()
}
