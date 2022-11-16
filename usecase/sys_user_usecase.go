package usecase

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/helper"
	"apex-ems-integration-clean-arch/repository/sysuserrepo"
)

type SysUserUsecase interface {
	Login(input web.LoginInput) (user web.LoginData, er error)
	ResetSysUserPassword(KodeLkm web.KodeLKMFilter) (web.ResetApexPwdResponse, error)
}

type sysUserUsecase struct{} // (e *employeeUsecase) => untuk menentukan hak kepemilikan

func NewSysUserUsecase() SysUserUsecase {
	return &sysUserUsecase{}
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
