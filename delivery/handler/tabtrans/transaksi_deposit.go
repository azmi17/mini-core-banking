package handler

import (
	"net/http"
	"new-apex-api/delivery/handler/httpio"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/helper"
	"new-apex-api/usecase"

	"github.com/gin-gonic/gin"
)

func TransaksiDeposit(ctx *gin.Context) {
	httpio := httpio.NewRequestIO(ctx)

	payload := entities.DepositRequest{}
	rerr := httpio.BindWithErr(&payload)
	if rerr != nil {
		errors := helper.FormatValidationError(rerr)
		errorMesage := gin.H{"errors": errors}
		response := helper.ApiResponse("Deposit transaction failed", http.StatusUnprocessableEntity, "failed", errorMesage)
		httpio.Response(http.StatusUnprocessableEntity, response)
		return
	}
	httpio.Bind(&payload)

	usecase := usecase.NewTabtransUsecase()
	er := usecase.TransaksiDeposit(payload)

	resp := entities.GlobalResponse{}
	if er != nil {
		if er == err.NoRecord {
			resp.ResponseCode = "1111"
			resp.ResponseMessage = er.Error()
		} else if er == err.RekeningBelumAktif {
			resp.ResponseCode = "1111"
			resp.ResponseMessage = er.Error()
		} else if er == err.RekeningNonAktif {
			resp.ResponseCode = "1111"
			resp.ResponseMessage = er.Error()
		} else if er == err.RekeningDitutup {
			resp.ResponseCode = "1111"
			resp.ResponseMessage = er.Error()
		} else if er == err.RekeningDiBlokir {
			resp.ResponseCode = "1111"
			resp.ResponseMessage = er.Error()
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
			return
		}
	} else {
		resp.ResponseCode = "0000"
		resp.ResponseMessage = "Deposit transaciton succeeded"
	}

	httpio.Response(http.StatusOK, resp)
}
