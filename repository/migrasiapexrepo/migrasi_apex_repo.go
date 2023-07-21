package migrasiapexrepo

import "new-apex-api/entities"

type MigrasiApexRepo interface {
	NorekLengthEqual4() (data []entities.NorekWithNID, er error)
	UpdateNasabahIDWithNorekOnNasabah() (er error)
}
