// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/CyanAsterisk/FreeCar/server/shared/errno"
	"github.com/CyanAsterisk/FreeCar/server/shared/tools"
	"net/http"

	"github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/handler"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	r.NoRoute(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 404
		c.JSON(http.StatusNotFound, tools.BuildBaseResp(errno.NoRoute))
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 405
		c.JSON(http.StatusMethodNotAllowed, tools.BuildBaseResp(errno.NoMethod))
	})
}
