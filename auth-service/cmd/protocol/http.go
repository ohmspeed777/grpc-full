package protocol

import (
	"app/handler"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ohmspeed777/go-pkg/logx"
	"github.com/ohmspeed777/go-pkg/middlewares"
	"github.com/tylerb/graceful"
)

func NewAPI() {
	e := echo.New()
	// e.HTTPErrorHandler = middleware.ErrorHandler(app.conf.App.Prefix)
	e.Use(middlewares.LogRequestMiddleware(app.key))
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middlewares.LogResponseMiddleware())
	e.HTTPErrorHandler = middlewares.CustomHTTPErrorHandler
	e.HideBanner = true
	// jwtMiddleware := jwtx.NewJWT(app.key)

	hdl := handler.NewHandler(handler.Dependencies{
		Config:  *app.Config,
		Service: app.Service,
	})

	baseAPI := e.Group("/api/v1")

	users := baseAPI.Group("/users")
	{
		users.POST("/sign-up", hdl.User.SignUp)
		users.POST("/sign-in", hdl.User.SignIn)
	}

	stream := baseAPI.Group("/stream")
	{
		stream.GET("/client", hdl.User.ClientStream)
		stream.GET("/bi", hdl.User.BidirectionalStream)
	}

	logx.GetLog().Infof("server starting on port: %d", app.Config.APP.APIPort)
	e.Server.Addr = fmt.Sprintf(":%d", app.Config.APP.APIPort)
	_ = graceful.ListenAndServe(e.Server, 5*time.Second)
}
