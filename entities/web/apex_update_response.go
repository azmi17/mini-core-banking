package web

type LKMUpdateResponse struct {
	KodeLkm      string `json:"kode_lkm"`
	Nama_Lembaga string `json:"nama_lembaga"`
	Alamat       string `json:"alamat"`
	Telepon      string `json:"telpon"`
	User_Id      int    `json:"user_id"`
}
