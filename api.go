package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
}

func CreateApiServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddress: listenAddr,
	}
}

func (a *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/handleFile", wrapHandleFunc(a.HandleFile))

	log.Println("API Serer running on port: ", a.listenAddress)

	http.ListenAndServe(a.listenAddress, router)
}

func (a *APIServer) HandleFile(w http.ResponseWriter, r *http.Request) error {
	const maxFileSize = 100 * 1024 * 1024 // 100 MB size
	file, header, err := r.FormFile("file")
	if err != nil {
		JsonStringifyResponse(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	defer file.Close()
	if header.Size > maxFileSize {
		JsonStringifyResponse(w, http.StatusExpectationFailed, ApiError{"File size exceeds the limit"})
	}
	hashMD5 := md5.New()
	hashSHA1 := sha1.New()
	hashSHA256 := sha256.New()

	_, err = io.Copy(io.MultiWriter(hashMD5, hashSHA1, hashSHA256), file)

	if err != nil {
		JsonStringifyResponse(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
	}

	fileName := header.Filename
	checksumMD5 := hex.EncodeToString(hashMD5.Sum(nil))
	checksumSHA1 := hex.EncodeToString(hashSHA1.Sum(nil))
	checksumSHA256 := hex.EncodeToString(hashSHA256.Sum(nil))

	return JsonStringifyResponse(w, http.StatusAccepted, UploadSuccessResp{ID: 132, FileName: fileName, Md5: checksumMD5, Sha1: checksumSHA1, Sha256: checksumSHA256})
}

func wrapHandleFunc(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			JsonStringifyResponse(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func JsonStringifyResponse(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "applicaiton/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
