// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/go-api-template/src/dao"
	"github.com/gsxhnd/go-api-template/src/handler"
	"github.com/gsxhnd/go-api-template/src/middleware"
	"github.com/gsxhnd/go-api-template/src/routes"
	"github.com/gsxhnd/go-api-template/src/service"
	"github.com/gsxhnd/go-api-template/src/utils"
)

// Injectors from wire.go:

func InitApp() (*Application, error) {
	engine := gin.New()
	config, err := utils.NewConfig()
	if err != nil {
		return nil, err
	}
	logger := utils.NewLogger(config)
	tracer, err := utils.NewTracer(config)
	if err != nil {
		return nil, err
	}
	middlewarer := middleware.NewMiddleware(logger, tracer)
	rootHandler := handler.NewRootHandle(logger)
	database := dao.NewDatabase(logger)
	demoService := service.NewDemoService(logger, database)
	demoHandler := handler.NewDemoHandle(logger, demoService)
	routesRoutes := &routes.Routes{
		Engine:     engine,
		Middleware: middlewarer,
		RootHandle: rootHandler,
		DemoHandle: demoHandler,
	}
	application := NewApplication(routesRoutes)
	return application, nil
}

func InitTask() (*Task, error) {
	task := NewTask()
	return task, nil
}
