package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Базовый Middleware для логгирования запросов
type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %sn", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

// Middleware ЛОГИРОВАНИЯ ДАННЫХ В ЗАПРОСЕ
func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("[%s] %s\n", r.Method, r.URL)
		next(w, r)
	}
}
