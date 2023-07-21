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

func GetListsTabtransTransaction(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	uriPayload := entities.LimitOffsetLkmUri{}
	httpio.BindUri(&uriPayload)

	formPayload := entities.GetListTabtrans{}
	rerr := httpio.BindWithErr(&formPayload)
	if rerr != nil {
		errors := helper.FormatValidationError(rerr)
		errorMesage := gin.H{"errors": errors}
		response := helper.ApiResponse("get tabtrans transaction failed", http.StatusUnprocessableEntity, "failed", errorMesage)
		httpio.Response(http.StatusUnprocessableEntity, response)
		return
	}

	usecase := usecase.NewTabtransUsecase()
	listTrx, er := usecase.GetListsTabtransTransaction(formPayload, uriPayload)
	// resp := entities.GetListTabtransInfoWithCountSumResp{}
	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "record not found", nil)
			return
		} else if er == err.BadRequest {
			httpio.ResponseString(http.StatusBadRequest, "invalid parameters", nil)
			return
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
			return
		}
	} else {
		// resp.TotalTrx = total.TotalTrx
		// resp.TotalPokok = total.TotalPokok
		// resp.Data = &listTrx
		httpio.Response(http.StatusOK, listTrx)
	}
	// httpio.Response(http.StatusOK, resp)
}
