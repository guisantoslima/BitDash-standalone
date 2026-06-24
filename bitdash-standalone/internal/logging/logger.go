
package logging

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

// New cria um *slog.Logger estruturado que escreve simultaneamente em
// BitDashData/logs/bitdash.log e em stderr, conforme README seção 18 e
// Blueprint seção 29.
//
// Eventos recomendados para logging (a serem emitidos pelos módulos que
// usarem este logger em sprints futuros):
//   inicialização da aplicação, criação/validação do workspace, aplicação
//   de migrations, erros de acesso ao banco, erros na execução da engine
//   Python, exportações e backups, cancelamento de movimentações.
func New(workspaceLogsDir string, debug bool) (*slog.Logger, error) {
	if err := os.MkdirAll(workspaceLogsDir, 0o755); err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile(
		filepath.Join(workspaceLogsDir, "bitdash.log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o644,
	)
	if err != nil {
		return nil, err
	}

	writer := io.MultiWriter(os.Stderr, logFile)

	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}

	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: level,
	})

	return slog.New(handler), nil
}

