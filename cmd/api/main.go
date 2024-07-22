package main

import (
	"github.com/andrefelizard/myFin-backend/configs/db"
	"github.com/andrefelizard/myFin-backend/internal/controller"
	"github.com/andrefelizard/myFin-backend/internal/repository"
	"github.com/andrefelizard/myFin-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	 }

	MinhasMetasRepository := repository.NewMinhasMetasRepository(dbConnection)

	MinhasMetasUseCase := usecase.NewMinhasMetasUseCase(MinhasMetasRepository)
	MinhasMetasController := controller.NewMinhasMetasController(MinhasMetasUseCase)

	AtivosRepository := repository.NewAtivosRepository(dbConnection)
	AtivosUseCase := usecase.NewAtivosUseCase(AtivosRepository)
	AtivosController := controller.NewAtivosController(AtivosUseCase)

	server.GET("/minhas-metas", MinhasMetasController.GetMinhasMetas)
	server.GET("/ativos", AtivosController.GetAtivos)

	server.Run(":8080")
}