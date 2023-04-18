package usecase

import (
	"new-apex-api/entities"
	"new-apex-api/repository/tabtransrepo"
	"new-apex-api/repository/tabunganrepo"
	"time"

	"github.com/kpango/glg"
	"github.com/schollz/progressbar/v3"
)

type RepostingUsecase interface {
	// RepostingSaldo(kodeLKM string) (er error)
	RepostingSaldoByBulk(listOfKodeLKM []string) (er error)
	RepostingSaldoByScheduler() (er error)
}

type repostingUsecase struct{}

func NewRepostingUsecase() RepostingUsecase {
	return &repostingUsecase{}
}

func (r *repostingUsecase) RepostingSaldoByBulk(listOfKodeLKM []string) (er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()

	er = tabtransRepo.RepostingSaldoOnRekeningLKM(listOfKodeLKM...)
	if er != nil {
		return er
	}

	return nil
}

func (r *repostingUsecase) RepostingSaldoByScheduler() (er error) {
	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
	tabungRepo, _ := tabunganrepo.NewTabunganRepo()

	list, er := tabungRepo.GetRekeningLKMByStatusActive()
	if er != nil {
		return er
	}
	// bar := progressbar.Default(int64(len(list)))
	bar := progressbar.NewOptions(len(list),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(20),
		progressbar.OptionSetDescription("[reset]Proccesing of balance reposting..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	var PrintRepoResultChan = make(chan entities.PrintRepo)
	defer close(PrintRepoResultChan)

	var numOfSuccess = 0
	var numOfFailed = 0
	go func() {
		for po := range PrintRepoResultChan {
			if po.Status == entities.PRINT_SUCCESS_STATUS_REPO_CHAN {
				numOfSuccess++
			} else {
				numOfFailed++
			}
			bar.Add(1)
			time.Sleep(1 * time.Nanosecond) // debug mode..
		}
	}()

	er = tabtransRepo.RepostingSaldoOnRekeningLKMByScheduler(PrintRepoResultChan, list...)
	if er != nil {
		return er
	}

	_ = glg.Log("REPOSTING OF SUCCESS:", numOfSuccess)
	_ = glg.Log("REPOSTING OF FAILED:", numOfFailed)

	return nil
}

// func (r *repostingUsecase) RepostingSaldoByScheduler() (er error) {
// 	tabtransRepo, _ := tabtransrepo.NewTabtransRepo()
// 	tabungRepo, _ := tabunganrepo.NewTabunganRepo()

// 	list, er := tabungRepo.GetRekeningLKMByStatusActive()
// 	if er != nil {
// 		return er
// 	}

// 	er = tabtransRepo.RepostingSaldoOnRekeningLKM(list...)
// 	if er != nil {
// 		return er
// 	}

// 	return nil
// }
