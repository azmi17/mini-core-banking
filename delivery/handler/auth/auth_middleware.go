package handler

import (
	"fmt"
	"net/http"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/helper"
	"new-apex-api/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthRequired(ctx *gin.Context) {
	getToken := ctx.GetHeader("x_token")
	getUserID := ctx.GetHeader("x_user_id")
	userID, _ := strconv.Atoi(getUserID)
	values := "token: " + getToken + " user_id: " + getUserID

	entities.PrintLog(fmt.Sprintf("%s, %s, %s, %s", "AUTHORIZATION", ctx.Request.RemoteAddr, ctx.Request.Method+":"+ctx.Request.URL.Path, values))

	usecase := usecase.NewAuthUsecase()
	er := usecase.HeaderValidation(getToken, userID)
	if er != nil {
		if er == err.HeaderRequired {
			response := helper.ApiResponse("header is required", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		} else if er == err.InvalidToken {
			response := helper.ApiResponse("invalid token", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		} else if er == err.UserIDDonthMatch {
			response := helper.ApiResponse("user id don't match", http.StatusBadRequest, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else if er == err.NoRecord {
			response := helper.ApiResponse("record not found", http.StatusNotFound, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		}
	}
}
