package middlewares

import (
	"net/http"
	"oauth2-authorization/utility"
)

func RequestLogger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utility.Log("Info", w, r)
	})
}
