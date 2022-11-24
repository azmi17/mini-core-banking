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

func CreateRoutingRekInduk(ctx *gin.Context) {

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding form
	payload := web.CreateRoutingRekInduk{}
	rerr := httpio.BindWithErr(&payload)
	if rerr != nil {
		errors := helper.FormatValidationError(rerr)
		errorMesage := gin.H{"errors": errors}
		response := helper.ApiResponse("Add routing rek induk failed", http.StatusUnprocessableEntity, "failed", errorMesage)
		httpio.Response(http.StatusUnprocessableEntity, response)
		return
	}

	usecase := usecase.NewRoutingIndukUsecase()
	routingData, er := usecase.CreateSysApexRoutingRekInduk(payload)

	if er != nil {
		entities.PrintError(er.Error())
		httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		return

	} else {
		httpio.Response(http.StatusOK, routingData)
	}

}
