package entities

type TransHistoryResponse struct {
	TransID        int     `json:"trans_id"`
	TglTrans       string  `json:"tgl_trans"`
	KodeLKM        string  `json:"kode_lkm"`
	NamaLembaga    string  `json:"nama_lembaga"`
	SubscriberId   string  `json:"subscriber_id"`
	Dc             string  `json:"dc"`
	Amount         float64 `json:"amount"`
	ResponseCode   string  `json:"response_code"`
	Stan           string  `json:"stan"`
	Ref            string  `json:"ref"`
	RekeningID     string  `json:"rekening_id"`
	ProcessingCode string  `json:"processing_code"`
	BillerCode     string  `json:"biller_code"`
	ProductCode    string  `json:"product_code"`
}
