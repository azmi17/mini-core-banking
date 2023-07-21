package alterapex

import (
	"net/http"
	"new-apex-api/delivery/handler/httpio"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/usecase"

	"github.com/gin-gonic/gin"
)

func AlterNasabahIDOnTabungAndNasabah(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)
	httpio.Recv()

	usecase := usecase.NewMigrasiApexUsecase()
	er := usecase.ReplaceNasabahIDOnNasabahWithNorek()

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
		resp.ResponseMessage = "Alter nasabah_id succeeded"
	}

	httpio.Response(http.StatusOK, resp)
}
