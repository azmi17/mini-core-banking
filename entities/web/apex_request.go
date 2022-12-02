package web

type StanFilter struct {
	Stan string `form:"stan"`
}

type KodeLKMFilter struct {
	KodeLkm string `form:"kode_lkm" binding:"required,min=4"`
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

type SaveLKMApex struct {
	KodeLkm      string `form:"kode_lkm" binding:"required,min=4"`
	Nama_Lembaga string `form:"nama_lembaga" binding:"required"`
	Alamat       string `form:"alamat" binding:"required"`
	Telepon      string `form:"telepon" binding:"required"`
	User_Id      int    `form:"user_id" binding:"required"`
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

type GetListTabtransByDate struct {
	TanggalAwal  string `form:"tanggal_awal" binding:"required"`
	TanggalAkhir string `form:"tanggal_akhir" binding:"required"`
	BankCode     string `form:"kode_lkm"`
}

type ChangeTglTransOnTabtrans struct {
	TabtransID int    `form:"tabtrans_id" binding:"required"`
	Tanggal    string `form:"tanggal" binding:"required,max=8"`
}

type RepostingTabungPayload struct {
	KodeLKM    string
	SaldoAkhir float64
}
