package api

type DumpRequest struct {
	Path    string `json:"path"`
	Timeout int64  `json:"timeout"`
}
