package manager

// TODO: Move these setting to user settings and breaks manager.Start function to some more functions link StartFullFlow...

type FlowType string

const (
	// Do all the steps this bot can do, including:
	// finding the podcast link
	FullFlow                FlowType = "full-flow"
	JustDownload            FlowType = "just-down"
	NoTranslation           FlowType = "no-translations"
	TranscribeDownloadedMP3 FlowType = "transcribe-downloaded-mp3"
	TranscribeDownloadedWAV FlowType = "transcribe-downloaded-wav"
	TranslateDownloadedMP3  FlowType = "translate-downloaded-mp3"
	TranslateDownloadedWAV  FlowType = "translate-downloaded-wav"
)

type managerSettings struct {
	flowType FlowType
}
