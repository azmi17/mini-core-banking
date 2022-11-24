package usecase

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/constants"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/helper"
	"apex-ems-integration-clean-arch/repository/nasabahrepo"
	"apex-ems-integration-clean-arch/repository/sysuserrepo"
	"apex-ems-integration-clean-arch/repository/tabunganrepo"
	"time"
)

type LkmUsecase interface {
	CreateLkm(payload web.SaveLKMApex) (web.LKMCreateResponse, error)
	UpdateLkm(payload web.SaveLKMApex) (web.LKMUpdateResponse, error)
	HardDeleteLkm(kodeLkm string) error
	DeleteLkm(kodeLkm string) error
}

type lkmUsecase struct{} // (e *employeeUsecase) => untuk menentukan hak kepemilikan

func NewLkmUsecase() LkmUsecase {
	return &lkmUsecase{}
}

func (lkm *lkmUsecase) CreateLkm(payload web.SaveLKMApex) (createLkm web.LKMCreateResponse, er error) {
	nasabahRepo, _ := nasabahrepo.NewNasabahRepo()
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	nasabah := entities.Nasabah{}

	// From payload:
	nasabah.Nasabah_Id = payload.KodeLkm
	nasabah.Nama_Nasabah = payload.Nama_Lembaga
	nasabah.Alamat = payload.Alamat
	nasabah.Telepon = payload.Telepon
	nasabah.Hp = payload.Telepon
	nasabah.UserId = payload.User_Id
	nasabah.Nama_Alias = payload.Nama_Lembaga
	nasabah.Nama_Nasabah_Sid = payload.Nama_Lembaga
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
	nasabah.Kode_Group2 = constants.Group2
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
	nasabah.Masa_Berlaku_Ktp = time.Now().AddDate(7, 0, 0)
	if nasabah, er = nasabahRepo.CreateNasabah(nasabah); er != nil {
		return createLkm, er
	}

	tabung := entities.Tabung{}

	// From payload:
	tabung.No_Rekening = payload.KodeLkm
	tabung.Nasabah_Id = payload.KodeLkm
	tabung.UserId = payload.User_Id
	tabung.No_Rekening_Virtual = payload.KodeLkm

	// Set static Val:
	tabung.Kode_Bi_Pemilik = constants.KodeBIPemilik
	tabung.Suku_Bunga = constants.ZeroValInt
	tabung.Persen_Pph = constants.ZeroValInt
	tabung.Tgl_Register = time.Now()
	tabung.Saldo_Akhir = constants.ZeroValInt
	tabung.Kode_Group1 = constants.KdGroup1
	tabung.Kode_Group2 = constants.KdGroup2
	tabung.Verifikasi = constants.Verifikasi
	tabung.Status = constants.StatusAktif
	tabung.Kode_Kantor = constants.KodeKantor
	tabung.Kode_Integrasi = constants.KdIntegrasi
	tabung.Kode_Produk = constants.KdProduk
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
	if tabung, er = tabunganRepo.CreateTabung(tabung); er != nil {
		return createLkm, er
	}

	sysDaftarUser := entities.SysDaftarUser{}

	// From payload:
	sysDaftarUser.User_Name = payload.KodeLkm
	sysDaftarUser.Nama_Lengkap = payload.Nama_Lembaga

	// Set static Val:
	sysDaftarUser.User_Password = constants.UserPwd
	sysDaftarUser.Unit_Kerja = constants.UnitKerja
	sysDaftarUser.Jabatan = constants.Jabatan
	sysDaftarUser.User_Code = constants.UserCode
	sysDaftarUser.Tgl_Expired = time.Now().AddDate(7, 0, 0)
	sysDaftarUser.User_Web_Password_Hash, sysDaftarUser.User_Web_Password = helper.HashSha1Pass()
	sysDaftarUser.Flag = constants.Flag
	sysDaftarUser.Status_Aktif = constants.StatusAktif
	sysDaftarUser.Penerimaan = constants.ZeroValInt
	sysDaftarUser.Pengeluaran = constants.ZeroValInt
	if sysDaftarUser, er = sysUserRepo.CreateSysDaftarUser(sysDaftarUser); er != nil {
		return createLkm, er
	}

	//Converting data => 3 repo to ApexResponse
	return helper.ApexFilterLKMResponse(nasabah, tabung, sysDaftarUser), nil
}

func (lkm *lkmUsecase) UpdateLkm(payload web.SaveLKMApex) (updateLkm web.LKMUpdateResponse, er error) {
	nasabahRepo, _ := nasabahrepo.NewNasabahRepo()
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	nasabah := entities.Nasabah{}
	nasabah.Nama_Nasabah = payload.Nama_Lembaga
	nasabah.Nama_Alias = payload.Nama_Lembaga
	nasabah.Nama_Nasabah_Sid = payload.Nama_Lembaga
	nasabah.Alamat = payload.Alamat
	nasabah.Alamat2 = payload.Alamat
	nasabah.Telepon = payload.Telepon
	nasabah.Hp = payload.Telepon
	nasabah.Hp1 = payload.Telepon
	nasabah.Hp2 = payload.Telepon
	nasabah.UserId = payload.User_Id
	nasabah.Nasabah_Id = payload.KodeLkm
	if nasabah, er = nasabahRepo.UpdateNasabah(nasabah); er != nil {
		return updateLkm, er
	}

	sysDaftarUser := entities.SysDaftarUser{
		Nama_Lengkap: payload.Nama_Lembaga,
		User_Name:    payload.KodeLkm,
	}
	if sysDaftarUser, er = sysUserRepo.UpdateSysDaftarUser(sysDaftarUser); er != nil {
		return updateLkm, er
	}

	return helper.ApexUpdateLKMResponse(nasabah, sysDaftarUser), nil

}

func (lkm *lkmUsecase) DeleteLkm(kodeLkm string) (er error) {
	nasabahRepo, _ := nasabahrepo.NewNasabahRepo()
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	if er = nasabahRepo.DeleteNasabah(kodeLkm); er != nil {
		return er
	}

	if er = tabunganRepo.DeleteTabung(kodeLkm); er != nil {
		return er
	}

	if er = sysUserRepo.DeleteSysDaftarUser(kodeLkm); er != nil {
		return er
	}

	return nil

}

func (lkm *lkmUsecase) HardDeleteLkm(kodeLkm string) (er error) {
	nasabahRepo, _ := nasabahrepo.NewNasabahRepo()
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	if er = nasabahRepo.HardDeleteNasabah(kodeLkm); er != nil {
		return er
	}

	if er = tabunganRepo.HardDeleteTabung(kodeLkm); er != nil {
		return er
	}

	if er = sysUserRepo.HardDeleteSysDaftarUser(kodeLkm); er != nil {
		return er
	}

	return nil

}
