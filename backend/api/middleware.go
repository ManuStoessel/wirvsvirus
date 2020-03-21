package api

import (
	"net/http"

	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
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

func (o *oauth2provider) verificationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		toValidate := map[string]string{}
		toValidate["aud"] = "api://default"
		toValidate["cid"] = o.ClientID

		jwtVerifierSetup := jwtverifier.JwtVerifier{
			Issuer:           o.Issuer,
			ClaimsToValidate: toValidate,
		}

		verifier := jwtVerifierSetup.New()

		_, err := verifier.VerifyAccessToken(r.Header.Get("Authorization"))

		if err != nil {
			log.WithFields(log.Fields{
				"method": r.Method,
				"path":   r.URL.EscapedPath(),
				"host":   r.URL.Hostname(),
				"token":  r.Header.Get("Authorization"),
			}).Warn("Failed to verify token")
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
