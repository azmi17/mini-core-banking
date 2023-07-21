package handler

import (
	"net/http"
	"new-apex-api/delivery/handler/httpio"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/usecase"

	"github.com/gin-gonic/gin"
)

func RepostingAllLKM(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)
	httpio.Recv()

	usecase := usecase.NewRepostingUsecase()
	er := usecase.RepostingSaldoByScheduler()

	resp := entities.GlobalResponse{}
	if er != nil {
		if er == err.NoRecord {
			resp.ResponseCode = "1111"
			resp.ResponseMessage = er.Error()
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
			return
		}
	} else {
		resp.ResponseCode = "0000"
		resp.ResponseMessage = "Reposting saldo succeeded"
	}

	httpio.Response(http.StatusOK, resp)
}
