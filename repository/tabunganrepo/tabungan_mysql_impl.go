package tabunganrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/entities/web"
	"new-apex-api/repository/constant"
)

func newTabunganMysqlImpl(apexConn *sql.DB) TabunganRepo {
	return &TabunganMysqlImpl{
		apexDb: apexConn,
	}
}

type TabunganMysqlImpl struct {
	apexDb *sql.DB
}

func (t *TabunganMysqlImpl) CreateTabung(newTabung entities.Tabung) (tabung entities.Tabung, er error) {
	stmt, er := t.apexDb.Prepare(`INSERT INTO tabung(
		no_rekening,
		nasabah_id,
		kode_bi_pemilik,
		suku_bunga,
		persen_pph,
		tgl_register,
		saldo_akhir,
		kode_group1,
		kode_group2,
		verifikasi,
		status,
		kode_kantor,
		kode_integrasi,
		kode_produk,
		userid,
		kode_group3,
		minimum,
		setoran_minimum,
		jkw,
		abp,
		setoran_wajib,
		adm_per_bln,
		target_nominal,
		saldo_akhir_titipan_bunga,
		kode_bi_lokasi,
		saldo_titipan_pokok,
		saldo_titipan_bunga_ks,
		saldo_blokir,
		premi,
		kode_keterkaitan,
		kode_kantor_kas,
		no_rekening_virtual
	) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	if er != nil {
		return tabung, errors.New(fmt.Sprint("error while prepare add tabung : ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Exec..
	if _, er := stmt.Exec(
		newTabung.No_Rekening,
		newTabung.Nasabah_Id,
		newTabung.Kode_Bi_Pemilik,
		newTabung.Suku_Bunga,
		newTabung.Persen_Pph,
		newTabung.Tgl_Register,
		newTabung.Saldo_Akhir,
		newTabung.Kode_Group1,
		newTabung.Kode_Group2,
		newTabung.Verifikasi,
		newTabung.Status,
		newTabung.Kode_Kantor,
		newTabung.Kode_Integrasi,
		newTabung.Kode_Produk,
		newTabung.UserId,
		newTabung.Kode_Group3,
		newTabung.Minimum,
		newTabung.Setoran_Minimum,
		newTabung.Jkw,
		newTabung.Abp,
		newTabung.Setoran_Wajib,
		newTabung.Adm_Per_Bln,
		newTabung.Target_Nominal,
		newTabung.Saldo_Akhir_Titipan_bunga,
		newTabung.Kode_Bi_Lokasi,
		newTabung.Saldo_Titipan_Pokok,
		newTabung.Saldo_Titipan_Bunga_Ks,
		newTabung.Saldo_Blokir,
		newTabung.Premi,
		newTabung.Kode_Keterkaitan,
		newTabung.Kode_Kantor_Kas,
		newTabung.No_Rekening_Virtual); er != nil {
		return tabung, errors.New(fmt.Sprint("error while add tabung: ", er.Error()))
	} else {
		return newTabung, nil
	}

}

func (t *TabunganMysqlImpl) HardDeleteTabung(kodeLkm string) (er error) {

	stmt, er := t.apexDb.Prepare("DELETE FROM tabung WHERE no_rekening = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete tabung: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete tabung : ", er.Error()))
	}

	return nil
}

func (t *TabunganMysqlImpl) DeleteTabung(kodeLkm string) (er error) {

	stmt, er := t.apexDb.Prepare(`UPDATE tabung SET status = 0, no_rekening = ? WHERE no_rekening = ?`)
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete tabung: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec("DEL-"+kodeLkm, kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete tabung: ", er.Error()))
	}

	return nil
}

func (t *TabunganMysqlImpl) GetTabScGroup() (list []web.TabSCGroup, er error) {
	rows, er := t.apexDb.Query("SELECT kode_group2, deskripsi_group2 FROM tab_kode_group2")
	if er != nil {
		return list, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var scGroup web.TabSCGroup
		if er = rows.Scan(&scGroup.KodeGroup, &scGroup.DeskripsiGroup); er != nil {
			return list, er
		}

		list = append(list, scGroup)
	}

	if len(list) == 0 {
		return list, err.NoRecord
	} else {
		return
	}
}

func (t *TabunganMysqlImpl) GetTabDetailInfo(KodeLkm string) (detail web.GetDetailLKMInfo, er error) {
	row := t.apexDb.QueryRow(`SELECT
		t.no_rekening, 
		g.deskripsi_group2,
		n.nama_nasabah,
		n.alamat,
		n.telpon,
		t.no_rekening,
		t.saldo_akhir,
		t.minimum,
		t.status
	FROM tabung AS t 
	INNER JOIN nasabah AS n ON(t.nasabah_id=n.nasabah_id) 
	INNER JOIN tab_kode_group2 AS g ON(t.kode_group2 = g.kode_group2) 
	WHERE t.status=1 AND t.no_rekening = ?`, KodeLkm)

	er = row.Scan(
		&detail.KodeLembaga,
		&constant.SQLVendor,
		&detail.NamaLembaga,
		&constant.SQLAlamat,
		&constant.SQLKontak,
		&detail.NoRekening,
		&detail.Saldo,
		&constant.SQLPlafond,
		&detail.StatusTab,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return detail, err.NoRecord
		} else {
			return detail, errors.New(fmt.Sprint("error while get instution detail: ", er.Error()))
		}
	}
	detail.Vendor = constant.SQLVendor.String
	detail.Alamat = constant.SQLAlamat.String
	detail.Kontak = constant.SQLKontak.String
	detail.Plafond = constant.SQLPlafond.Float64

	// constant.ConvertSQLDataType()
	return
}

func (t *TabunganMysqlImpl) GetTabInfoList(limitOffset web.LimitOffsetLkmUri) (list []web.GetDetailLKMInfo, er error) {

	args := []interface{}{}
	limit := ""
	if limitOffset.Limit > 0 {
		limit = "LIMIT ? OFFSET ?"
		args = append(args, limitOffset.Limit, limitOffset.Offset)
	} else {
		limit = "LIMIT ? OFFSET ?"
		args = append(args, -1, limitOffset.Offset)
	}
	rows, er := t.apexDb.Query(`SELECT 
		t.no_rekening, 
		g.deskripsi_group2,
		n.nama_nasabah,
		n.alamat,
		n.telpon,
		t.no_rekening,
		t.saldo_akhir,
		t.minimum,
		t.status
		FROM tabung AS t 
	INNER JOIN nasabah AS n ON(t.nasabah_id=n.nasabah_id) 
	INNER JOIN tab_kode_group2 AS g ON(t.kode_group2 = g.kode_group2) `+limit+``, args...)
	if er != nil {
		return list, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var lkmList web.GetDetailLKMInfo
		if er = rows.Scan(
			&lkmList.KodeLembaga,
			&constant.SQLVendor,
			&lkmList.NamaLembaga,
			&constant.SQLAlamat,
			&constant.SQLKontak,
			&lkmList.NoRekening,
			&lkmList.Saldo,
			&constant.SQLPlafond,
			&lkmList.StatusTab,
		); er != nil {
			return list, er
		}

		lkmList.Vendor = constant.SQLVendor.String
		lkmList.Alamat = constant.SQLAlamat.String
		lkmList.Kontak = constant.SQLKontak.String
		lkmList.Plafond = constant.SQLPlafond.Float64
		list = append(list, lkmList)
	}

	if len(list) == 0 {
		return list, nil // no.record
	} else {
		return
	}
}

func (t *TabunganMysqlImpl) FindTabunganLkm(tabunganLkm string) (tabung entities.Tabung, er error) {
	row := t.apexDb.QueryRow(`SELECT
		no_rekening, 
		nasabah_id,
		saldo_akhir,
		status
	FROM tabung WHERE no_rekening = ? LIMIT 1`, tabunganLkm)
	er = row.Scan(
		&tabung.No_Rekening,
		&tabung.Nasabah_Id,
		&tabung.Saldo_Akhir,
		&tabung.Status,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return tabung, err.NoRecord
		} else {
			return tabung, errors.New(fmt.Sprint("error while get tabungan LKM: ", er.Error()))
		}
	}
	return
}

func (t *TabunganMysqlImpl) GetRekeningLKMByStatusActive() (lists []string, er error) {
	rows, er := t.apexDb.Query("SELECT no_rekening FROM tabung WHERE status = 1")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var list web.LKMlist
		if er = rows.Scan(&list.KodeLKM); er != nil {
			return lists, er
		}

		lists = append(lists, list.KodeLKM)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}
