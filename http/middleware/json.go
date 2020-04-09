package middleware

import (
	"errors"
	"net/http"
	"strings"

	"Sharykhin/buffstream-questionnaire/http/response"
)

//JsonContentType validates that a correct header is represented in a request
func JsonContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc := r.Header.Get("Content-Type")
		if (r.Method == "POST" || r.Method == "PUT") &&
			!strings.Contains(cc, "application/json") {
			response.BadRequest(w, errors.New("Content-Type must be application/json"))
			return
		}

		h.ServeHTTP(w, r)
		return
	})
}
