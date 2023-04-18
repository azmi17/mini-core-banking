package delivery

import (
	"fmt"
	"new-apex-api/entities"

	"github.com/kpango/glg"
)

func PrintoutObserver() {
	for po := range entities.PrintOutChan {
		if po.Type == entities.PRINTOUT_TYPE_ERR {
			_ = glg.Error(po.Message...)
		} else if po.Type == entities.PRINTOUT_TYPE_LOG {
			_ = glg.Log(po.Message...)
		}
	}
}

func PrintRepoResult() {

	for po := range entities.PrintRepoChan {
		if po.Status == "0000" {
			fmt.Println("sukses reposting:", po.KodeLKM)
		} else {
			fmt.Println("gagal reposting:", po.KodeLKM)
		}
	}
}
