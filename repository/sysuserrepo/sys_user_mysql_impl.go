package sysuserrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/repository/constant"
	"os"
	"strings"
)

func newSysUserMysqlImpl(apexConn *sql.DB) SysUserRepo {
	return &SysUserMysqlImpl{
		apexDb: apexConn,
	}
}

type SysUserMysqlImpl struct {
	apexDb *sql.DB
}

func (s *SysUserMysqlImpl) GetSingleUserByUserName(userName string) (user entities.ManajemenUserDataResponse, er error) {
	row := s.apexDb.QueryRow(`SELECT 
		user_id,
		user_name,
		nama_lengkap,
		jabatan,
		unit_kerja,
		COALESCE(DATE_FORMAT(TGL_EXPIRED, "%d/%m/%Y"),'') AS tgl_expired,
		status_aktif,
		user_code
	FROM sys_daftar_user 
	WHERE user_name = ? LIMIT 1`, userName)
	er = row.Scan(
		&user.User_ID,
		&user.User_Name,
		&user.Nama_Lengkap,
		&constant.SQLJabatan,
		&user.Unit_Kerja,
		&user.Tgl_Expired,
		&user.StatusAktif,
		&user.User_Code,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return user, err.NoRecord
		} else {
			return user, errors.New(fmt.Sprint("error while get user: ", er.Error()))
		}
	}
	user.Jabatan = constant.SQLJabatan.String

	return
}

func (s *SysUserMysqlImpl) GetUserID(userID int) (user entities.SysDaftarUser, er error) {
	row := s.apexDb.QueryRow(`SELECT 
		user_id,
		user_name
	FROM sys_daftar_user 
	WHERE user_id = ? LIMIT 1`, userID)
	er = row.Scan(
		&user.User_Id,
		&user.User_Name,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return user, err.NoRecord
		} else {
			return user, errors.New(fmt.Sprint("error while get user id: ", er.Error()))
		}
	}
	return
}

func (s *SysUserMysqlImpl) GetListOfUsers(payload entities.GlobalFilter, limitOffset entities.LimitOffsetLkmUri) (lists []entities.ManajemenUserDataResponse, er error) {
	var rows *sql.Rows

	args := []interface{}{}
	sqlCond := ""
	sqlStmt := `SELECT 
		user_id,
		user_name,
		nama_lengkap,
		jabatan,
		unit_kerja,
		COALESCE(DATE_FORMAT(TGL_EXPIRED, "%d/%m/%Y"),'') AS tgl_expired,
		status_aktif,
		user_code
	FROM sys_daftar_user `

	if payload.Filter == "" {
		if limitOffset.Limit > 0 {
			sqlCond = "LIMIT ? OFFSET ?"
			args = append(args, limitOffset.Limit, limitOffset.Offset)
		} else {
			sqlCond = "LIMIT ? OFFSET ?"
			args = append(args, -1, limitOffset.Offset)
		}
		rows, er = s.apexDb.Query(sqlStmt+sqlCond+``, args...)
	} else {
		if limitOffset.Limit > 0 {
			sqlCond = `
			WHERE
			(user_name LIKE "%` + payload.Filter + `%" OR nama_lengkap LIKE "%` + payload.Filter + `%") 
			LIMIT ? OFFSET ?`
			args = append(args, limitOffset.Limit, limitOffset.Offset)
		} else {
			sqlCond = `
			WHERE
			(user_name LIKE "%` + payload.Filter + `%" OR nama_lengkap LIKE "%` + payload.Filter + `%") 
			LIMIT ? OFFSET ?`
			args = append(args, -1, limitOffset.Offset)
		}
		rows, er = s.apexDb.Query(sqlStmt+sqlCond+``, args...)
	}
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var users entities.ManajemenUserDataResponse
		if er = rows.Scan(
			&users.User_ID,
			&constant.SQLUSerName,
			&constant.SQLNamaLengkap,
			&constant.SQLJabatan,
			&constant.SQLUnitKerja,
			&users.Tgl_Expired,
			&users.StatusAktif,
			&users.User_Code,
		); er != nil {
			return lists, er
		}
		users.User_Name = constant.SQLUSerName.String
		users.Nama_Lengkap = constant.SQLNamaLengkap.String
		users.Jabatan = constant.SQLJabatan.String
		users.Unit_Kerja = constant.SQLUnitKerja.String

		lists = append(lists, users)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (s *SysUserMysqlImpl) CreateSysDaftarUser(newSysUser entities.SysDaftarUser) (sysUser entities.SysDaftarUser, er error) {

	stmt, er := s.apexDb.Prepare(`INSERT INTO sys_daftar_user(
		user_name,
		user_password,
		nama_lengkap,
		unit_kerja,
		jabatan,
		user_code,
		tgl_expired,
		user_web_password,
		flag,
		status_aktif,
		penerimaan,
		pengeluaran
	) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)`)
	if er != nil {
		return sysUser, errors.New(fmt.Sprint("error while prepare add sys user: ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Exec..
	if stmt, er := stmt.Exec(
		newSysUser.User_Name,
		newSysUser.User_Password,
		newSysUser.Nama_Lengkap,
		newSysUser.Unit_Kerja,
		newSysUser.Jabatan,
		newSysUser.User_Code,
		newSysUser.Tgl_Expired,
		newSysUser.User_Web_Password,
		newSysUser.Flag,
		newSysUser.Status_Aktif,
		newSysUser.Penerimaan,
		newSysUser.Pengeluaran); er != nil {
		return sysUser, errors.New(fmt.Sprint("error while add sys user: ", er.Error()))
	} else {

		lastId, txErr := stmt.LastInsertId()
		if txErr != nil {
			return sysUser, errors.New(fmt.Sprint("error while get last insert id sys user: ", txErr.Error()))
		}
		newSysUser.User_Id = int(lastId)

		return newSysUser, nil
	}

}

func (s *SysUserMysqlImpl) UpdateSysDaftarUser(updSysuser entities.SysDaftarUser) (sysUser entities.SysDaftarUser, er error) {

	thisRepo, _ := NewSysUserRepo()
	user, er := thisRepo.GetSingleUserByUserName(updSysuser.User_Name)
	if er != nil {
		return sysUser, err.NoRecord
	}

	stmt, er := s.apexDb.Prepare(`UPDATE sys_daftar_user SET 
	nama_lengkap = ?,
	unit_kerja = ?,
	jabatan = ?,
	user_code = ?,
	tgl_expired = ?,
	user_web_password = ?,
	flag = ?,
	status_aktif = ?,
	penerimaan = ?,
	pengeluaran = ?
		WHERE user_name = ?`)
	if er != nil {
		return sysUser, errors.New(fmt.Sprint("error while prepare update user: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(
		updSysuser.Nama_Lengkap,
		updSysuser.Unit_Kerja,
		updSysuser.Jabatan,
		updSysuser.User_Code,
		updSysuser.Tgl_Expired,
		updSysuser.User_Web_Password,
		updSysuser.Flag,
		updSysuser.Status_Aktif,
		updSysuser.Penerimaan,
		updSysuser.Pengeluaran,
		updSysuser.User_Name); er != nil {
		return sysUser, errors.New(fmt.Sprint("error while update user: ", er.Error()))
	}

	updSysuser.User_Id = user.User_ID

	return updSysuser, nil

}

func (s *SysUserMysqlImpl) HardDeleteSysDaftarUser(kodeLkm ...string) (er error) {

	stmt, er := s.apexDb.Prepare("DELETE FROM sys_daftar_user WHERE user_name = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete user : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	for _, v := range kodeLkm {
		if _, er := stmt.Exec(v); er != nil {
			return errors.New(fmt.Sprint("error while delete user : ", er.Error()))
		}
	}

	return nil
}

func (s *SysUserMysqlImpl) SoftDeleteSysDaftarUser(kodeLkm ...string) (er error) {

	// thisRepo, _ := NewSysUserRepo()
	// _, er = thisRepo.GetSingleUserByUserName(kodeLkm)
	// if er != nil {
	// 	return err.NoRecord
	// }

	stmt, er := s.apexDb.Prepare("UPDATE sys_daftar_user SET status_aktif = 0, user_name = ? WHERE user_name = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete user : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	for _, v := range kodeLkm {
		if _, er := stmt.Exec("DEL-"+v, v); er != nil {
			return errors.New(fmt.Sprint("error while delete user : ", er.Error()))
		}
	}

	return nil
}

func (s *SysUserMysqlImpl) ResetUserPassword(user entities.SysDaftarUser) (sysUser entities.SysDaftarUser, er error) {

	thisRepo, _ := NewSysUserRepo()
	_, er = thisRepo.GetSingleUserByUserName(user.User_Name)
	if er != nil {
		return sysUser, err.NoRecord
	}

	stmt, er := s.apexDb.Prepare("UPDATE sys_daftar_user SET user_web_password = ? WHERE user_name = ?")
	if er != nil {
		return sysUser, errors.New(fmt.Sprint("error while prepare update apex password: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(user.User_Web_Password, user.User_Name); er != nil {
		return sysUser, errors.New(fmt.Sprint("error while update apex password: ", er.Error()))
	}

	return user, nil
}

func (s *SysUserMysqlImpl) FindByUserName(userName string) (user entities.SysDaftarUser, er error) {
	row := s.apexDb.QueryRow(`SELECT
		user_id, 
		user_name,
		nama_lengkap,
		DATE_FORMAT(tgl_expired, "%d/%m/%Y") AS tgl_expired,
		user_web_password
	FROM sys_daftar_user WHERE user_name = ? LIMIT 1`, userName)
	er = row.Scan(
		&user.User_Id,
		&user.User_Name,
		&user.Nama_Lengkap,
		&constant.SQLTglExpired,
		&user.User_Web_Password_Hash,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return user, err.NoRecord
		} else {
			return user, errors.New(fmt.Sprint("error while get user name: ", er.Error()))
		}
	}

	user.TglExpiredStr = constant.SQLTglExpired.String
	return
}

func (s *SysUserMysqlImpl) UpdateLKMName(updSysuser entities.SysDaftarUser) (er error) {

	thisRepo, _ := NewSysUserRepo()
	user, er := thisRepo.GetSingleUserByUserName(updSysuser.User_Name)
	if er != nil {
		return err.NoRecord
	}

	stmt, er := s.apexDb.Prepare(`UPDATE sys_daftar_user SET nama_lengkap = ? WHERE user_name = ?`)
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare update lkm name: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(updSysuser.Nama_Lengkap, updSysuser.User_Name); er != nil {
		return errors.New(fmt.Sprint("error while update lkm name: ", er.Error()))
	}

	updSysuser.User_Id = user.User_ID

	return nil

}

func (s *SysUserMysqlImpl) GetListsOtorisator() (lists []entities.Otorisators, er error) {

	item := strings.Split(os.Getenv("app.list_otorisator_id"), ",")
	listOtorisatorID := (strings.Join(item, ","))

	rows, er := s.apexDb.Query(`SELECT user_id, nama_lengkap, jabatan FROM sys_daftar_user WHERE user_id IN (` + listOtorisatorID + `) ORDER BY nama_lengkap ASC`)
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var otorisator entities.Otorisators
		if er = rows.Scan(&otorisator.UserID, &otorisator.NamaOtorisator, &otorisator.Jabatan); er != nil {
			return lists, er
		}

		lists = append(lists, otorisator)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}
