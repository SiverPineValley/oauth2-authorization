package middlewares

import (
	"net/http"
	"oauth2-authorization/models"
	"oauth2-authorization/utility"
)

func RequestLogger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utility.Log(models.LogLevelInfo, w, r)
	})
}
