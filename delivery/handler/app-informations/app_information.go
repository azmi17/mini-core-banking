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
		"App Name":         helper.AppName,
		"App Description":  helper.AppDescription,
		"App Version":      helper.AppVersion,
		"App Latest Build": helper.LastBuild,
	}

	httpio.Response(http.StatusOK, appInfo)
}
