package err

import "errors"

var (
	NoRecord             = errors.New("no records found")
	PasswordDontMatch    = errors.New("password don't match")
	InternalServiceError = errors.New("internal service error")
	DuplicateEntry       = errors.New("duplicate entry")
	BadRequest           = errors.New("bad request")

	HeaderRequired   = errors.New("token or user_id is required")
	UserIDDonthMatch = errors.New("user id don't match")
	InvalidToken     = errors.New("invalid token")
)

var (
	RekeningBelumAktif = errors.New("rekening belum aktif")
	RekeningNonAktif   = errors.New("rekening nonaktif")
	RekeningDitutup    = errors.New("rekening ditutup")
	RekeningDiBlokir   = errors.New("rekening diblokir")
)
