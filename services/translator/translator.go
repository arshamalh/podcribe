package translator

import "podcribe/entities"

type I interface {
	// translate transcribed podcast line by line according to the predefined settings
	Translate(podcast *entities.Podcast) error
}

type translator struct {
	WrapWithTranslation   bool
	ExtractWordsWithCount bool
}

func New() *translator {
	return &translator{}
}

func (t translator) Translate(podcast *entities.Podcast) error {
	return nil
}
