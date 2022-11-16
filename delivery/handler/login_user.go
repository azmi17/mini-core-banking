package handler

import (
	"apex-ems-integration-clean-arch/delivery/handler/httpio"
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(ctx *gin.Context) {

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding form
	payload := web.LoginInput{}
	httpio.Bind(&payload)

	usecase := usecase.NewSysUserUsecase()
	user, er := usecase.Login(payload)

	resp := web.LoginResponse{}

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
