package web

type KodeLKMFilter struct {
	KodeLkm string `form:"kode_lkm"`
}

type KodeLKMUri struct {
	KodeLkm string `uri:"kode_lkm"`
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
