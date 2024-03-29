package entities

type GlobalResponse struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

type LKMCreateResponse struct {
	KodeLkm        string `json:"kode_lkm"`
	NamaLembaga    string `json:"nama_lembaga"`
	Alamat         string `json:"alamat"`
	Telepon        string `json:"telpon"`
	KodeSC         string `json:"kode_sc"`
	NoRekening     string `json:"apex_norek"`
	Saldo          int    `json:"saldo_akhir"`
	Plafond        int    `json:"plafond"`
	SetoranMinimum int    `json:"setoran_minimum"`
	Status         int    `json:"status"`
	IsSaldoChecked int    `json:"is_saldo_checked"`
	UserNameSmec   string `json:"user_name_smec"`
	PwdSmec        string `json:"password_smec"`
	User_Id        int    `json:"user_id"`
}

type LKMUpdateResponse struct {
	KodeLkm        string  `json:"kode_lkm"`
	NamaLembaga    string  `json:"nama_lembaga"`
	KodeSC         string  `json:"kode_sc"`
	Alamat         string  `json:"alamat"`
	Telepon        string  `json:"telpon"`
	NoRekening     string  `json:"apex_norek"`
	Plafond        float64 `json:"plafond"`
	SetoranMinimum float64 `json:"setoran_minimum"`
	Status         int     `json:"status"`
	IsSaldoChecked int     `json:"is_saldo_checked"`
	UserID         int     `json:"user_id"`
}

type TabSCGroup struct {
	KodeGroup      string `json:"kode_group"`
	DeskripsiGroup string `json:"deskripsi_group"`
}

type GetDetailLKMInfo struct {
	// KodeLembaga string  `json:"kode_lkm"`
	// NamaLembaga string  `json:"nama_lembaga"`
	// Vendor      string  `json:"vendor"`
	// Alamat      string  `json:"alamat"`
	// Kontak      string  `json:"kontak"`
	// NoRekening  string  `json:"apex_norek"`
	// Saldo       float64 `json:"saldo_akhir"`
	// Plafond     float64 `json:"plafond"`
	// StatusTab   string  `json:"status_tab"`

	KodeLkm        string  `json:"kode_lkm"`
	NamaLembaga    string  `json:"nama_lembaga"`
	Vendor         string  `json:"vendor"`
	Alamat         string  `json:"alamat"`
	Telepon        string  `json:"telpon"`
	NoRekening     string  `json:"apex_norek"`
	Saldo          float64 `json:"saldo_akhir"`
	Plafond        float64 `json:"plafond"`
	SetoranMinimum float64 `json:"setoran_minimum"`
	Status         int     `json:"status"`
	IsSaldoChecked int     `json:"is_saldo_checked"`
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

type RoutingRekIndukData struct {
	KodeLkm    string `json:"kode_lkm"`
	NorekInduk string `json:"norek_induk"`
}

type SaveRoutingRekIndukResponse struct {
	Response_Code string               `json:"response_code"`
	Response_Msg  string               `json:"response_message"`
	Data          *RoutingRekIndukData `json:"data,omitempty"`
}

type CreateManajemenUserDataResponse struct {
	User_ID      int    `json:"user_id"`
	User_Name    string `json:"user_name"`
	Nama_Lengkap string `json:"nama_lengkap"`
	Password     string `json:"password"`
	Jabatan      string `json:"jabatan"`
	Unit_Kerja   string `json:"unit_kerja"`
	Tgl_Expired  string `json:"tgl_expired"`
	StatusAktif  int    `json:"status_aktif"`
	User_Code    string `json:"user_code"`
}

type UpdateManajemenUserDataResponse struct {
	User_ID      int    `json:"user_id"`
	User_Name    string `json:"user_name"`
	Nama_Lengkap string `json:"nama_lengkap"`
	Jabatan      string `json:"jabatan"`
	Unit_Kerja   string `json:"unit_kerja"`
	Tgl_Expired  string `json:"tgl_expired"`
	StatusAktif  int    `json:"status_aktif"`
	User_Code    string `json:"user_code"`
}

type ManajemenUserDataResponse struct {
	User_ID      int    `json:"user_id"`
	User_Name    string `json:"user_name"`
	Nama_Lengkap string `json:"nama_lengkap"`
	Jabatan      string `json:"jabatan"`
	Unit_Kerja   string `json:"unit_kerja"`
	Tgl_Expired  string `json:"tgl_expired"`
	StatusAktif  int    `json:"status_aktif"`
	User_Code    string `json:"user_code"`
}

type GetListTabtransTrx struct {
	TabtransID  int     `json:"tabtrans_id"`
	TglTrans    string  `json:"tgl_trans"`
	KodeLKM     string  `json:"kode_lkm"`
	NamaLembaga string  `json:"nama_lembaga"`
	Pokok       float64 `json:"pokok"`
	Dk          string  `json:"dk"`
	Lkm_Norek   string  `json:"lkm_no_rekening"`
	Idpel       string  `json:"idpel"`
	KodeTrans   string  `json:"kode_trans"`
	Kuitansi    string  `json:"kuitansi"`
	Keterangan  string  `json:"keterangan"`
	BillerCode  string  `json:"biller_code"`
	ProductCode string  `json:"product_code"`
	UserID      int     `json:"user_id"`
}

type GetCountWithSumTabtrans struct {
	TotalTrx   int     `json:"total_trx"`
	TotalPokok float64 `json:"total_pokok"`
}

type GetListTabtransInfoWithCountSumResp struct {
	TotalTrx   int                   `json:"total_trx"`
	TotalPokok float64               `json:"total_pokok"`
	Data       *[]GetListTabtransTrx `json:"data,omitempty"`
}

type RepostingData struct {
	KodeLKM     string
	TotalDebet  float64
	TotalKredit float64
}

type CalculateSaldoResult struct {
	KodeLKM    string
	SaldoAkhir float64
}

type LKMlist struct {
	KodeLKM string
}

type GetListsDepositTrxRes struct {
	TransID       int     `json:"trans_id"`
	TglTrans      string  `json:"tgl_trans"`
	KodeLKM       string  `json:"kode_lkm"`
	NamaLembaga   string  `json:"nama_lembaga"`
	Tunai         float64 `json:"tunai"`
	Transfer      float64 `json:"transfer"`
	Piutang       float64 `json:"piutang"`
	ViaBrhancless float64 `json:"via_branchless"`
	Total         float64 `json:"total"`
	NoBukti       string  `json:"no_bukti"`
	Keterangan    string  `json:"keterangan"`
	KodeTrans     string  `json:"kode_trans"`
	DK            string  `json:"d/k"`
	UserID        int     `json:"user_id"`
}

type VirtualAccountResponse struct {
	NomorVA     string  `json:"no_va"`
	NasabahID   string  `json:"nasabah_id"`
	NamaLembaga string  `json:"nama_lembaga"`
	Vendor      string  `json:"vendor_code"`
	SLAType     string  `json:"kode_pembayaran_sla"`
	Alamat      string  `json:"alamat"`
	Telepon     string  `json:"telepon"`
	Saldo       float64 `json:"saldo_akhir"`
	Status      int     `json:"status"`
}

type VirtualAccountTransResponse struct {
	TabtransID  int     `json:"tabtrans_id"`
	NamaLembaga string  `json:"nama_lembaga"`
	NoVA        string  `json:"no_va"`
	SLAType     string  `json:"type_sla"`
	TglTrans    string  `json:"tgl_trans"`
	Kuitansi    string  `json:"kuitansi"`
	Keterangan  string  `json:"keterangan"`
	Nominal     float64 `json:"nominal"`
}
