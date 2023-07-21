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

func UpdateRoutingRekInduk(ctx *gin.Context) {

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding form
	payload := entities.UpdateRoutingRekInduk{}
	rerr := httpio.BindWithErr(&payload)
	if rerr != nil {
		errors := helper.FormatValidationError(rerr)
		errorMesage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update routing rek induk failed", http.StatusUnprocessableEntity, "failed", errorMesage)
		httpio.Response(http.StatusUnprocessableEntity, response)
		return
	}

	usecase := usecase.NewRoutingIndukUsecase()
	routingData, er := usecase.UpdateSysApexRoutingRekInduk(payload)

	if er != nil {
		if er == err.NoRecord {
			entities.PrintLog(er.Error())
			httpio.ResponseString(statuscode.StatusNoRecord, "Record not found!", nil)
			return
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
			return
		}

	} else {
		httpio.Response(http.StatusOK, routingData)
	}

}
