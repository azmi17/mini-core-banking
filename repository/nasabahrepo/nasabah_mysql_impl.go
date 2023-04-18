package nasabahrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
)

func newNasbahMysqlImpl(apexConn *sql.DB) NasabahRepo {
	return &NasbahMysqlImpl{
		apexDb: apexConn,
	}
}

type NasbahMysqlImpl struct {
	apexDb *sql.DB
}

func (n *NasbahMysqlImpl) CreateNasabah(newNasabah entities.Nasabah) (nasabah entities.Nasabah, er error) {
	stmt, er := n.apexDb.Prepare(`INSERT INTO nasabah(
		nasabah_id, 
		nama_nasabah, 
		alamat, 
		telpon, 
		jenis_kelamin, 
		tempatlahir, 
		tgllahir, 
		jenis_id, 
		no_id, 
		kode_group1, 
		kode_group2, 
		kode_group3, 
		kode_agama, 
		desa, 
		kecamatan, 
		kota_kab, 
		propinsi, 
		verifikasi, 
		hp, 
		tgl_register, 
		nama_ibu_kandung, 
		kodepos, 
		kode_kantor, 
		userid, 
		nama_alias, 
		status_gelar, 
		jenis_debitur, 
		kode_area, 
		negara_domisili, 
		gol_debitur, 
		langgar_bmpk, 
		lampaui_bmpk, 
		nama_nasabah_sid, 
		alamat2, 
		flag_masa_berlaku,
		status_marital, 
		hp1, 
		hp2, 
		status_tempat_tinggal, 
		masa_berlaku_ktp
	) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	if er != nil {
		return nasabah, errors.New(fmt.Sprint("error while prepare add nasabah : ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(
		newNasabah.Nasabah_Id,
		newNasabah.Nama_Nasabah,
		newNasabah.Alamat,
		newNasabah.Telepon,
		newNasabah.Jenis_Kelamin,
		newNasabah.TempatLahir,
		newNasabah.TglLahir,
		newNasabah.Jenis_Id,
		newNasabah.No_Id,
		newNasabah.Kode_Group1,
		newNasabah.Kode_Group2,
		newNasabah.Kode_Group3,
		newNasabah.Kode_Agama,
		newNasabah.Desa,
		newNasabah.Kecamatan,
		newNasabah.Kota_Kab,
		newNasabah.Provinsi,
		newNasabah.Verifikasi,
		newNasabah.Hp,
		newNasabah.Tgl_Register,
		newNasabah.Nama_Ibu_Kandung,
		newNasabah.Kodepos,
		newNasabah.Kode_Kantor,
		newNasabah.UserId,
		newNasabah.Nama_Alias,
		newNasabah.Status_Gelar,
		newNasabah.Jenis_Debitur,
		newNasabah.Kode_Area,
		newNasabah.Negara_Domisili,
		newNasabah.Gol_Debitur,
		newNasabah.Langgar_Bmpk,
		newNasabah.Lampaui_Bmpk,
		newNasabah.Nama_Nasabah_Sid,
		newNasabah.Alamat2,
		newNasabah.Flag_Masa_Berlaku,
		newNasabah.Status_Marital,
		newNasabah.Hp1,
		newNasabah.Hp2,
		newNasabah.Status_Tempat_Tinggal,
		newNasabah.Masa_Berlaku_Ktp,
	); er != nil {
		return nasabah, errors.New(fmt.Sprint("error while add nasabah : ", er.Error()))
	} else {
		return newNasabah, nil
	}

}

func (n *NasbahMysqlImpl) UpdateNasabah(updNasabah entities.Nasabah) (nasabah entities.Nasabah, er error) {

	stmt, er := n.apexDb.Prepare(`UPDATE nasabah SET 
		nama_nasabah = ?, 
		nama_nasabah_sid = ?, 
		nama_alias = ?, 
		alamat = ?, 
		alamat2 = ?, 
		telpon = ?, 
		hp = ?, 
		hp1 = ?,  
		hp2 = ?, 
		userid = ?  
		WHERE nasabah_id = ?`)
	if er != nil {
		return nasabah, errors.New(fmt.Sprint("error while prepare update nasabah : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(
		updNasabah.Nama_Nasabah,
		updNasabah.Nama_Nasabah_Sid,
		updNasabah.Nama_Alias,
		updNasabah.Alamat,
		updNasabah.Alamat2,
		updNasabah.Telepon,
		updNasabah.Hp,
		updNasabah.Hp1,
		updNasabah.Hp2,
		updNasabah.UserId,
		updNasabah.Nasabah_Id); er != nil {
		return nasabah, errors.New(fmt.Sprint("error while update nasabah : ", er.Error()))
	}
	return updNasabah, nil
}

func (n *NasbahMysqlImpl) HardDeleteNasabah(kodeLkm string) (er error) {
	stmt, er := n.apexDb.Prepare("DELETE FROM nasabah WHERE nasabah_id = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete nasabah : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete nasabah : ", er.Error()))
	}
	return nil
}

func (n *NasbahMysqlImpl) DeleteNasabah(kodeLkm string) (er error) {
	stmt, er := n.apexDb.Prepare("UPDATE nasabah SET nasabah_id = ? WHERE nasabah_id = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete nasabah : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec("DEL-"+kodeLkm, kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete nasabah : ", er.Error()))
	}
	return nil
}

func (n *NasbahMysqlImpl) FindNasabahLkm(nasabahId string) (nasabahLKM entities.Nasabah, er error) {
	row := n.apexDb.QueryRow(`SELECT
		nasabah_id, 
		nama_nasabah
	FROM nasabah WHERE nasabah_id = ? LIMIT 1`, nasabahId)
	er = row.Scan(
		&nasabahLKM.Nasabah_Id,
		&nasabahLKM.Nama_Nasabah,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return nasabahLKM, err.NoRecord
		} else {
			return nasabahLKM, errors.New(fmt.Sprint("error while get nasabah LKM: ", er.Error()))
		}
	}
	return
}
