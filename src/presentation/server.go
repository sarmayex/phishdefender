package presentation

import (
	"errors"
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net"
)

func Start() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	//g := e.Group("api/v1")

	//handler := di.DefenderHandler()
	//g.POST("/phishinger/store", handler.Store)

	// Serving the HTTP-Server
	go func(e *echo.Echo) {
		logrus.WithFields(logrus.Fields{
			"http_started_at": htenvier.Env("HTTP_ADDRESS"),
		}).Info("start http-server")
		err := e.Start(htenvier.EnvOrDefault("HTTP_ADDRESS", "0.0.0.0:8000"))
		if err != nil {
			var netOpError *net.OpError
			ok := errors.As(err, &netOpError)
			if ok && netOpError.Err.Error() != "use of closed network connection" {
				logrus.WithFields(logrus.Fields{
					"err": err,
				}).Panic("failed to serve Http")
			} else {
				logrus.Info("http server has been stopped")
			}
		}

	}(e)

}
