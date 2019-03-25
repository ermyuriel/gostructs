package gostructs

import "time"

type CeleryMessage struct {
	ID      string      `json:"id"`
	Task    string      `json:"task"`
	Args    []string    `json:"args"`
	Kwargs  interface{} `json:"kwargs"`
	Retries int         `json:"retries"`
	Eta     time.Time   `json:"eta"`
	Expires time.Time   `json:"expires"`
}
