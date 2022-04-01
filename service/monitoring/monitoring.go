package monitoring

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/readreceipt/api/config"
)

func Init() error {
	return sentry.Init(sentry.ClientOptions{
		Dsn:              config.Sentry(),
		TracesSampleRate: 0.2,
		Environment:      config.Env(),
	})
}

func CaptureError(err error) {
	if config.IsLocal() {
		fmt.Println(err)
	}

	sentry.CaptureException(err)
}
