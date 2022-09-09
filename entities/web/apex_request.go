package web

type KodeLKMFilter struct {
	KodeLkm string `form:"kode_lkm"`
}

type KodeLKMUri struct {
	UserName string `uri:"user_name"`
}

type LimitOffsetLkmUri struct {
	Limit  int `uri:"limit"`
	Offset int `uri:"offset"`
}

type SaveLKMApex struct {
	KodeLkm      string `form:"kode_lkm" binding:"required,min=4"`
	Nama_Lembaga string `form:"nama_lembaga" binding:"required"`
	Alamat       string `form:"alamat" binding:"required"`
	Telepon      string `form:"telepon" binding:"required"`
	User_Id      int    `form:"user_id" binding:"required"`
}

type LKMCreateResponse struct {
	KodeLkm        string `json:"kode_lkm"`
	Nama_Lembaga   string `json:"nama_lembaga"`
	Alamat         string `json:"alamat"`
	Telepon        string `json:"telpon"`
	No_rekening    string `json:"no_rekening"`
	Saldo_Akhir    int    `json:"saldo_akhir"`
	User_Name_Smec string `json:"user_name_smec"`
	Password_Smec  string `json:"password_smec"`
	User_Id        int    `json:"user_id"`
}

type LKMUpdateResponse struct {
	KodeLkm      string `json:"kode_lkm"`
	Nama_Lembaga string `json:"nama_lembaga"`
	Alamat       string `json:"alamat"`
	Telepon      string `json:"telpon"`
	User_Id      int    `json:"user_id"`
}

type SCGroup struct {
	KodeGroup      string `json:"kode_group"`
	DeskripsiGroup string `json:"deskripsi_group"`
}

type GetDetailLKMInfo struct {
	KodeLembaga string  `json:"kode_lkm"`
	NamaLembaga string  `json:"nama_lembaga"`
	Vendor      string  `json:"vendor"`
	Alamat      string  `json:"alamat"`
	Kontak      string  `json:"kontak"`
	NoRekening  string  `json:"apex_norek"`
	Saldo       float64 `json:"saldo_akhir"`
	Plafond     float64 `json:"plafond"`
	StatusTab   string  `json:"status_tab"`
	// omitempty

}

type ResetApexPwdResponse struct {
	KodeLkm       string `json:"kode_lkm"`
	Password_Smec string `json:"password_smec"`
}
