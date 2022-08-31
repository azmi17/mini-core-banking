package web

type PayloadApex struct {
	KodeLkm      string `form:"kode_lkm"`
	Nama_Lembaga string `form:"nama_lembaga"`
	Alamat       string `form:"alamat"`
	Telepon      string `form:"telepon"`
	User_Id      int    `form:"user_id"`
}
