package handler

import (
	"net/http"
	"new-apex-api/delivery/handler/httpio"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/entities/statuscode"
	"new-apex-api/helper"
	"new-apex-api/usecase"

	"github.com/gin-gonic/gin"
)

func PermanentlyDeleteApexTransaction(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	payload := entities.MultipleTabtransID{}
	rerr := httpio.BindWithErr(&payload)
	if rerr != nil {
		errors := helper.FormatValidationError(rerr)
		errorMesage := gin.H{"errors": errors}
		response := helper.ApiResponse("delete tabtrans transaction failed", http.StatusUnprocessableEntity, "failed", errorMesage)
		httpio.Response(http.StatusUnprocessableEntity, response)
		return
	}
	httpio.Bind(&payload)

	usecase := usecase.NewTabtransUsecase()
	er := usecase.HardDeleteApexTransaction(payload.ListOfTabtransID)
	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "Record not found", nil)
		} else {
			entities.PrintError(er.Error())
			entities.PrintLog(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		httpio.Response(http.StatusOK, nil)
	}

}
