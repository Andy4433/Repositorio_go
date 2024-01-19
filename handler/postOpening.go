package handler

import (
	"net/http"

	"github.com/Andy4433/Repositorio_go.git/schemas"
	"github.com/gin-gonic/gin"
)

func PostOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err!=nil{
		logger.Errorf("validation error: %v", err.Error())
		sendErro(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,

	}

	if err := db.Create(&opening).Error; err!=nil{
		logger.Errorf("error creating opening: %v", err.Error())
		sendErro(ctx, http.StatusInternalServerError, "error creating opening on database")
		return 
	}
}