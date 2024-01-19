package handler

import "github.com/gin-gonic/gin"

func sendErro(ctx *gin.Context, code int, msg string){
	ctx.Header("Content-type", "application/json")

	ctx.JSON(code, gin.H{
		"message": msg,
		"erroCode": code,
	})
}