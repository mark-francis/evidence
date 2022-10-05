package upload

import (
	"io"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/mark.francis/evidence/app"
)

func init() {
	// Register an HTTP function with the Functions Framework
	functions.HTTP("EvidenceUpload", myHTTPFunction)
}

// Function myHTTPFunction is an HTTP handler
func myHTTPFunction(w http.ResponseWriter, r *http.Request) {
	// Your code here
	msg, err := app.Upload("test")
	if err != nil {
		panic(err)
	}

	// Send an HTTP response
	if _, err := io.WriteString(w, msg); err != nil {
		panic(err)
	}
}
