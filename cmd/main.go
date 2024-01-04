package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	config "github.com/inouttt/test-go-mezink/config"
	"github.com/inouttt/test-go-mezink/pkg/db"
	"github.com/inouttt/test-go-mezink/src/v1/records/handler"
	"github.com/inouttt/test-go-mezink/src/v1/records/repository"
	"github.com/inouttt/test-go-mezink/src/v1/records/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	customValidator struct {
		validator *validator.Validate
	}
)

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	appConf := config.InitConfig()

	dbConn := db.NewMongo(&db.MongoDBConfig{
		Host:     appConf.DbHost,
		Username: appConf.DbUser,
		Password: appConf.DbPass,
		Schema:   appConf.DbName,
	})
	defer func() {
		if err := dbConn.Client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	timeoutContext := time.Duration(appConf.ContextTimeout) * time.Second

	e := echo.New()
	newValidator := validator.New()
	e.Validator = &customValidator{validator: newValidator}

	configLimit := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 100, ExpiresIn: 1 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	e.Use(middleware.RateLimiterWithConfig(configLimit))

	// repo
	rRepo := repository.NewRecordMongo(dbConn)

	// usecase
	rUsecase := usecase.NewRecordUsecase(rRepo, timeoutContext)

	// handler
	handler.NewHttpOrderHandler(e, rUsecase)

	// health check
	e.GET("/", func(ctx echo.Context) error {
		dbStatus := false
		if dbConn.Client.Ping(ctx.Request().Context(), nil) == nil {
			dbStatus = true
		}

		return ctx.JSON(200, echo.Map{
			"http_server": true,
			"db_conn":     dbStatus,
		})
	})
	e.Logger.Fatal(e.Start(":" + appConf.ServerPort))
}
