package internal

type Country struct {
	Base
	Name   string `json:"name"`
	Status uint   `json:"status"`
	Code   string `json:"code"`
}
