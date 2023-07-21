package usecase

import (
	"new-apex-api/entities"
	"new-apex-api/entities/constants"
	"new-apex-api/entities/err"
	"new-apex-api/helper"
	"new-apex-api/repository/informationlkmrepo"
	"new-apex-api/repository/tabtransrepo"
	"time"
)

type TabtransUsecase interface {
	GetListsTabtransTransaction(payload entities.GetListTabtrans, limitOffset entities.LimitOffsetLkmUri) (data []entities.GetListTabtransTrx, er error)
	GetListsTabtransTrxBySTAN(stan string) ([]entities.GetListTabtransTrx, error)
	HardDeleteApexTransaction(data []int) error
	ChangeDateOnTabtransTrx(tabtransID int, tglTrans string) (entities.GetListTabtransTrx, error)
	MultipleChangeDateOnApexTransaction(payload entities.MultipleChangeDateTransaction) error
	GetRekeningKoranLKMDetail(entities.RekeningKoranRequest) (data entities.RekeningKoranResponse, er error)
	GetNominatifDeposit(payload entities.NominatifDepositRequest, limitOffset entities.LimitOffsetLkmUri) (data []entities.NominatifDepositResponse, er error)
	GetLaporanTransaksi(payload entities.DaftarTransaksiRequest, limitOffset entities.LimitOffsetLkmUri) (data []entities.DaftarTransaksiResponse, er error)
	GetListsTransaksiDeposit(payload entities.GetListsDepositTrxReq) ([]entities.GetListsDepositTrxRes, error)
	TransaksiDeposit(payload entities.DepositRequest) (er error)
	PembatalanTransaksiDeposit(payload entities.ReversalDepositRequest) (er error)
}

type tabtransUsecase struct{}

func NewTabtransUsecase() TabtransUsecase {
	return &tabtransUsecase{}
}

func (t *tabtransUsecase) GetListsTabtransTransaction(payload entities.GetListTabtrans, limitOffset entities.LimitOffsetLkmUri) (tabtransTxList []entities.GetListTabtransTrx, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return tabtransTxList, err.BadRequest
	}

	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
	tabtransTxList, er = tabtransRepo.GetListsApexTransaction(payload, limitOffset)
	if er != nil {
		return tabtransTxList, er
	}

	if len(tabtransTxList) == 0 {
		return make([]entities.GetListTabtransTrx, 0), nil
		// ^ []entities.GetDetailLKMInfo (untuk membuat sebuah slice kosong agar tidak return null di JSON) | err.NoRecord
	}

	return tabtransTxList, nil
}

func (t *tabtransUsecase) GetListsTabtransTrxBySTAN(stan string) (tx []entities.GetListTabtransTrx, er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	tx, er = tabtransRepo.GetListsApexTransactionBySTAN(stan)
	if er != nil {
		return tx, er
	}

	if len(tx) == 0 {
		return tx, err.NoRecord
	}

	return tx, nil
}

func (t *tabtransUsecase) HardDeleteApexTransaction(data []int) (er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	if er = tabtransRepo.HardDeleteApexTransaction(data...); er != nil {
		return er
	}

	return nil
}

func (t *tabtransUsecase) ChangeDateOnTabtransTrx(tabtransID int, tglTrans string) (data entities.GetListTabtransTrx, er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	if data, er = tabtransRepo.ChangeDateOnApexTransaction(tabtransID, tglTrans); er != nil {
		return data, er
	}

	return
}

func (t *tabtransUsecase) MultipleChangeDateOnApexTransaction(payload entities.MultipleChangeDateTransaction) (er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	if er = tabtransRepo.MultipleChangeDateOnApexTransaction(payload); er != nil {
		return er
	}

	return
}

func (t *tabtransUsecase) GetRekeningKoranLKMDetail(payload entities.RekeningKoranRequest) (resp entities.RekeningKoranResponse, er error) {
	// Add separator on date payload
	payloadPeriodeAwal, _ := helper.AddSeparatorsOnDateStr(payload.PeriodeAwal)
	payloadPeriodeAkhir, _ := helper.AddSeparatorsOnDateStr(payload.PeriodeAkhir)

	periodeAwalFormatStr, _ := helper.FormatTimeStrDDMMYYY(payloadPeriodeAwal)
	periodeAkhirFormatStr, _ := helper.FormatTimeStrDDMMYYY(payloadPeriodeAkhir)

	informationLKMrepo, _ := informationlkmrepo.NewInformationLKMRepo()
	header, er := informationLKMrepo.RekeningKoranLKMDetailHeader(payload.KodeLKM)
	if er != nil {
		return resp, er
	}
	resp.Norek = header.Norek
	resp.NamaLembaga = header.NamaLembaga
	resp.ProdukTab = header.ProdukTab
	resp.NamaSC = header.NamaSC
	resp.PeriodeAwal = periodeAwalFormatStr
	resp.PeriodeAkhir = periodeAkhirFormatStr

	periodeAwal, _ := helper.ParseTimeStrToDateOnYesterday(payloadPeriodeAkhir)
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
	saldoAwal, er := tabtransRepo.GetSaldo(payload.KodeLKM, periodeAwal) // tanggal periode awal harus di substract
	if er != nil {
		return resp, er
	}
	resp.SaldoAwal = saldoAwal

	// disini parsing dd/mm/yyyy ke yyyymmdd
	data, er := tabtransRepo.GetRekeningKoranLKMDetail(payload.KodeLKM, payload.PeriodeAwal, payload.PeriodeAkhir)
	if er != nil {
		return resp, nil
	}

	saldoTemp := saldoAwal
	resp.Detail = make([]entities.RekeningKoranDetail, 0) // default 0 for detail field
	for _, v := range data {
		var detail entities.RekeningKoranDetail
		if v.MyKodeTrans == "200" {
			detail.Debet = v.Pokok
		} else {
			detail.Kredit = v.Pokok
		}
		saldoTemp = saldoTemp + detail.Kredit - detail.Debet // kalkulasi mutasi akhir saldo
		detail.TglTrans = v.TglTrans
		detail.Uraian = v.Keterangan
		detail.KodeTrans = v.KodeTrans
		detail.Saldo = saldoTemp
		detail.Kuitansi = v.Kuitansi
		detail.NorekLKM = v.PayLkmNorek
		detail.Idpel = v.PayIdpel
		detail.Biller = v.PayBillerCode
		detail.Produk = v.PayProductCode
		resp.Detail = append(resp.Detail, detail)
	}
	return resp, nil
}

func (t *tabtransUsecase) GetNominatifDeposit(payload entities.NominatifDepositRequest, limitOffset entities.LimitOffsetLkmUri) (data []entities.NominatifDepositResponse, er error) {

	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return data, err.BadRequest
	}

	// Add separator on date payload
	DtWithSeparators, _ := helper.AddSeparatorsOnDateStr(payload.TanggalHitung)

	// parse time : String to Date
	date, er := time.Parse("2006-01", DtWithSeparators[0:7]) // <- ini harus dicegah jika request client tidak menggunakan separator pada date..
	if er != nil {
		return data, er
	}

	// Substract date
	YYYYMM1 := date.AddDate(0, -1, 0)
	YYYYMM2 := date.AddDate(0, 0, -1)

	lastMonthBeginDt := YYYYMM1.Format(helper.YYYYMMDDV2) // <- Get begin date in last month
	lastMonthEndDt := YYYYMM2.Format(helper.YYYYMMDDV2)   // <- Get end date in last month
	lastMonthEndTg := YYYYMM2.Day()                       // <- Get end date in last month

	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
	data, er = tabtransRepo.GetNominatifDeposit(payload.TanggalHitung, lastMonthBeginDt, lastMonthEndDt, lastMonthEndTg, limitOffset)
	if er != nil {
		return data, nil
	}

	return data, nil
}

func (t *tabtransUsecase) GetLaporanTransaksi(payload entities.DaftarTransaksiRequest, limitOffset entities.LimitOffsetLkmUri) (data []entities.DaftarTransaksiResponse, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return data, err.BadRequest
	}

	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
	data, er = tabtransRepo.GetLaporanTransaksi(payload, limitOffset)
	if er != nil {
		return data, er
	}

	if len(data) == 0 {
		return make([]entities.DaftarTransaksiResponse, 0), nil
		// ^ []entities.DaftarTransaksiResponse (untuk membuat sebuah slice kosong agar tidak return null di JSON) | err.NoRecord
	}

	return data, nil
}

func (t *tabtransUsecase) GetListsTransaksiDeposit(payload entities.GetListsDepositTrxReq) (tx []entities.GetListsDepositTrxRes, er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	tx, er = tabtransRepo.GetListsTransaksiDeposit(payload)
	if er != nil {
		return tx, er
	}

	if len(tx) == 0 {
		return tx, err.NoRecord
	}

	return tx, nil
}

func (t *tabtransUsecase) TransaksiDeposit(payload entities.DepositRequest) (er error) {
	referensiRepo, _ := informationlkmrepo.NewInformationLKMRepo()
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	lkmInfo, er := referensiRepo.LKMInformation(payload.KodeLKM)
	if er != nil {
		return er
	}

	if lkmInfo.StatusRekening == 0 {
		return err.RekeningBelumAktif
	} else if lkmInfo.StatusRekening == 2 {
		return err.RekeningNonAktif
	} else if lkmInfo.StatusRekening == 3 {
		return err.RekeningDitutup
	} else if lkmInfo.StatusRekening == 4 {
		return err.RekeningDiBlokir
	}

	// Generating Kuitansi
	dateStr := helper.GetCurrentDate(helper.YYYYMMDDV2)
	randomNum := helper.String(8)

	data := entities.Deposit{}
	data.Tgl_trans = time.Now()
	data.NoRekening = payload.KodeLKM
	data.Pokok = payload.JumlahTransaksi
	data.Keterangan = payload.Keterangan
	data.Verifikasi = constants.Verifikasi
	data.UserID = payload.UserID
	data.ModulIDSource = constants.EmptyStr
	data.TransIDSource = constants.TransIDSource
	data.KodeTrans = payload.JenisTransaksi
	data.Tob = constants.EmptyStr
	data.PostedToGl = constants.PostedToGl
	data.KodePerkOB = constants.EmptyStr
	data.KodeKantor = constants.KodeKantor
	data.SandiTrans = constants.EmptyStr
	data.Kuitansi = dateStr + "#" + randomNum
	data.CounterSign = constants.CounterSign
	data.NoRekeningABA = constants.EmptyStr
	data.MyKodeTrans = constants.Kredit

	// Validasi jika yg di kirim adalah kode_trans 200 (penarikan dana deposit)
	if payload.JenisTransaksi == "200" {
		data.MyKodeTrans = constants.Debit
	}

	er = tabtransRepo.TransaksiDeposit(data)
	if er != nil {
		return er
	}

	er = tabtransRepo.RepostingSaldoOnRekeningLKM(payload.KodeLKM)
	if er != nil {
		return er
	}

	// TODO: Validasi jika saldo akhir rekening LIKM <= 0 set nonaktif

	return
}

func (t *tabtransUsecase) PembatalanTransaksiDeposit(payload entities.ReversalDepositRequest) (er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	trx, er := tabtransRepo.GetTransaksiDeposit(payload.Kuitansi)
	if er != nil {
		return er
	}

	reversalTrx := trx
	if reversalTrx.KodeTrans == "100" {
		reversalTrx.KodeTrans = "200"
	} else if reversalTrx.KodeTrans == "102" {
		reversalTrx.KodeTrans = "202"
	} else if reversalTrx.KodeTrans == "111" {
		reversalTrx.KodeTrans = "211"
	} else if reversalTrx.KodeTrans == "115" {
		reversalTrx.KodeTrans = "215"
	} else if reversalTrx.KodeTrans == "200" {
		reversalTrx.KodeTrans = "200"
	}
	reversalTrx.Tgl_trans = time.Now()
	reversalTrx.Keterangan = "PEMBTALAN TRANSAKSI : " + trx.Keterangan
	reversalTrx.MyKodeTrans = "200"
	reversalTrx.UserID = payload.UserID

	er = tabtransRepo.TransaksiDeposit(reversalTrx)
	if er != nil {
		return er
	}

	er = tabtransRepo.RepostingSaldoOnRekeningLKM(reversalTrx.NoRekening)
	if er != nil {
		return er
	}

	return
}
