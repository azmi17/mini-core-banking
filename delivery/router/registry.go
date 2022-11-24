/*
 * Copyright (c) 2022 Randy Ardiansyah https://github.com/randyardiansyah25/<repo>
 *
 * Created Date: Wednesday, 16/03/2022, 10:32:08
 * Author: Randy Ardiansyah
 *
 * Filename: /home/Documents/workspace/go/src/router-template/delivery/router/registry.go
 * Project : /home/Documents/workspace/go/src/router-template/delivery/router
 *
 * HISTORY:
 * Date                  	By                 	Comments
 * ----------------------	-------------------	--------------------------------------------------------------------------------------------------------------------
 */

package router

import (
	appInformationsHandler "apex-ems-integration-clean-arch/delivery/handler/app-informations"
	institutionsHandler "apex-ems-integration-clean-arch/delivery/handler/institutions"
	institutionRoutingsHandler "apex-ems-integration-clean-arch/delivery/handler/lkm-routings"
	tabungansHandler "apex-ems-integration-clean-arch/delivery/handler/tabungans"
	usersHandler "apex-ems-integration-clean-arch/delivery/handler/users"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(router *gin.Engine) {

	// API Versioning:
	apiv1 := router.Group("api/v1")

	// API Endpoint:
	apiv1.GET("/version", appInformationsHandler.AppInfo)
	apiv1.GET("/vendors", tabungansHandler.GetTabScGroup)

	apiv1.GET("/institutions/:limit/:offset", tabungansHandler.GetTabungansLkmInfoList)
	apiv1.GET("/institution/:kode_lkm", tabungansHandler.GetTabunganLkmInfo)
	apiv1.POST("/institution", institutionsHandler.CreateLKM)
	apiv1.PUT("/institution", institutionsHandler.UpdateLKM)
	apiv1.DELETE("/institution/:kode_lkm", institutionsHandler.DeleteLKM)

	apiv1.POST("/user", usersHandler.CreateSysUser)
	apiv1.PUT("/user", usersHandler.UpdateSysUser)
	apiv1.POST("/user/search", usersHandler.FindSingleUserByUserName)
	apiv1.GET("/user/:kode_lkm", usersHandler.GetSingleUserByUserName)
	apiv1.GET("/users/:limit/:offset", usersHandler.GetListOfUsers)
	apiv1.POST("/user/login", usersHandler.LoginUser)
	apiv1.PUT("/user/reset-password", usersHandler.ResetApexPassword)

	apiv1.GET("/routing/:kode_lkm", institutionRoutingsHandler.GetRoutingRekInduk)
	apiv1.GET("/routings/:limit/:offset", institutionRoutingsHandler.GetListRoutingRekInduk)
	apiv1.POST("/routing", institutionRoutingsHandler.CreateRoutingRekInduk)
	apiv1.PUT("/routing", institutionRoutingsHandler.UpdateRoutingRekInduk)
	apiv1.DELETE("/routing/:kode_lkm", institutionRoutingsHandler.DeleteRoutingRekIndukByKodeLKM)

	apiv1.DELETE("/flush", institutionsHandler.HardDeleteLKM)
}
