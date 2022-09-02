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

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding from
	payload := web.SaveApex{}
	er := httpio.Bind(&payload)
	if er != nil {
		// Errors validations
		errors := helper.FormatValidationError(er)
		errorMesage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create institution failed", "failed", errorMesage)
		httpio.Response(http.StatusUnprocessableEntity, response)
		return // <= return-kan jika ada validasi detect error supaya tidak di eksekusi ke bawah..
	}
	uscase := usecase.NewApexUsecase()
	lkm, er := uscase.SaveLkm(payload)
	if er != nil {
		if er == err.DuplicateEntry {
			httpio.ResponseString(statuscode.StatusDuplicate, "LKM data is available!", nil)
		} else {
			entities.PrintError(er.Error())
			entities.PrintLog(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		response := helper.ApiResponse("New lkm has been created", "success", lkm)
		httpio.Response(http.StatusOK, response)
	}

}
