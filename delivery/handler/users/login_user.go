package handler

import (
	"net/http"
	"new-apex-api/delivery/handler/httpio"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/helper"
	"new-apex-api/usecase"

	"github.com/gin-gonic/gin"
)

func LoginUser(ctx *gin.Context) {

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding form
	payload := entities.LoginInput{}
	rerr := httpio.BindWithErr(&payload)
	if rerr != nil {
		errors := helper.FormatValidationError(rerr)
		errorMesage := gin.H{"errors": errors}
		response := helper.ApiResponse("login user failed", http.StatusUnprocessableEntity, "failed", errorMesage)
		httpio.Response(http.StatusUnprocessableEntity, response)
		return
	}

	usecase := usecase.NewSysUserUsecase()
	user, er := usecase.Login(payload)

	resp := entities.LoginResponse{}
	if er != nil {
		if er == err.NoRecord || er == err.PasswordDontMatch {
			resp.Response_Code = "1111"
			resp.Response_Msg = er.Error()
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
			return
		}
	} else {
		resp.Response_Code = "0000"
		resp.Response_Msg = "Successfully logged in"
		resp.Data = &user
	}

	httpio.Response(http.StatusOK, resp)
}
