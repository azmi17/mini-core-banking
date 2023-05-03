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
	appInformationsHandler "new-apex-api/delivery/handler/app-informations"
	institutionsHandler "new-apex-api/delivery/handler/institutions"
	institutionRoutingsHandler "new-apex-api/delivery/handler/lkm-routings"
	referensiHandler "new-apex-api/delivery/handler/referensi-apex"
	repostingHandler "new-apex-api/delivery/handler/repostings"
	tabtransHandler "new-apex-api/delivery/handler/tabtrans"
	tabungansHandler "new-apex-api/delivery/handler/tabungans"
	usersHandler "new-apex-api/delivery/handler/users"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(router *gin.Engine) {

	// API Versioning:
	apiv1 := router.Group("api/v1")

	// API Endpoint:
	apiv1.GET("/version", appInformationsHandler.AppInfo)
	// apiv1.GET("/vendors", tabungansHandler.GetTabScGroup)

	// a:= apiv1.Group("/", validator)

	apiv1.GET("/institutions/:limit/:offset", tabungansHandler.GetTabungansLkmInfoList)
	apiv1.GET("/institution/:kode_lkm", tabungansHandler.GetTabunganLkmInfo)
	apiv1.POST("/institution", institutionsHandler.CreateLKM)
	apiv1.PUT("/institution", institutionsHandler.UpdateLKM)
	apiv1.DELETE("/institution/:kode_lkm", institutionsHandler.DeleteLKM) // => Jangan Embedd ke EMS (proses di lakukan di Apex)

	apiv1.POST("/repostings", repostingHandler.RepostingSaldoByBulk)
	apiv1.POST("/repostings/all", repostingHandler.RepostingAllByApi)

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

	apiv1.POST("/tabtrans/:limit/:offset", tabtransHandler.GetListsTabtransTrx)
	apiv1.POST("/tabtrans/by-stan", tabtransHandler.GetListsTabtransTrxBySTAN)
	apiv1.PUT("/tabtrans", tabtransHandler.ChangeDateOnTabtransTrx)
	apiv1.DELETE("/tabtrans/:tabtrans_id", tabtransHandler.DeleteTabtransTrx)
	apiv1.POST("/tabtrans/laporan/rekening_koran", tabtransHandler.GetRekeningKoranLKMDetail)
	apiv1.POST("/tabtrans/laporan/nominatif_deposit/:limit/:offset", tabtransHandler.GetNominatifDeposit)
	apiv1.POST("/tabtrans/laporan/daftar_transaksi/:limit/:offset", tabtransHandler.GetLaporanTransaksi)
	apiv1.POST("/tabtrans/deposits", tabtransHandler.GetListsDepositHisotry)

	apiv1.GET("/referensi/vendors", referensiHandler.GetListsScGroup)
	apiv1.GET("/referensi/banks", referensiHandler.GetListsBankGroup)
	apiv1.GET("/referensi/jenis_deposit", referensiHandler.GetListsJenisTransaksiDeposit)
	apiv1.GET("/referensi/jenis_tabungan", referensiHandler.GetListsJenisTransaksiTabungan)
	apiv1.GET("/referensi/tabproducts", referensiHandler.GetlistsProdukTabungan)

	apiv1.DELETE("/flush", institutionsHandler.HardDeleteLKM)
}
