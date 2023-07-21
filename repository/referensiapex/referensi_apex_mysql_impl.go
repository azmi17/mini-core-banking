package referensiapex

import (
	"database/sql"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"os"
	"strings"
)

func newReferensiApexMysqlImpl(apexConn *sql.DB) ReferensiApexRepo {
	return &referensiApexMysqlImpl{
		apexDb: apexConn,
	}
}

type referensiApexMysqlImpl struct {
	apexDb *sql.DB
}

func (r *referensiApexMysqlImpl) GetListsScGroup() (lists []entities.ScGroup, er error) {
	rows, er := r.apexDb.Query("SELECT kode_group2, deskripsi_group2 FROM tab_kode_group2")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var scGroup entities.ScGroup
		if er = rows.Scan(&scGroup.KodeGroup, &scGroup.DeskripsiGroup); er != nil {
			return lists, er
		}

		lists = append(lists, scGroup)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (r *referensiApexMysqlImpl) GetListsJenisTransaksiTabungan() (lists []entities.JenisTransaksi, er error) {
	rows, er := r.apexDb.Query("SELECT kode_trans, deskripsi_trans FROM tab_kode_trans")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var tabKodeTrans entities.JenisTransaksi
		if er = rows.Scan(&tabKodeTrans.KodeTrans, &tabKodeTrans.DeskripsiTrans); er != nil {
			return lists, er
		}

		lists = append(lists, tabKodeTrans)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (r *referensiApexMysqlImpl) GetListsJenisTransaksiDeposit() (lists []entities.JenisTransaksi, er error) {

	item := strings.Split(os.Getenv("app.jenis_transaksi_deposit"), ",")
	kodeTrans := (strings.Join(item, ","))

	rows, er := r.apexDb.Query(`SELECT kode_trans, deskripsi_trans FROM tab_kode_trans WHERE kode_trans IN (` + kodeTrans + `) ORDER BY kode_trans ASC`)
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var tabKodeTrans entities.JenisTransaksi
		if er = rows.Scan(&tabKodeTrans.KodeTrans, &tabKodeTrans.DeskripsiTrans); er != nil {
			return lists, er
		}

		lists = append(lists, tabKodeTrans)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (r *referensiApexMysqlImpl) GetListsBankGroup() (lists []entities.BankGroup, er error) {
	rows, er := r.apexDb.Query("SELECT id, nama_bank, no_rekening, deskripsi FROM deposit_bank_info")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var banks entities.BankGroup
		if er = rows.Scan(&banks.Id, &banks.NamaBank, &banks.NoRekeningBank, &banks.Deskripsi); er != nil {
			return lists, er
		}

		lists = append(lists, banks)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (r *referensiApexMysqlImpl) GetlistsProdukTabungan() (lists []entities.ProdukTabungan, er error) {
	rows, er := r.apexDb.Query("SELECT kode_produk, deskripsi_produk FROM tab_produk")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var tabProduk entities.ProdukTabungan
		if er = rows.Scan(&tabProduk.KodeProduk, &tabProduk.DeskripsiProduk); er != nil {
			return lists, er
		}

		lists = append(lists, tabProduk)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (r *referensiApexMysqlImpl) GetListsJenisPembayaranSLA() (lists []entities.JenisPembayaranSLA, er error) {

	item := strings.Split(os.Getenv("app.list_kode_tipe_sla"), ",")
	kodeProdukSLA := (strings.Join(item, ","))

	rows, er := r.apexDb.Query(`SELECT kode_produk, deskripsi_produk FROM tab_produk WHERE kode_produk IN (` + kodeProdukSLA + `) ORDER BY kode_produk ASC`)
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var produkSLA entities.JenisPembayaranSLA
		if er = rows.Scan(&produkSLA.KodePembayaran, &produkSLA.NamaPembayaran); er != nil {
			return lists, er
		}

		lists = append(lists, produkSLA)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (r *referensiApexMysqlImpl) GetListsTabunganIntegrasi() (lists []entities.TabunganIntegrasi, er error) {
	rows, er := r.apexDb.Query("SELECT kode_integrasi, deskripsi_integrasi FROM tab_integrasi")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var tabIntegrasi entities.TabunganIntegrasi
		if er = rows.Scan(&tabIntegrasi.KodeIntegrasi, &tabIntegrasi.DeskripsiIntegrasi); er != nil {
			return lists, er
		}

		lists = append(lists, tabIntegrasi)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}
