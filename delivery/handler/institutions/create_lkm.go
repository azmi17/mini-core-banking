package handler

import (
	"net/http"
	"new-apex-api/delivery/handler/httpio"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/entities/statuscode"
	"new-apex-api/entities/web"
	"new-apex-api/helper"
	"new-apex-api/usecase"

	"github.com/gin-gonic/gin"
)

func CreateLKM(ctx *gin.Context) {

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding form
	payload := web.SaveLKMApex{}
	rerr := httpio.BindWithErr(&payload)
	if rerr != nil {
		errors := helper.FormatValidationError(rerr)
		errorMesage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register institution failed", http.StatusUnprocessableEntity, "failed", errorMesage)
		httpio.Response(http.StatusUnprocessableEntity, response)
		return
	}

	usecase := usecase.NewLkmUsecase()
	lkm, er := usecase.CreateLkm(payload)
	if er != nil {
		if er == err.DuplicateEntry {
			httpio.ResponseString(statuscode.StatusDuplicate, "Institution data is available!", nil)
		} else {
			entities.PrintError(er.Error())
			entities.PrintLog(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		httpio.Response(http.StatusOK, lkm)
	}

}
