package entities

type StanFilter struct {
	Stan string `form:"stan"`
}

type KodeLKMFilter struct {
	KodeLkm string `form:"kode_lkm" binding:"required"`
}

type GlobalFilter struct {
	Filter string `form:"filter"`
}

type MultipleKodeLKM struct {
	ListOfKodeLKM []string `form:"kode_lkm[]" binding:"required"`
}

type MultipleVANumberFilter struct {
	VaNumber []string `form:"no_va[]" binding:"required"`
}

type MultipleChangeDateTransaction struct {
	ListOfTabtransID []int  `form:"tabtrans_id[]" binding:"required"`
	Tanggal          string `form:"tanggal" binding:"required,max=8"`
}

type MultipleTabtransID struct {
	ListOfTabtransID []int `form:"tabtrans_id[]" binding:"required"`
}

type KodeLKMUri struct {
	KodeLkm string `uri:"kode_lkm"`
}

type TabtransIDUri struct {
	TabtransID int `uri:"tabtrans_id"`
}

type LimitOffsetLkmUri struct {
	Limit  int `uri:"limit"`
	Offset int `uri:"offset"`
}

type CreateLKM struct {
	KodeLkm        string `form:"kode_lkm" binding:"required,min=4"`
	KodeSC         string `form:"kode_sc" binding:"required"`
	NamaLembaga    string `form:"nama_lembaga" binding:"required"`
	Alamat         string `form:"alamat" binding:"required"`
	Telepon        string `form:"telepon" binding:"required"`
	Plafond        int    `form:"plafond"`
	IsSaloChecked  int    `form:"is_saldo_checked"`
	SetoranMinimum int    `form:"setoran_minimum"`
	UserID         int    `form:"user_id" binding:"required"`
}

type UpdateLKM struct {
	KodeLkm        string  `form:"kode_lkm" binding:"required,min=4"`
	NamaLembaga    string  `form:"nama_lembaga" binding:"required"`
	KodeSC         string  `form:"kode_sc" binding:"required"`
	Alamat         string  `form:"alamat" binding:"required"`
	Telepon        string  `form:"telepon" binding:"required"`
	Plafond        float64 `form:"plafond"`
	SetoranMinimum float64 `form:"setoran_minimum"`
	Status         int     `form:"status"`
	IsSaldoChecked int     `form:"is_saldo_checked"`
	UserID         int     `form:"user_id" binding:"required"`
}

type LoginInput struct {
	User_Name string `form:"user_name" binding:"required"`
	Password  string `form:"password" binding:"required"`
}

type CreateRoutingRekInduk struct {
	KodeLkm    string `form:"kode_lkm" binding:"required"`
	NorekInduk string `form:"norek_induk" binding:"required"`
}

type UpdateRoutingRekInduk struct {
	KodeLkm       string `form:"kode_lkm" binding:"required"`
	NorekInduk    string `form:"norek_induk" binding:"required"`
	KodeLkmTarget string `form:"kode_lkm_target" binding:"required"`
}

type CreateManajemenUser struct {
	UserName string `form:"user_name" binding:"required,min=4"`
	NamaUser string `form:"nama_user" binding:"required"`
	Jabatan  string `form:"jabatan" binding:"required"`
	UserCode string `form:"user_code" binding:"required"`
}

type UpdateManajemenUser struct {
	UserName    string `form:"user_name" binding:"required,min=4"`
	NamaUser    string `form:"nama_user" binding:"required"`
	Jabatan     string `form:"jabatan" binding:"required"`
	StatusAktif int    `form:"status_aktif" binding:"required"`
	UserCode    string `form:"user_code" binding:"required"`
}

type GetListTabtrans struct {
	TanggalAwal  string `form:"tanggal_awal" binding:"required"`
	TanggalAkhir string `form:"tanggal_akhir" binding:"required"`
	Filter       string `form:"filter"` // norek atau kuitansi
}

type ChangeTglTransOnTabtrans struct {
	TabtransID int    `form:"tabtrans_id" binding:"required"`
	Tanggal    string `form:"tanggal" binding:"required,max=8"`
}

type GetListsDepositTrxReq struct {
	TanggalAkhir string `form:"tanggal_awal" binding:"required"`
	TanggalAwal  string `form:"tanggal_akhir" binding:"required"`
	// BankCode     string `form:"kode_lkm"`
}

type UpdateRekeningLKM struct {
	KodeLKM        string `form:"kode_lkm"`
	KodeSC         string
	Status         int     `form:"status"`
	Plafond        float64 `form:"plafond"`
	SetoranMinimum float64 `form:"setoran_minimum"`
	IsSaldoChecked int
	UserID         int `form:"user_id"`
}

type CreateVirtualAccount struct {
	NamaLembaga       string `form:"nama_lembaga" binding:"required"`
	VendorCode        string `form:"vendor_code" binding:"required"`
	Alamat            string `form:"alamat" binding:"required"`
	Telepon           string `form:"telepon" binding:"required"`
	KodePembayaranSLA string `form:"kode_pembayaran_sla" binding:"required"`
	Status            int    `form:"status" binding:"required"`
	UserID            int    `form:"user_id" binding:"required"`
}

type UpdateVirtualAccount struct {
	NasabahID         string `form:"nasabah_id" binding:"required"`
	NamaLembaga       string `form:"nama_lembaga" binding:"required"`
	VendorCode        string `form:"vendor_code" binding:"required"`
	Alamat            string `form:"alamat" binding:"required"`
	Telepon           string `form:"telepon" binding:"required"`
	KodePembayaranSLA string `form:"kode_pembayaran_sla" binding:"required"`
	Status            int    `form:"status" binding:"required"`
	UserID            int    `form:"user_id" binding:"required"`
}

// Nitip
type NorekWithNID struct {
	NoRekening string
	NasabahID  string
}
