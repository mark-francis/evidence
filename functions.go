package upload

import (
	"context"
	"io"
	"net/http"

	"github.com/A-SCEND/muxxer"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gorilla/mux"
	"github.com/mark-francis/evidence/app"
)

var srv *muxxer.Server

func init() {
	// Register an HTTP function with the Functions Framework

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv = muxxer.New(ctx, "evidence")

	srv.AddRoutes(
		muxxer.Get("/upload/{fileName}", upload),
		muxxer.Get("/download/{fileName}", download),
	)

	functions.HTTP("EvidenceService", serve)
}

func serve(w http.ResponseWriter, r *http.Request) {
	srv.ServeRequest(w, r)
}

// HTTP handler
func upload(w io.Writer, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	return app.Upload(vars["fileName"])
}

// HTTP handler
func download(w io.Writer, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	return app.Download(vars["fileName"])
}
