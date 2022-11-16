package err

import "errors"

var (
	NoRecord             = errors.New("no records found")
	PasswordDontMatch    = errors.New("password don't match")
	InternalServiceError = errors.New("internal service error")
	DuplicateEntry       = errors.New("duplicate entry")
	BadRequest           = errors.New("bad request")
)
