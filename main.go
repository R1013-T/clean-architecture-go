package main

import (
	"clean-architecture/controller"
	"clean-architecture/db"
	"clean-architecture/repository"
	"clean-architecture/router"
	"clean-architecture/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
