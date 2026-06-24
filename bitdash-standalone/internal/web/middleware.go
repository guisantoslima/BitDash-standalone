
package web

import (
	"log/slog"
	"net/http"
	"time"
)

// LoggingMiddleware loga método, path, status e duração de cada requisição,
// usando o logger estruturado da aplicação (Blueprint 29 - eventos de log).
func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}

			next.ServeHTTP(sw, req)

			logger.Info("http_request",
				"method", req.Method,
				"path", req.URL.Path,
				"status", sw.status,
				"duration_ms", time.Since(start).Milliseconds(),
			)
		})
	}
}

// RecoverMiddleware captura panics em handlers, registra no log e responde
// com 500, evitando que a aplicação caia (Blueprint 30 - tratamento de erro).
func RecoverMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					logger.Error("panic_recovered", "error", rec, "path", req.URL.Path)
					http.Error(w, "Erro interno. Verifique os logs.", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, req)
		})
	}
}

// statusWriter captura o status code escrito pelo handler, para fins de log.
type statusWriter struct {
	http.ResponseWriter
	status int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.status = code
	sw.ResponseWriter.WriteHeader(code)
}

