package validation

type ValidationError struct {
	Key       string `json:"key"`
	ErrorMsg  string `json:"error_msg"`
	Condition string `json:"condition"`
}
