package controller

import (
	"net/http"

	"github.com/andrefelizard/myFin-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ativosController struct {
	ativosUseCase usecase.AtivosUseCase
}

func NewAtivosController(usecase usecase.AtivosUseCase) ativosController {
	return ativosController{
		ativosUseCase: usecase,
	}
}

func (a *ativosController) GetAtivos(ctx *gin.Context) {

	ativos, err := a.ativosUseCase.GetAtivos()
	if(err != nil) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ativos)
}