package controller

import (
	"net/http"
	"strconv"

	"github.com/andrefelizard/myFin-backend/internal/model"
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

func (a *ativosController) CreateAtivo(ctx *gin.Context) {
	var ativo model.Ativo
	err := ctx.BindJSON(&ativo)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedAtivo, err := a.ativosUseCase.CreateAtivo(ativo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, insertedAtivo)
}

func (a *ativosController) GetAtivoById(ctx *gin.Context) {
	id := ctx.Param("id")
	if(id == "") {
		response := model.Response {
			Message: "id is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ativoId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response {
			Message: "id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ativo, err := a.ativosUseCase.GetAtivoById(ativoId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if ativo == nil {
		response := model.Response {
			Message: "id must be a number",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, ativo)
}