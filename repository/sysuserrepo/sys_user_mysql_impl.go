package sysuserrepo

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"database/sql"
	"errors"
	"fmt"
)

func newSysUserMysqlImpl(apexConn *sql.DB) SysUserRepo {
	return &SysUserMysqlImpl{
		apexDb: apexConn,
	}
}

type SysUserMysqlImpl struct {
	apexDb *sql.DB
}

func (s *SysUserMysqlImpl) GetSingleUserByUserName(userName string) (user web.ManajemenUserDataResponse, er error) {
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
		&user.Jabatan,
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
	return
}

func (s *SysUserMysqlImpl) GetListOfUsers(limitOffset web.LimitOffsetLkmUri) (lists []web.ManajemenUserDataResponse, er error) {
	args := []interface{}{}
	limit := ""
	if limitOffset.Limit > 0 {
		limit = "LIMIT ? OFFSET ?"
		args = append(args, limitOffset.Limit, limitOffset.Offset)
	} else {
		limit = "LIMIT ? OFFSET ?"
		args = append(args, -1, limitOffset.Offset)
	}

	rows, er := s.apexDb.Query(`SELECT 
		user_id,
		user_name,
		nama_lengkap,
		jabatan,
		unit_kerja,
		COALESCE(DATE_FORMAT(TGL_EXPIRED, "%d/%m/%Y"),'') AS tgl_expired,
		status_aktif,
		user_code
		FROM sys_daftar_user `+limit+``, args...)
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var users web.ManajemenUserDataResponse
		if er = rows.Scan(
			&users.User_ID,
			&users.User_Name,
			&users.Nama_Lengkap,
			&users.Jabatan,
			&users.Unit_Kerja,
			&users.Tgl_Expired,
			&users.StatusAktif,
			&users.User_Code,
		); er != nil {
			return lists, er
		}

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
	pengeluaran =?
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

func (s *SysUserMysqlImpl) HardDeleteSysDaftarUser(kodeLkm string) (er error) {

	stmt, er := s.apexDb.Prepare("DELETE FROM sys_daftar_user WHERE user_name = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete user : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete user : ", er.Error()))
	}

	return nil
}

func (s *SysUserMysqlImpl) DeleteSysDaftarUser(kodeLkm string) (er error) {

	thisRepo, _ := NewSysUserRepo()
	_, er = thisRepo.GetSingleUserByUserName(kodeLkm)
	if er != nil {
		return err.NoRecord
	}

	stmt, er := s.apexDb.Prepare("UPDATE sys_daftar_user SET status_aktif = 0, user_name = ? WHERE user_name = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete user : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec("DEL-"+kodeLkm, kodeLkm); er != nil {
		return errors.New(fmt.Sprint("error while delete user : ", er.Error()))
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
		tgl_expired,
		user_web_password
	FROM sys_daftar_user WHERE user_name = ? LIMIT 1`, userName)
	er = row.Scan(
		&user.User_Id,
		&user.User_Name,
		&user.Nama_Lengkap,
		&user.Tgl_Expired,
		&user.User_Web_Password_Hash,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return user, err.NoRecord
		} else {
			return user, errors.New(fmt.Sprint("error while get user name: ", er.Error()))
		}
	}
	return
}
