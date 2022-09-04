package handler

import (
	"apex-ems-integration-clean-arch/delivery/handler/httpio"
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/statuscode"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/helper"
	"apex-ems-integration-clean-arch/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteLKM(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	payload := web.KodeLKMFilter{}
	httpio.Bind(&payload)

	usecase := usecase.NewApexUsecase()
	er := usecase.DeleteLkm(payload.KodeLkm)
	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "Record not found!", nil)
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		response := helper.ApiResponse("Institution has been deleted", "success", nil)
		httpio.Response(http.StatusOK, response)
	}

}
