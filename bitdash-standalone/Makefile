
.PHONY: dev test migrate seed package

# Sobe a aplicação em modo desenvolvimento
dev:
	go run ./cmd/bitdash

# Executa a suite de testes Go
test:
	go test ./... -v

# Aplica migrations manualmente (útil para debug; o app já aplica no startup)
migrate:
	go run ./cmd/bitdash --migrate-only

# Popula o workspace com dados de exemplo (seed) - implementado em sprints futuros
seed:
	@echo "seed: a implementar em sprint futuro"

# Gera o binário de distribuição
package:
	go build -o bin/bitdash ./cmd/bitdash

