package constants

// KEY:VAL MAPPING
/*
	nasabah.Nasabah_Id = payload.KodeLkm
	nasabah.Nama_Nasabah = payload.Nama_Lembaga
	nasabah.Alamat = payload.Alamat
	nasabah.Telepon = payload.Telepon
	nasabah.Jenis_Kelamin = "L"
	nasabah.TempatLahir = "Bandung"
	nasabah.TglLahir = time.Now()
	nasabah.Jenis_Id = "1"
	nasabah.No_Id = helper.GenerateIdKTP()
	nasabah.Kode_Group1 = "1"
	nasabah.Kode_Group2 = "01"
	nasabah.Kode_Group3 = "001"
	nasabah.Kode_Agama = "Islam"
	nasabah.Desa = "Desa"
	nasabah.Kecamatan = "Kecamatan"
	nasabah.Kota_Kab = "0121"
	nasabah.Provinsi = "008"
	nasabah.Verifikasi = "1"
	nasabah.Hp = payload.Telepon
	nasabah.Tgl_Register = time.Now()
	nasabah.Nama_Ibu_Kandung = "Ibu"
	nasabah.Kodepos = "12345"
	nasabah.Kode_Kantor = "001"
	nasabah.UserId = payload.User_Id
	nasabah.Nama_Alias = payload.Nama_Lembaga
	nasabah.Status_Gelar = "0100"
	nasabah.Jenis_Debitur = "0"
	nasabah.Kode_Area = "23"
	nasabah.Negara_Domisili = "ID"
	nasabah.Gol_Debitur = "907"
	nasabah.Langgar_Bmpk = "T"
	nasabah.Lampaui_Bmpk = "T"
	nasabah.Nama_Nasabah_Sid = payload.Nama_Lembaga
	nasabah.Alamat2 = payload.Alamat
	nasabah.Flag_Masa_Berlaku = "1"
	nasabah.Status_Marital = "Single"
	nasabah.Hp1 = payload.Telepon
	nasabah.Hp2 = payload.Telepon
	nasabah.Status_Tempat_Tinggal = "Milik Sendiri"
	nasabah.Masa_Berlaku_Ktp = time.Now().AddDate(7, 0, 0)

	tabung.No_Rekening = payload.KodeLkm
	tabung.Nasabah_Id = payload.KodeLkm
	tabung.Kode_Bi_Pemilik = "874"
	tabung.Suku_Bunga = 0
	tabung.Persen_Pph = 0
	tabung.Tgl_Register = time.Now()
	tabung.Saldo_Akhir = 0
	tabung.Kode_Group1 = "001"
	tabung.Kode_Group2 = "01"
	tabung.Verifikasi = "1"
	tabung.Status = 1
	tabung.Kode_Kantor = "001"
	tabung.Kode_Integrasi = "01"
	tabung.Kode_Produk = "01"
	tabung.UserId = payload.User_Id
	tabung.Kode_Group3 = ""
	tabung.Minimum = 0
	tabung.Setoran_Minimum = 0
	tabung.Jkw = 0
	tabung.Abp = 0
	tabung.Setoran_Wajib = 0
	tabung.Adm_Per_Bln = 0
	tabung.Target_Nominal = 0
	tabung.Saldo_Akhir_Titipan_bunga = 0
	tabung.Kode_Bi_Lokasi = "1"
	tabung.Saldo_Akhir_Titipan_bunga = 0
	tabung.Saldo_Titipan_Bunga_Ks = 0
	tabung.Saldo_Blokir = 0
	tabung.Premi = 0
	tabung.Kode_Keterkaitan = "2"
	tabung.Kode_Kantor_Kas = "01"
	tabung.No_Rekening_Virtual = payload.KodeLkm

	sysDaftarUser.User_Name = payload.KodeLkm
	sysDaftarUser.User_Password = "TKkRamfizZc="
	sysDaftarUser.Nama_Lengkap = payload.Nama_Lembaga
	sysDaftarUser.Unit_Kerja = "001"
	sysDaftarUser.Jabatan = "Echannel"
	sysDaftarUser.User_Code = "1"
	sysDaftarUser.Tgl_Expired = time.Now().AddDate(7, 0, 0)
	sysDaftarUser.User_Web_Password_Hash, sysDaftarUser.User_Web_Password = helper.HashSha1Pass()
	sysDaftarUser.Flag = 1
	sysDaftarUser.Status_Aktif = 1
	sysDaftarUser.Penerimaan = 0
	sysDaftarUser.Pengeluaran = 0

*/
const (

	// Global Reusable var
	ZeroValInt  = 0
	KodeKantor  = "001"
	Verifikasi  = "1"
	StatusAktif = 1
	EmptyStr    = ""

	// Static value untuk nasabah
	JenisKelamin = "L"
	TempatLahir  = "Bandung"
	JenisId      = "1"
	Group1       = "1"
	Group2       = "01"
	Group3       = "001"
	Agama        = "Islam"
	Desa         = "Desa"
	Kecamatan    = "Kecamatan"
	KotaKabKode  = "0121"
	Prov         = "008"
	NamaIbu      = "Ibu"
	KodePos      = "12345"

	StatusGelar         = "0100"
	JenisDebitur        = "0"
	KodeArea            = "23"
	NegaraDomisili      = "ID"
	GolDebitur          = "907"
	LampauLanggarBmpk   = "T"
	FlagMasaBerlaku     = "1"
	StatusMarital       = "Single"
	StatusTempatTinggal = "Milik Sendiri"

	// Static value untuk tabung
	KodeBIPemilik         = "874"
	KdGroup1              = "001"
	KdGroup2              = "01"
	KdIntegrasi           = "01"
	KdProduk              = "01"
	KdBILokasi            = "1"
	KdKeterkaitan         = "2"
	KdKantorKas           = "01"
	FlagPayEchannelActive = "Y"

	// Static value untuk sys_daftar_user
	UserPwd   = "TKkRamfizZc"
	UnitKerja = "001"
	Jabatan   = "Echannel"
	Flag      = 1
	UserCode  = "1"

	// Static value untuk tabtrans
	Kredit        = "100"
	Debit         = "200"
	TransIDSource = 0
	CounterSign   = 0
	PostedToGl    = 1

	// Static value untuk approval
	NeedAprove = 0
	Approved   = 1
	Rejected   = 2
	Used       = 3
)
