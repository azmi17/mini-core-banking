package usecase

import "new-apex-api/repository/migrasiapexrepo"

type MigrasiApexUsecase interface {
	ReplaceNasabahIDOnNasabahWithNorek() (er error)
}

type migrasiApexUsecase struct{}

func NewMigrasiApexUsecase() MigrasiApexUsecase {
	return &migrasiApexUsecase{}
}

func (lkm *migrasiApexUsecase) ReplaceNasabahIDOnNasabahWithNorek() (er error) {
	migrasiApexrepo, _ := migrasiapexrepo.NewMigrasiApexRepo()
	// data, er := migrasiApexrepo.NorekLengthEqual4()
	// if er != nil {
	// 	return er
	// }
	er = migrasiApexrepo.UpdateNasabahIDWithNorekOnNasabah()
	if er != nil {
		return er
	}

	return nil
}
