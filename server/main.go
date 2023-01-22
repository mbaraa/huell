package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	// bytes
	maxFileSize     int64 = 512 * 1024 * 1024
	defaultFileSize int64 = 1024 * 1024
)

func createFile(fileSize int64) (*os.File, error) {
	f, err := os.Create(fmt.Sprint(time.Now().UnixMicro()))
	defer f.Close()
	if err != nil {
		return nil, err
	}

	f.WriteString("I'll Dance")

	for i := int64(0); i < fileSize/6; i++ {
		f.WriteString(" Dance")
	}

	return f, nil
}

func handleCreateAndDownloadFile(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Access-Control-Allow-Origin", "*")
	fileSizeStr := req.URL.Query().Get("fs")

	fileSize, err := strconv.ParseInt(fileSizeStr, 10, 64)
	if err != nil || fileSize > maxFileSize {
		fileSize = defaultFileSize
	}

	file, err := createFile(fileSize)
	if err != nil {
		resp.WriteHeader(http.StatusInsufficientStorage)
	}

	bb, err := os.ReadFile(file.Name())
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
	}

	resp.Write(bb)
	os.Remove(file.Name())
}

func main() {
	http.HandleFunc("/meow", handleCreateAndDownloadFile)
	http.ListenAndServe(":9090", nil)
}
