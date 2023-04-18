package handler

import (
	"net/http"
	"new-apex-api/delivery/handler/httpio"
	"new-apex-api/helper"

	"github.com/gin-gonic/gin"
)

func AppInfo(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)
	httpio.Recv()

	appInfo := map[string]interface{}{
		"App Name":        helper.AppAuthor,
		"App Description": helper.AppDescription,
		"App Version":     helper.AppVersion,
		"App Author":      helper.AppAuthor,
	}

	httpio.Response(http.StatusOK, appInfo)
}
