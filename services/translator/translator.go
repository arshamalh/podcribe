package translator

type I interface {
	// translate transcribed podcast line by line according to the predefined settings
	Translate(string) (string, error)
}

type translator struct {
	WrapWithTranslation   bool
	ExtractWordsWithCount bool
}

func New() *translator {
	return &translator{}
}

func (t translator) Translate(text2translate string) (string, error) {
	return "", nil
}
