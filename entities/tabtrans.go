package entities

import "time"

type Tabtrans struct {
	TabtransID       int
	Tgl_trans        time.Time
	No_rekening      string
	Kode_trans       string
	My_kode_trans    string
	Pokok            float64
	Kuitansi         string
	Userid           int
	Keterangan       string
	Verifikasi       string
	Tob              string
	Sandi_trans      string
	Posted_to_gl     string
	Kode_kantor      string
	Jam              string
	Pay_lkm_source   string
	Pay_lkm_norek    string
	Pay_idpel        string
	Pay_biller_code  string
	Pay_product_code string
}
