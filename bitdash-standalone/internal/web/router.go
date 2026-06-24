
package web

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewRouter monta o router base da aplicação com os middlewares globais e a
// rota placeholder "/". Handlers reais de cada módulo (assets, transactions,
// dashboard, settings, workspace) serão registrados em sprints futuros, mas
// a função já recebe o Renderer para que os próximos sprints só precisem
// adicionar `r.Get(...)` / `r.Post(...)` aqui.
func NewRouter(logger *slog.Logger, renderer *Renderer) http.Handler {
	r := chi.NewRouter()

	r.Use(LoggingMiddleware(logger))
	r.Use(RecoverMiddleware(logger))

	// Placeholder: a partir do Sprint 3 isto redireciona para /dashboard.
	// Por ora, apenas confirma que o servidor está de pé.
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("BitDash-standalone — Foundation OK. Workspace e banco inicializados."))
	})

	r.Get("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return r
}

