package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/acme/autocert"
)

func RunServe(_ *cobra.Command, _ []string) error {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("cert-cache"),
		HostPolicy: autocert.HostWhitelist("jrocha.eu"),
	}

	server := NewServer(certManager)

	if server.Err != nil {
		err := fmt.Errorf("error: a problem occurred setting the file system: %w", server.Err)
		return err
	}
	log.Println("Server started at http://localhost" + server.Config.Addr)

	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	server.Config.ListenAndServeTLS("", "")

	return nil
}
