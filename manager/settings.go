package manager

// TODO: Move these setting to user settings and breaks manager.Start function to some more functions link StartFullFlow...

type FlowType string

const (
	// Do all the steps this bot can do, including:
	// finding the podcast link
	FullFlow      FlowType = "full-flow"
	JustDownload  FlowType = "just-down"
	NoTranslation FlowType = "no-translations"
)

type managerSettings struct {
	flowType FlowType
}
