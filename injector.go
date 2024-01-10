//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/ckn01/belajar-golang-restful-api/app"
	"github.com/ckn01/belajar-golang-restful-api/controller"
	"github.com/ckn01/belajar-golang-restful-api/middleware"
	"github.com/ckn01/belajar-golang-restful-api/repository"
	"github.com/ckn01/belajar-golang-restful-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func provideValidatorOptions() []validator.Option {
	return []validator.Option{validator.WithRequiredStructEnabled()}
}

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		provideValidatorOptions,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
