package main

import (
	"net/http"

	"github.com/ckn01/belajar-golang-restful-api/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

// func main() {
// 	server := InitializedServer()
// 	err := server.ListenAndServe()
// 	helper.PanicIfError(err)
// }

func main() {

	// 	db := app.NewDB()
	// 	validate := validator.New()
	// 	categoryRepository := repository.NewCategoryRepository()
	// 	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// 	categoryController := controller.NewCategoryController(categoryService)
	// 	router := app.NewRouter(categoryController)

	// 	server := http.Server{
	// 		Addr:    "localhost:3000",
	// 		Handler: middleware.NewAuthMiddleware(router),
	// 	}

	// err := server.ListenAndServe()
	// helper.PanicIfError(err)
}
