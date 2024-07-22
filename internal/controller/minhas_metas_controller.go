package controller

import (
	"net/http"

	"github.com/andrefelizard/myFin-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type minhasMetasController struct {
	minhasMetasUseCase usecase.MinhasMetasUseCase
}

func NewMinhasMetasController(usecase usecase.MinhasMetasUseCase) minhasMetasController {
	return minhasMetasController{
		minhasMetasUseCase: usecase,
	}
}

func (m *minhasMetasController) GetMinhasMetas(ctx *gin.Context) {

	metas, err := m.minhasMetasUseCase.GetMinhasMetas("a123")
	if(err != nil) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, metas)
}