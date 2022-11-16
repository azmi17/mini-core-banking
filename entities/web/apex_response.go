package web

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

type TabSCGroup struct {
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
}

type ResetApexPwdResponse struct {
	KodeLkm       string `json:"kode_lkm"`
	Password_Smec string `json:"password_smec"`
}

type LoginData struct {
	User_Id         int    `json:"user_id"`
	User_Name       string `json:"user_name"`
	Nama_lengkap    string `json:"nama_lengkap"`
	Tanggal_Expried string `json:"tgl_expired"`
}

type LoginResponse struct {
	Response_Code string     `json:"response_code"`
	Response_Msg  string     `json:"response_message"`
	Data          *LoginData `json:"data,omitempty"`
}
