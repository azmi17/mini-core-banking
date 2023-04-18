package usecase

import (
	"new-apex-api/entities/err"
	"new-apex-api/entities/web"
	"new-apex-api/helper"
	"new-apex-api/repository/informationlkmrepo"
	"new-apex-api/repository/tabtransrepo"
	"time"
)

type TabtransUsecase interface {
	GetListsTabtransTrx(tglTrans web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (
		[]web.GetListTabtransTrx,
		web.GetCountWithSumTabtransTrx,
		error,
	)
	GetListsTabtransTrxBySTAN(stan string) ([]web.GetListTabtransTrx, error)
	DeleteTabtransTrx(tabtransID int) error
	ChangeDateOnTabtransTrx(tabtransID int, tglTrans string) (web.GetListTabtransTrx, error)
	GetRekeningKoranLKMDetail(web.RekeningKoranRequest) (data web.RekeningKoranResponse, er error)
	GetNominatifDeposit(payload web.NominatifDepositRequest, limitOffset web.LimitOffsetLkmUri) (data []web.NominatifDepositResponse, er error)
	GetLaporanTransaksi(payload web.DaftarTransaksiRequest, limitOffset web.LimitOffsetLkmUri) (data []web.DaftarTransaksiResponse, er error)
	GetListsTransaksiDeposit(payload web.GetListsDepositTrxReq) ([]web.GetListsDepositTrxRes, error)
}

type tabtransUsecase struct{}

func NewTabtransUsecase() TabtransUsecase {
	return &tabtransUsecase{}
}

func (t *tabtransUsecase) GetListsTabtransTrx(tglTrans web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (tabtransTxList []web.GetListTabtransTrx, total web.GetCountWithSumTabtransTrx, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return tabtransTxList, total, err.BadRequest
	}

	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
	tabtransTxList, total, er = tabtransRepo.GetListsTabtransTrx(tglTrans, limitOffset)
	if er != nil {
		return tabtransTxList, total, er
	}

	if len(tabtransTxList) == 0 {
		return make([]web.GetListTabtransTrx, 0), total, nil
		// ^ []web.GetDetailLKMInfo (untuk membuat sebuah slice kosong agar tidak return null di JSON) | err.NoRecord
	}

	return tabtransTxList, total, nil
}

func (t *tabtransUsecase) GetListsTabtransTrxBySTAN(stan string) (tx []web.GetListTabtransTrx, er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	tx, er = tabtransRepo.GetListsTabtransTrxBySTAN(stan)
	if er != nil {
		return tx, er
	}

	if len(tx) == 0 {
		return tx, err.NoRecord
	}

	return tx, nil
}

func (t *tabtransUsecase) DeleteTabtransTrx(tabtransID int) (er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	if er = tabtransRepo.DeleteTabtransTrx(tabtransID); er != nil {
		return er
	}

	return nil
}

func (t *tabtransUsecase) ChangeDateOnTabtransTrx(tabtransID int, tglTrans string) (data web.GetListTabtransTrx, er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	if data, er = tabtransRepo.ChangeDateOnTabtransTrx(tabtransID, tglTrans); er != nil {
		return data, er
	}

	return
}

func (t *tabtransUsecase) GetRekeningKoranLKMDetail(payload web.RekeningKoranRequest) (resp web.RekeningKoranResponse, er error) {
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

	periodeAwal, _ := helper.ParseTimeStrToDate(payloadPeriodeAkhir)
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
	saldoAwal, er := tabtransRepo.GetSaldo(payload.KodeLKM, periodeAwal) // tanggal periode awal harus di substract
	if er != nil {
		return resp, er
	}
	resp.SaldoAwal = saldoAwal

	data, er := tabtransRepo.GetRekeningKoranLKMDetail(payload.KodeLKM, payload.PeriodeAwal, payload.PeriodeAkhir)
	if er != nil {
		return resp, nil
	}

	saldoTemp := saldoAwal
	resp.Detail = make([]web.RekeningKoranDetail, 0) // default 0 for detail field
	for _, v := range data {
		var detail web.RekeningKoranDetail
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

func (t *tabtransUsecase) GetNominatifDeposit(payload web.NominatifDepositRequest, limitOffset web.LimitOffsetLkmUri) (data []web.NominatifDepositResponse, er error) {

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

func (t *tabtransUsecase) GetLaporanTransaksi(payload web.DaftarTransaksiRequest, limitOffset web.LimitOffsetLkmUri) (data []web.DaftarTransaksiResponse, er error) {
	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 {
		return data, err.BadRequest
	}

	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
	data, er = tabtransRepo.GetLaporanTransaksi(payload, limitOffset)
	if er != nil {
		return data, er
	}

	if len(data) == 0 {
		return make([]web.DaftarTransaksiResponse, 0), nil
		// ^ []web.DaftarTransaksiResponse (untuk membuat sebuah slice kosong agar tidak return null di JSON) | err.NoRecord
	}

	return data, nil
}

func (t *tabtransUsecase) GetListsTransaksiDeposit(payload web.GetListsDepositTrxReq) (tx []web.GetListsDepositTrxRes, er error) {
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
