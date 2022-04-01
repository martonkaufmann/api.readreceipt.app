package serializer

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type JSON struct {
	echo.DefaultJSONSerializer
}

func (j JSON) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := json.NewEncoder(c.Response())
	enc.SetEscapeHTML(false)
	if indent != "" {
		enc.SetIndent("", indent)
	}

	return enc.Encode(i)
}

func (j JSON) Deserialize(c echo.Context, i interface{}) error {
	return j.DefaultJSONSerializer.Deserialize(c, i)
}
