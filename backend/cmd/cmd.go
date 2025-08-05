package cmd

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/ddd-todolist/backend/api/handler"
	"github.com/crazyfrankie/ddd-todolist/backend/api/middleware"
	"github.com/crazyfrankie/ddd-todolist/backend/application"
)

func Init() (*gin.Engine, error) {
	ctx := context.Background()
	services, err := application.Init(ctx)
	if err != nil {
		return nil, err
	}

	userHandler := handler.NewUserHandler(services.UserSvc)
	taskHandler := handler.NewTaskHandler(services.TaskSvc)

	srv := gin.Default()
	srv.Use(middleware.CtxCache())
	srv.Use(middleware.CORS())
	srv.Use(middleware.NewAuthnHandler(services.Infra.JWTGen).
		IgnorePath("/api/user/register").
		IgnorePath("/api/user/login").
		JWTAuthMW())

	apiGroup := srv.Group("api")

	userHandler.RegisterRoute(apiGroup)
	taskHandler.RegisterRoute(apiGroup)

	return srv, nil
}
