package api

type DumpRequest struct {
	Key     string `json:"path"`
	Timeout int64  `json:"timeout"`
}
