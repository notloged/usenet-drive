package webdav

import (
	"fmt"
	"net/http"

	"github.com/javi11/usenet-drive/internal/domain"
	"golang.org/x/net/webdav"
)

func StartServer(config domain.Config) (*http.Server, error) {
	handler := &webdav.Handler{
		FileSystem: nzbFilesystem(config.NzbPath),
		LockSystem: webdav.NewMemLS(),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.ServerPort),
		Handler: http.DefaultServeMux,
	}

	return srv, nil
}
