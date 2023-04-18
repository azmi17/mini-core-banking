package entities

const (
	PRINTOUT_TYPE_LOG = iota
	PRINTOUT_TYPE_ERR

	PRINT_SUCCESS_STATUS_REPO_CHAN = "00"
	PRINT_FAILED_STATUS_REPO_CHAN  = "01"
	PRINT_MSG_REPO_CHAN            = "SUCCESS"
)

var (
	PrintOutChan  = make(chan PrintOut)
	PrintRepoChan = make(chan PrintRepo)
)

type PrintRepo struct {
	KodeLKM string
	Status  string
	Message string
}

type PrintOut struct {
	Type    int
	Message []interface{}
}

func PrintError(message ...interface{}) {
	po := PrintOut{
		Type:    PRINTOUT_TYPE_ERR,
		Message: message,
	}

	PrintOutChan <- po
}

func PrintLog(message ...interface{}) {
	po := PrintOut{
		Type:    PRINTOUT_TYPE_LOG,
		Message: message,
	}
	PrintOutChan <- po
}
