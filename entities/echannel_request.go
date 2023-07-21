package entities

type TransHistoryRequest struct {
	TglAwal  string `form:"tgl_awal" binding:"required"`
	TglAkhir string `form:"tgl_akhir" binding:"required"`
	Filter   string `form:"filter"`
}
