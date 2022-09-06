package handler

import (
	"apex-ems-integration-clean-arch/delivery/handler/httpio"
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/helper"
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

	usecase := usecase.NewApexUsecase()
	user, er := usecase.ResetApexPassword(payload)
	if er != nil {
		entities.PrintError(er.Error())
		entities.PrintLog(er.Error())
		httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
	} else {
		response := helper.ApiResponse("Reset apex password succeeded", "success", user)
		httpio.Response(http.StatusOK, response)
	}

}
