package api

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Access Logs
		log.WithFields(log.Fields{
			"method": r.Method,
			"path":   r.URL.EscapedPath(),
			"host":   r.URL.Hostname(),
		}).Info("Incoming request")

		next.ServeHTTP(w, r)
	})
}

type oauth2provider struct {
	Issuer   string
	ClientID string
}
