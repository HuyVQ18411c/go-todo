package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		next.ServeHTTP(response, request)
		log.Println(request.Method, request.URL.Path, time.Since(start))
	})
}
