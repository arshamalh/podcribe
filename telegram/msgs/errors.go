package msgs

import "fmt"

const (
	CantDownloadFile     = "can't download file from telegram, isn't it less than 20 MB?"
	FileTypeNotSupported = "file type is not supported, message the admin: "
	CantTranscribe       = "something unexpected happened when transcribing"
	UnableToConvert      = "unable to convert telegram voice to AI understandable format"
)

type NotEnoughBalance float64

func (neb NotEnoughBalance) Error() string {
	return fmt.Sprintf("not enough balance, you should have at least %.3f dollars", neb)
}
