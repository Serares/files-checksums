package main

import "net/http"

type UploadSuccessResp struct {
	ID       int    `json:"id"`
	FileName string `json:"fileName"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
}

type File struct {
	ID       int    `json:"id"`
	FileType string `json:"fileType"`
	Checksum string `json:"checksum"`
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}
