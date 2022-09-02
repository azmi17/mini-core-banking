package web

type SaveApex struct {
	KodeLkm      string `form:"kode_lkm" binding:"required,min=4"`
	Nama_Lembaga string `form:"nama_lembaga" binding:"required"`
	Alamat       string `form:"alamat" binding:"required"`
	Telepon      string `form:"telepon" binding:"required"`
	User_Id      int    `form:"user_id" binding:"required"`
}
