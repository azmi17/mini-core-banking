package entities

type Approval struct {
	Id           int
	UserID       int
	OtorisatorID int
	Token        string
	Status       int
	Description  string
	Time         string
	Expired      string
}

type ApprovalRequest struct {
	UserID       int    `form:"user_id" binding:"required"`
	OtorisatorID int    `form:"otorisator_id" binding:"required"`
	Description  string `form:"description" binding:"required"`
}

type ApprovalResponse struct {
	Id           int    `json:"id"`
	UserID       int    `json:"user_d"`
	UserName     string `json:"nama_user"`
	OtorisatorID int    `json:"otorisator_id"`
	// OtorisatorName string    `json:"nama_otorisator"`
	Token       string `json:"token"`
	Status      int    `json:"status"`
	Description string `json:"deskripsi"`
	Time        string `json:"waktu"`
	Expired     string `json:"expired"`
}
