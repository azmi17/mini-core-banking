package handler

import (
	"apex-ems-integration-clean-arch/delivery/handler/httpio"
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResetApexPassword(ctx *gin.Context) {

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding form
	payload := web.KodeLKMFilter{}
	httpio.Bind(&payload)

	usecase := usecase.NewSysUserUsecase()
	user, er := usecase.ResetSysUserPassword(payload)
	if er != nil {
		entities.PrintError(er.Error())
		entities.PrintLog(er.Error())
		httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
	} else {
		httpio.Response(http.StatusOK, user)
	}

}
