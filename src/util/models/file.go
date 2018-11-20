package models

type File struct {
	UID      int64  `json:"uid"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	MimeType string `json:"type"`
	Path     string `json:"path"`
	Hash     string `json:"hash"`
}
