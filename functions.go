package upload

import (
	"io"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/mark-francis/evidence/app"
)

func init() {
	// Register an HTTP function with the Functions Framework
	functions.HTTP("EvidenceUpload", upload)
	functions.HTTP("EvidenceDownload", download)
}

// HTTP handler
func upload(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Trim(r.URL.Path, "/")

	msg, err := app.Upload(fileName)
	if err != nil {
		panic(err)
	}

	// Send an HTTP response
	if _, err := io.WriteString(w, msg); err != nil {
		panic(err)
	}
}

// HTTP handler
func download(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Trim(r.URL.Path, "/")

	msg, err := app.Download(fileName)
	if err != nil {
		panic(err)
	}

	// Send an HTTP response
	if _, err := io.WriteString(w, msg); err != nil {
		panic(err)
	}
}
