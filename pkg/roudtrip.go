package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Базовый Middleware для логгирования запросов
type LoggingRoundTripper struct {
	Logger io.Writer
	Next   http.RoundTripper
}

func (l LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.Logger, "[%s] %s %sn", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.Next.RoundTrip(r)
}

// Middleware ЛОГИРОВАНИЯ ДАННЫХ В ЗАПРОСЕ
func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("[%s] %s\n", r.Method, r.URL)
		next(w, r)
	}
}
