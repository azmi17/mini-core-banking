package apexvirtualaccountrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/constants"
	"new-apex-api/entities/err"
	"new-apex-api/helper"
	"new-apex-api/repository/constant"
	"new-apex-api/repository/nasabahrepo"
	"new-apex-api/repository/tabtransrepo"
	"new-apex-api/repository/tabunganrepo"
	"os"
	"strings"
	"time"
)

func newApexVirtualAccountMysqlImpl(apexConn *sql.DB) ApexVirtualAccountRepo {
	return &ApexVirtualAccountMysqlImpl{
		apexDb: apexConn,
	}
}

type ApexVirtualAccountMysqlImpl struct {
	apexDb *sql.DB
}

func (va *ApexVirtualAccountMysqlImpl) CreateApexVirtualAccount(payload entities.CreateVirtualAccount) (resp entities.VirtualAccountResponse, er error) {
	nasabahRepo, _ := nasabahrepo.NewNasabahRepo()
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	nID, er := tabtransRepo.GetNextNasabahID()
	if er != nil {
		return resp, er
	}

	norek, er := tabtransRepo.GetNextNoRekeningBiggerThanFour()
	if er != nil {
		return resp, er
	}

	nasabah := entities.Nasabah{}

	// dynamic val:
	nasabah.Nasabah_Id = nID
	nasabah.Nama_Nasabah = payload.NamaLembaga
	nasabah.Kode_Group2 = payload.VendorCode
	nasabah.Alamat = payload.Alamat
	nasabah.Telepon = payload.Telepon
	nasabah.Hp = payload.Telepon
	nasabah.UserId = payload.UserID
	nasabah.Nama_Alias = payload.NamaLembaga
	nasabah.Nama_Nasabah_Sid = payload.NamaLembaga
	nasabah.Alamat2 = payload.Alamat
	nasabah.Hp1 = payload.Telepon
	nasabah.Hp2 = payload.Telepon

	// Set static Val:
	nasabah.Jenis_Kelamin = constants.JenisKelamin
	nasabah.TempatLahir = constants.TempatLahir
	nasabah.TglLahir = time.Now()
	nasabah.Jenis_Id = constants.JenisId
	nasabah.No_Id = helper.GenerateIdKTP()
	nasabah.Kode_Group1 = constants.Group1
	nasabah.Kode_Group3 = constants.Group3
	nasabah.Kode_Agama = constants.Agama
	nasabah.Desa = constants.Desa
	nasabah.Kecamatan = constants.Kecamatan
	nasabah.Kota_Kab = constants.KotaKabKode
	nasabah.Provinsi = constants.Prov
	nasabah.Verifikasi = constants.Verifikasi
	nasabah.Tgl_Register = time.Now()
	nasabah.Nama_Ibu_Kandung = constants.NamaIbu
	nasabah.Kodepos = constants.KodePos
	nasabah.Kode_Kantor = constants.KodeKantor
	nasabah.Status_Gelar = constants.StatusGelar
	nasabah.Jenis_Debitur = constants.JenisDebitur
	nasabah.Kode_Area = constants.KodeArea
	nasabah.Negara_Domisili = constants.NegaraDomisili
	nasabah.Gol_Debitur = constants.GolDebitur
	nasabah.Langgar_Bmpk = constants.LampauLanggarBmpk
	nasabah.Lampaui_Bmpk = constants.LampauLanggarBmpk
	nasabah.Flag_Masa_Berlaku = constants.FlagMasaBerlaku
	nasabah.Status_Marital = constants.StatusMarital
	nasabah.Status_Tempat_Tinggal = constants.StatusTempatTinggal
	nasabah.Masa_Berlaku_Ktp = time.Now().AddDate(15, 0, 0)
	if nasabah, er = nasabahRepo.CreateNasabah(nasabah); er != nil {
		return resp, er
	}

	tabung := entities.Tabung{}

	// dynamic val:
	tabung.No_Rekening = norek
	tabung.Status = payload.Status
	tabung.Nasabah_Id = nID
	tabung.UserId = payload.UserID
	tabung.No_Rekening_Virtual = norek
	tabung.Kode_Group2 = payload.VendorCode
	tabung.Kode_Integrasi = payload.KodePembayaranSLA
	tabung.Kode_Produk = payload.KodePembayaranSLA

	// Set static Val:
	tabung.Kode_Bi_Pemilik = constants.KodeBIPemilik
	tabung.Suku_Bunga = constants.ZeroValInt
	tabung.Persen_Pph = constants.ZeroValInt
	tabung.Tgl_Register = time.Now()
	tabung.Saldo_Akhir = constants.ZeroValInt
	tabung.Kode_Group1 = constants.KdGroup1
	tabung.Verifikasi = constants.Verifikasi
	tabung.Kode_Kantor = constants.KodeKantor
	tabung.Kode_Group3 = constants.EmptyStr
	tabung.Minimum = constants.ZeroValInt
	tabung.Setoran_Minimum = constants.ZeroValInt
	tabung.Jkw = constants.ZeroValInt
	tabung.Abp = constants.ZeroValInt
	tabung.Setoran_Wajib = constants.ZeroValInt
	tabung.Adm_Per_Bln = constants.ZeroValInt
	tabung.Target_Nominal = constants.ZeroValInt
	tabung.Saldo_Akhir_Titipan_bunga = constants.ZeroValInt
	tabung.Kode_Bi_Lokasi = constants.KdBILokasi
	tabung.Saldo_Akhir_Titipan_bunga = constants.ZeroValInt
	tabung.Saldo_Titipan_Bunga_Ks = constants.ZeroValInt
	tabung.Saldo_Blokir = constants.ZeroValInt
	tabung.Premi = constants.ZeroValInt
	tabung.Kode_Keterkaitan = constants.KdKeterkaitan
	tabung.Kode_Kantor_Kas = constants.KdKantorKas
	tabung.IsSaldoChecked = constants.StatusAktif
	tabung.FlagPayEchannel = constants.FlagPayEchannelActive
	if tabung, er = tabunganRepo.CreateTabung(tabung); er != nil {
		return resp, er
	}

	resp.NasabahID = nID
	resp.NamaLembaga = payload.NamaLembaga
	resp.Vendor = payload.VendorCode
	resp.NomorVA = norek
	resp.SLAType = payload.KodePembayaranSLA
	resp.Alamat = payload.Alamat
	resp.Telepon = payload.Telepon
	resp.Saldo = tabung.Saldo_Akhir
	resp.Status = payload.Status

	return resp, nil
}

func (va *ApexVirtualAccountMysqlImpl) GetListsApexVirtualAccount(limitOffset entities.LimitOffsetLkmUri) (list []entities.VirtualAccountResponse, er error) {

	item := strings.Split(os.Getenv("app.list_kode_tipe_sla"), ",")
	listKodeTipeSLA := (strings.Join(item, ","))

	args := []interface{}{}
	limit := ""
	if limitOffset.Limit > 0 {
		limit = "LIMIT ? OFFSET ?"
		args = append(args, limitOffset.Limit, limitOffset.Offset)
	} else {
		limit = "LIMIT ? OFFSET ?"
		args = append(args, -1, limitOffset.Offset)
	}
	rows, er := va.apexDb.Query(`SELECT
	t.no_rekening AS 'Nomor VA', 
	t.nasabah_id AS NasabahID,
	n.nama_nasabah AS 'Nama Lembaga',
	g.deskripsi_group2 AS 'Vendor',
	tb.deskripsi_produk AS 'Jenis SLA',
	n.alamat AS 'Alamat',
	n.telpon AS 'Telepon',
	t.saldo_akhir AS 'Saldo',
	t.status
   FROM tabung AS t 
		INNER JOIN tab_produk AS tb ON (t.kode_produk = tb.kode_produk)
		INNER JOIN nasabah AS n ON (t.nasabah_id = n.nasabah_id)
		INNER JOIN tab_kode_group2 AS g ON (t.kode_group2 = g.kode_group2)
   WHERE LENGTH(t.no_rekening)>4 AND t.kode_produk IN (`+listKodeTipeSLA+`) AND t.status=1 ORDER BY t.nasabah_id ASC `+limit+``, args...)
	if er != nil {
		return list, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var data entities.VirtualAccountResponse
		if er = rows.Scan(
			&data.NomorVA,
			&data.NasabahID,
			&data.NamaLembaga,
			&data.Vendor,
			&data.SLAType,
			&data.Alamat,
			&constant.SQLKontak,
			&data.Saldo,
			&data.Status,
		); er != nil {
			return list, er
		}
		data.Telepon = constant.SQLKontak.String
		list = append(list, data)
	}

	if len(list) == 0 {
		return list, nil
	} else {
		return
	}
}

func (va *ApexVirtualAccountMysqlImpl) updateNasabahBelongingToVA(payload entities.UpdateVirtualAccount) (er error) {
	stmt, er := va.apexDb.Prepare("UPDATE nasabah SET nama_nasabah = ?, alamat = ?, kode_group2 = ?, telpon = ?, userid = ? WHERE nasabah_id = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare update nasabah VA: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(payload.NamaLembaga, payload.Alamat, payload.VendorCode, payload.Telepon, payload.UserID, payload.NasabahID); er != nil {
		return errors.New(fmt.Sprint("error while update nasabah VA: ", er.Error()))
	}

	return nil
}

func (va *ApexVirtualAccountMysqlImpl) updateTabungBelongingToVA(payload entities.UpdateVirtualAccount) (er error) {
	stmt, er := va.apexDb.Prepare("UPDATE tabung SET kode_group2 = ?, kode_produk = ?, status = ?, kode_integrasi = ?, userid = ? WHERE nasabah_id = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare update tabung VA: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(payload.VendorCode, payload.KodePembayaranSLA, payload.Status, payload.KodePembayaranSLA, payload.UserID, payload.NasabahID); er != nil {
		return errors.New(fmt.Sprint("error while update tabung VA: ", er.Error()))
	}

	return nil
}

func (va *ApexVirtualAccountMysqlImpl) UpdateApexVirtualAccount(payload entities.UpdateVirtualAccount) (er error) {

	er = va.updateNasabahBelongingToVA(payload)
	if er != nil {
		return er
	}

	er = va.updateTabungBelongingToVA(payload)
	if er != nil {
		return er
	}

	return nil

}

func (va *ApexVirtualAccountMysqlImpl) GetListsSLATransactionVirtualAccount(payload entities.GetListTabtrans, limitOffset entities.LimitOffsetLkmUri) (lists []entities.VirtualAccountTransResponse, er error) {
	var rows *sql.Rows

	item := strings.Split(os.Getenv("app.list_kode_tipe_sla"), ",")
	listKodeTipeSLA := (strings.Join(item, ","))

	/*
		Dibawah adalah slice of interface untuk kebutuhan custom query,
		di deklarasikan dengan scope Global Variable agar re-usable pada 2 kondisi
	*/
	args := []interface{}{}
	sqlCond := ""
	sqlStmt := `
	SELECT 
		trans.tabtrans_id,
		n.nama_nasabah AS 'Nama Lembaga',
		t.no_rekening AS 'Nomor VA',
		tb.deskripsi_produk AS 'Jenis SLA',
		DATE_FORMAT(trans.tgl_trans,'%d/%m/%Y') AS tgltrans,
		trans.kuitansi,
		trans.keterangan,
		trans.pokok
	FROM tabung AS t 
		INNER JOIN tab_produk AS tb ON (t.kode_produk = tb.kode_produk)
		INNER JOIN nasabah AS n ON (t.nasabah_id = n.nasabah_id)
		INNER JOIN tab_kode_group2 AS g ON (t.kode_group2 = g.kode_group2)
		INNER JOIN tabtrans AS trans ON (t.no_rekening = trans.no_rekening)
	WHERE LENGTH(t.no_rekening)>4 
	AND t.kode_produk IN (` + listKodeTipeSLA + `)
	AND trans.kode_trans <> '900' AND `

	if payload.Filter == "" {
		if limitOffset.Limit > 0 {
			sqlCond = "trans.tgl_trans >= ? AND trans.tgl_trans <= ? ORDER BY trans.tgl_trans ASC LIMIT ? OFFSET ?"
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, limitOffset.Limit, limitOffset.Offset)
		} else {
			sqlCond = "trans.tgl_trans >= ? AND trans.tgl_trans <= ? ORDER BY trans.tgl_trans ASC LIMIT ? OFFSET ?"
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, -1, limitOffset.Offset)
		}
		rows, er = va.apexDb.Query(sqlStmt+sqlCond+``, args...)
	} else {
		if limitOffset.Limit > 0 {
			sqlCond = `
			trans.tgl_trans >= ? 
			AND trans.tgl_trans <= ? 
			AND (trans.no_rekening LIKE "%` + payload.Filter + `%" OR trans.kuitansi LIKE "%` + payload.Filter + `%") 
			ORDER BY trans.tgl_trans ASC LIMIT ? OFFSET ?`
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, limitOffset.Limit, limitOffset.Offset)
		} else {
			sqlCond = `
			trans.tgl_trans >= ? 
			AND trans.tgl_trans <= ? 
			AND t.no_rekening = ? 
			AND (trans.no_rekening LIKE "%?%" OR trans.kuitansi LIKE "%?%") 
			ORDER BY trans.tgl_trans ASC LIMIT ? OFFSET ?`
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, -1, limitOffset.Offset)
		}
		rows, er = va.apexDb.Query(sqlStmt+sqlCond+``, args...)
	}

	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var vaSLAtrans entities.VirtualAccountTransResponse
		if er = rows.Scan(
			&vaSLAtrans.TabtransID,
			&vaSLAtrans.NamaLembaga,
			&vaSLAtrans.NoVA,
			&vaSLAtrans.SLAType,
			&vaSLAtrans.TglTrans,
			&vaSLAtrans.Kuitansi,
			&vaSLAtrans.Keterangan,
			&vaSLAtrans.Nominal,
		); er != nil {
			return lists, er
		}
		lists = append(lists, vaSLAtrans)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	}

	return lists, nil
}

func (va *ApexVirtualAccountMysqlImpl) SoftDeleteVAAccount(vaNumber ...string) (er error) {
	nasabahRepo, _ := nasabahrepo.NewNasabahRepo()
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	if er = nasabahRepo.SoftDeleteNasabah(vaNumber...); er != nil {
		return er
	}

	if er = tabunganRepo.SoftDeleteTabung(vaNumber...); er != nil {
		return er
	}

	if er = tabtransRepo.SoftDeleteApexTransaction(vaNumber...); er != nil {
		return er
	}

	return nil
}
