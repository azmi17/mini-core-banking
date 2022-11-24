package usecase

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/constants"
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/helper"
	"apex-ems-integration-clean-arch/repository/sysuserrepo"
	"time"
)

type SysUserUsecase interface {
	CreateSysUser(input web.CreateManajemenUser) (user web.CreateManajemenUserDataResponse, er error)
	UpdateSysUser(input web.UpdateManajemenUser) (user web.UpdateManajemenUserDataResponse, er error)
	GetSingleUserByUserName(userName string) (web.ManajemenUserDataResponse, error)
	GetListOfUsers(limitOffset web.LimitOffsetLkmUri) ([]web.ManajemenUserDataResponse, error)
	Login(input web.LoginInput) (user web.LoginData, er error)
	ResetSysUserPassword(KodeLkm web.KodeLKMFilter) (web.ResetApexPwdResponse, error)
}

type sysUserUsecase struct{} // (e *employeeUsecase) => untuk menentukan hak kepemilikan

func NewSysUserUsecase() SysUserUsecase {
	return &sysUserUsecase{}
}

func (s *sysUserUsecase) CreateSysUser(input web.CreateManajemenUser) (user web.CreateManajemenUserDataResponse, er error) {
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	sysDaftarUser := entities.SysDaftarUser{}

	// From payload:
	sysDaftarUser.User_Name = input.UserName
	sysDaftarUser.Nama_Lengkap = input.NamaUser

	// Set static Val:
	sysDaftarUser.User_Password = constants.UserPwd
	sysDaftarUser.Unit_Kerja = constants.UnitKerja
	sysDaftarUser.Jabatan = input.Jabatan
	sysDaftarUser.User_Code = input.UserCode
	sysDaftarUser.Tgl_Expired = time.Now().AddDate(15, 0, 0)
	sysDaftarUser.User_Web_Password_Hash, sysDaftarUser.User_Web_Password = helper.HashSha1Pass()
	sysDaftarUser.Flag = constants.Flag
	sysDaftarUser.Status_Aktif = constants.StatusAktif
	sysDaftarUser.Penerimaan = constants.ZeroValInt
	sysDaftarUser.Pengeluaran = constants.ZeroValInt

	if sysDaftarUser, er = sysUserRepo.CreateSysDaftarUser(sysDaftarUser); er != nil {
		return user, er
	}

	user.User_ID = sysDaftarUser.User_Id
	user.User_Name = sysDaftarUser.User_Name
	user.Nama_Lengkap = sysDaftarUser.Nama_Lengkap
	user.Password = sysDaftarUser.User_Web_Password_Hash
	user.Jabatan = sysDaftarUser.Jabatan
	user.Unit_Kerja = sysDaftarUser.Unit_Kerja
	user.Tgl_Expired = string(sysDaftarUser.Tgl_Expired.Format("02-01-2006"))
	user.StatusAktif = sysDaftarUser.Status_Aktif
	user.User_Code = sysDaftarUser.User_Code

	return user, nil
}

func (s *sysUserUsecase) UpdateSysUser(payload web.UpdateManajemenUser) (user web.UpdateManajemenUserDataResponse, er error) {
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	sysDaftarUser := entities.SysDaftarUser{
		Nama_Lengkap: payload.NamaUser,
		Jabatan:      payload.Jabatan,
		Status_Aktif: payload.StatusAktif,
		User_Code:    payload.UserCode,
		User_Name:    payload.UserName,
		Tgl_Expired:  time.Now().AddDate(15, 0, 0),
		Unit_Kerja:   constants.UnitKerja,
		Flag:         constants.Flag,
		Penerimaan:   constants.ZeroValInt,
		Pengeluaran:  constants.ZeroValInt,
	}
	if sysDaftarUser, er = sysUserRepo.UpdateSysDaftarUser(sysDaftarUser); er != nil {
		return user, er
	}

	user.User_ID = sysDaftarUser.User_Id
	user.User_Name = sysDaftarUser.User_Name
	user.Nama_Lengkap = sysDaftarUser.Nama_Lengkap
	user.Jabatan = sysDaftarUser.Jabatan
	user.Unit_Kerja = sysDaftarUser.Unit_Kerja
	user.Tgl_Expired = string(sysDaftarUser.Tgl_Expired.Format("02-01-2006"))
	user.StatusAktif = sysDaftarUser.Status_Aktif
	user.User_Code = sysDaftarUser.User_Code

	return user, nil
}

func (s *sysUserUsecase) GetSingleUserByUserName(userName string) (routingInfo web.ManajemenUserDataResponse, er error) {
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	if routingInfo, er = sysUserRepo.GetSingleUserByUserName(userName); er != nil {
		return routingInfo, er
	}

	return routingInfo, nil
}

func (s *sysUserUsecase) GetListOfUsers(limitOffset web.LimitOffsetLkmUri) (routingList []web.ManajemenUserDataResponse, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return routingList, err.BadRequest
	}

	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()
	routingList, er = sysUserRepo.GetListOfUsers(limitOffset)
	if er != nil {
		return routingList, er
	}

	if len(routingList) == 0 {
		return routingList, err.NoRecord
	}

	return routingList, nil
}

func (s *sysUserUsecase) Login(input web.LoginInput) (user web.LoginData, er error) {
	userName := input.User_Name
	password := input.Password

	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()
	userData := entities.SysDaftarUser{}
	if userData, er = sysUserRepo.FindByUserName(userName); er != nil {
		return user, er
	}

	hashPass := helper.HashSha1PassByInput(password)
	if userData.User_Web_Password_Hash != hashPass {
		return user, err.PasswordDontMatch
	}

	return helper.SysUserLoginResponseFilter(userData), nil

}

func (s *sysUserUsecase) ResetSysUserPassword(KodeLkm web.KodeLKMFilter) (resp web.ResetApexPwdResponse, er error) {
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	sysDaftarUser := entities.SysDaftarUser{}
	sysDaftarUser.User_Name = KodeLkm.KodeLkm
	sysDaftarUser.User_Web_Password_Hash, sysDaftarUser.User_Web_Password = helper.HashSha1Pass()

	if sysDaftarUser, er = sysUserRepo.ResetUserPassword(sysDaftarUser); er != nil {
		return resp, er
	}

	updResp := web.ResetApexPwdResponse{}
	updResp.KodeLkm = sysDaftarUser.User_Name
	updResp.Password_Smec = sysDaftarUser.User_Web_Password_Hash

	return updResp, nil
}
