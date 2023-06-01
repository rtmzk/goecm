package v1

type Rule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Func        string `json:"func"`
	ScriptOut
}

type CheckRules struct {
	Rules []Rule `json:"rules"`
}

type ScriptOut struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
