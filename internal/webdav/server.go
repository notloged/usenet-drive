package webdav

import (
	"fmt"
	"net/http"

	"github.com/javi11/usenet-drive/internal/domain"
	"golang.org/x/net/webdav"
)

func StartServer(config domain.Config, cp UsenetConnectionPool) (*http.Server, error) {
	handler := &webdav.Handler{
		FileSystem: NewNzbFilesystem(config.NzbPath, cp),
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
