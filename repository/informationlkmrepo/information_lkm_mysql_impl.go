package informationlkmrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/entities/web"
)

func newInformationLKMMysqlImpl(apexConn *sql.DB) InformationLKMRepo {
	return &informationLKMMysqlImpl{
		apexDb: apexConn,
	}
}

type informationLKMMysqlImpl struct {
	apexDb *sql.DB
}

func (i *informationLKMMysqlImpl) RekeningKoranLKMDetailHeader(kodeLKM string) (header web.RekeningKoranHeader, er error) {
	row := i.apexDb.QueryRow(`SELECT
		t.no_rekening,
		n.nama_nasabah,
		tp.deskripsi_produk,
		g.deskripsi_group2
	FROM tabung AS t 
		INNER JOIN nasabah AS n ON (t.nasabah_id = n.nasabah_id)	
		INNER JOIN tab_kode_group2 AS g ON (t.kode_group2 = g.kode_group2)	
		INNER JOIN tab_produk AS tp ON (t.kode_produk = tp.kode_produk)
		WHERE t.no_rekening = ?`, kodeLKM)
	if er != nil {
		return header, er
	}
	er = row.Scan(
		&header.Norek,
		&header.NamaLembaga,
		&header.ProdukTab,
		&header.NamaSC,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return header, err.NoRecord
		} else {
			return header, errors.New(fmt.Sprint("error while get saldo awal: ", er.Error()))
		}
	}
	return header, nil
}

func (i *informationLKMMysqlImpl) LKMInformation(kodeLKM string) (lkmInfo entities.LkmInfo, er error) {
	row := i.apexDb.QueryRow(`SELECT
		t.no_rekening,
		t.nasabah_id,
		n.nama_nasabah,
		DATE_FORMAT(n.tgl_register, "%d/%m/%Y") AS Tgl_Register,
		t.minimum,
		t.setoran_minimum,
		t.saldo_akhir,
		t.status
	FROM tabung AS t 
		INNER JOIN nasabah AS n ON (t.nasabah_id = n.nasabah_id)	
		WHERE t.no_rekening= ?`, kodeLKM)
	if er != nil {
		return lkmInfo, er
	}
	er = row.Scan(
		&lkmInfo.KodeLKM,
		&lkmInfo.NasabahId,
		&lkmInfo.NamaLembaga,
		&lkmInfo.Alamat,
		&lkmInfo.TanggalRegister,
		&lkmInfo.Plafond,
		&lkmInfo.SetoranMinimal,
		&lkmInfo.SaldoAkhir,
		&lkmInfo.StatusRekening,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return lkmInfo, err.NoRecord
		} else {
			return lkmInfo, errors.New(fmt.Sprint("error while get lkm information: ", er.Error()))
		}
	}
	return lkmInfo, nil
}
