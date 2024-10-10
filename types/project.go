package types

type MstrProject struct {
	MstrObject
}

type MstrProjectSetting struct {
	Value interface{} `json:"value"`
}

type MstrProjectSettings map[string]MstrProjectSetting
