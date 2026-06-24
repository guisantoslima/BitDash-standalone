
package web

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// Renderer carrega e renderiza os templates html/template a partir da pasta
// web/templates, conforme Blueprint 16 (estrutura do repositório) e ADR-006
// (server-side templates + HTMX, sem SPA).
type Renderer struct {
	templatesDir string
}

// NewRenderer cria um Renderer apontando para o diretório base de templates.
func NewRenderer(templatesDir string) *Renderer {
	return &Renderer{templatesDir: templatesDir}
}

// Render renderiza um template nomeado (relativo a templatesDir) combinado
// com o layout base, escrevendo a resposta HTML.
//
// Nesta fase (Sprint 0) o helper é minimalista: nos próximos sprints cada
// handler usará Render para suas páginas (dashboard, assets, transactions,
// settings).
func (r *Renderer) Render(w http.ResponseWriter, page string, data any) error {
	layout := filepath.Join(r.templatesDir, "layouts", "base.html")
	pagePath := filepath.Join(r.templatesDir, page)

	tmpl, err := template.ParseFiles(layout, pagePath)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, "base", data)
}

