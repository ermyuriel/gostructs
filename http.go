package gostructs

type HTTPRequest struct {
	Method    string              `json:"method"`
	Header    map[string][]string `json:"header"`
	Body      interface{}         `json:"body"`
	Source    string              `json:"source"`
	Target    string              `json:"target"`
	Timestamp int64               `json:"timestamp"`
}
