package usecase

import (
	"new-apex-api/entities"
	"new-apex-api/entities/constants"
	"new-apex-api/entities/err"
	"new-apex-api/helper"
	"new-apex-api/repository/nasabahrepo"
	"new-apex-api/repository/sysuserrepo"
	"new-apex-api/repository/tabtransrepo"
	"new-apex-api/repository/tabunganrepo"
	"time"
)

type LkmUsecase interface {
	GetLKMInfoLists(payload entities.GlobalFilter, limitOffset entities.LimitOffsetLkmUri) ([]entities.GetDetailLKMInfo, error)
	GetLKMDetailInfo(Id string) (entities.GetDetailLKMInfo, error)
	CreateLkm(payload entities.CreateLKM) (entities.LKMCreateResponse, error)
	UpdateLkm(payload entities.UpdateLKM) (entities.LKMUpdateResponse, error)
	HardDeleteLkm(kodeLkm []string) error
	DeleteLkm(kodeLkm []string) error
}

type lkmUsecase struct{}

func NewLkmUsecase() LkmUsecase {
	return &lkmUsecase{}
}

func (lkm *lkmUsecase) GetLKMDetailInfo(Id string) (detailTabInfo entities.GetDetailLKMInfo, er error) {
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()

	if detailTabInfo, er = tabunganRepo.GetTabDetailInfo(Id); er != nil {
		return detailTabInfo, er
	}

	return detailTabInfo, nil
}

func (lkm *lkmUsecase) GetLKMInfoLists(payload entities.GlobalFilter, limitOffset entities.LimitOffsetLkmUri) (lkmTabList []entities.GetDetailLKMInfo, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return lkmTabList, err.BadRequest
	}

	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()

	lkmTabList, er = tabunganRepo.GetTabInfoList(payload, limitOffset)
	if er != nil {
		return lkmTabList, er
	}

	if len(lkmTabList) == 0 {
		return make([]entities.GetDetailLKMInfo, 0), nil
		//^ []entities.GetDetailLKMInfo (untuk membuat sebuah slice kosong agar tidak return null di JSON) |err.NoRecord
	}

	return lkmTabList, nil
}

func (lkm *lkmUsecase) CreateLkm(payload entities.CreateLKM) (createLkm entities.LKMCreateResponse, er error) {
	nasabahRepo, _ := nasabahrepo.NewNasabahRepo()
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	nasabah := entities.Nasabah{}

	// From payload:
	nasabah.Nasabah_Id = payload.KodeLkm
	nasabah.Nama_Nasabah = payload.NamaLembaga
	nasabah.Kode_Group2 = payload.KodeSC
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
		return createLkm, er
	}

	tabung := entities.Tabung{}

	// From payload:
	tabung.No_Rekening = payload.KodeLkm
	tabung.Nasabah_Id = payload.KodeLkm
	tabung.IsSaldoChecked = payload.IsSaloChecked
	tabung.UserId = payload.UserID
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
	tabung.Minimum = float32(payload.Plafond)
	tabung.Setoran_Minimum = float32(payload.SetoranMinimum)
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
	tabung.FlagPayEchannel = constants.FlagPayEchannelActive
	if tabung, er = tabunganRepo.CreateTabung(tabung); er != nil {
		return createLkm, er
	}

	sysDaftarUser := entities.SysDaftarUser{}

	// From payload:
	sysDaftarUser.User_Name = payload.KodeLkm
	sysDaftarUser.Nama_Lengkap = payload.NamaLembaga

	// Set static Val:
	sysDaftarUser.User_Password = constants.UserPwd
	sysDaftarUser.Unit_Kerja = constants.UnitKerja
	sysDaftarUser.Jabatan = constants.Jabatan
	sysDaftarUser.User_Code = constants.UserCode
	sysDaftarUser.Tgl_Expired = time.Now().AddDate(15, 0, 0)
	sysDaftarUser.User_Web_Password_Hash, sysDaftarUser.User_Web_Password = helper.HashSha1Pass()
	sysDaftarUser.Flag = constants.Flag
	sysDaftarUser.Status_Aktif = constants.StatusAktif
	sysDaftarUser.Penerimaan = constants.ZeroValInt
	sysDaftarUser.Pengeluaran = constants.ZeroValInt
	if sysDaftarUser, er = sysUserRepo.CreateSysDaftarUser(sysDaftarUser); er != nil {
		return createLkm, er
	}

	//Converting data => 3 repo to ApexResponse
	return helper.ApexCreateLKMResponse(nasabah, tabung, sysDaftarUser), nil
}

func (lkm *lkmUsecase) UpdateLkm(payload entities.UpdateLKM) (updateLkm entities.LKMUpdateResponse, er error) {
	nasabahRepo, _ := nasabahrepo.NewNasabahRepo()
	tabungRepo, _ := tabunganrepo.NewTabunganRepo()
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	nasabah := entities.Nasabah{}
	nasabah.Nama_Nasabah = payload.NamaLembaga
	nasabah.Kode_Group2 = payload.KodeSC
	nasabah.Nama_Alias = payload.NamaLembaga
	nasabah.Nama_Nasabah_Sid = payload.NamaLembaga
	nasabah.Alamat = payload.Alamat
	nasabah.Alamat2 = payload.Alamat
	nasabah.Telepon = payload.Telepon
	nasabah.Hp = payload.Telepon
	nasabah.Hp1 = payload.Telepon
	nasabah.Hp2 = payload.Telepon
	nasabah.UserId = payload.UserID
	nasabah.Nasabah_Id = payload.KodeLkm
	if nasabah, er = nasabahRepo.UpdateNasabah(nasabah); er != nil {
		return updateLkm, er
	}

	sysDaftarUser := entities.SysDaftarUser{
		Nama_Lengkap: payload.NamaLembaga,
		User_Name:    payload.KodeLkm,
	}
	if er = sysUserRepo.UpdateLKMName(sysDaftarUser); er != nil {
		return updateLkm, er
	}

	tabungan := entities.UpdateRekeningLKM{
		KodeLKM:        payload.KodeLkm,
		KodeSC:         payload.KodeLkm,
		Status:         payload.Status,
		Plafond:        payload.Plafond,
		SetoranMinimum: payload.SetoranMinimum,
		IsSaldoChecked: payload.IsSaldoChecked,
		UserID:         payload.UserID,
	}
	_, er = tabungRepo.EditRekeningLKM(tabungan)
	if er != nil {
		return updateLkm, er
	}

	return helper.ApexUpdateLKMResponse(nasabah, tabungan, sysDaftarUser), nil

}

// Harus multiple atau tidak (integrasi EMS)
func (lkm *lkmUsecase) DeleteLkm(kodeLkm []string) (er error) {
	nasabahRepo, _ := nasabahrepo.NewNasabahRepo()
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	if er = nasabahRepo.SoftDeleteNasabah(kodeLkm...); er != nil {
		return er
	}

	if er = tabunganRepo.SoftDeleteTabung(kodeLkm...); er != nil {
		return er
	}

	if er = tabtransRepo.SoftDeleteApexTransaction(kodeLkm...); er != nil {
		return er
	}

	if er = sysUserRepo.SoftDeleteSysDaftarUser(kodeLkm...); er != nil {
		return er
	}

	return nil

}

func (lkm *lkmUsecase) HardDeleteLkm(kodeLkm []string) (er error) {
	nasabahRepo, _ := nasabahrepo.NewNasabahRepo()
	tabunganRepo, _ := tabunganrepo.NewTabunganRepo()
	sysUserRepo, _ := sysuserrepo.NewSysUserRepo()

	if er = tabunganRepo.HardDeleteTabung(kodeLkm...); er != nil {
		return er
	}

	if er = nasabahRepo.HardDeleteNasabah(kodeLkm...); er != nil {
		return er
	}

	if er = sysUserRepo.HardDeleteSysDaftarUser(kodeLkm...); er != nil {
		return er
	}

	return nil

}
