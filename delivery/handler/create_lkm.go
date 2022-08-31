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

func CreateLKM(ctx *gin.Context) {
	httpio := httpio.NewRequestIO(ctx) // catch request

	payload := web.PayloadApex{}
	httpio.Bind(&payload)

	uscase := usecase.NewApexUsecase()
	lkm, er := uscase.CreateLkm(payload)
	if er != nil {
		if er == err.DuplicateEntry {
			httpio.ResponseString(statuscode.StatusDuplicate, "LKM data is available!", nil)
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		response := helper.ApiResponse("New lkm has been created", "success", lkm)
		httpio.Response(http.StatusOK, response)
	}
}
