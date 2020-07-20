package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"summerCourse/service"
)

func SelectGoods(ctx *gin.Context) {
	goods := service.SelectGoods()
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info": "success",
		"data": struct {
			Goods []service.Goods `json:"goods"`
		}{goods},
	})
}

 

