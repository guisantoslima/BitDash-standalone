
package workspace

import (
	"encoding/json"
	"os"
	"time"
)

// WorkspaceConfig é o conteúdo persistido em bitdash.config.json
// (Blueprint 28.2 / README seção 7).
type WorkspaceConfig struct {
	WorkspaceVersion string `json:"workspace_version"`
	DefaultCurrency  string `json:"default_currency"`
	DefaultLocale    string `json:"default_locale"`
	DefaultPort      string `json:"default_port"`
	CreatedAt        string `json:"created_at"`
	LastBackupAt      string `json:"last_backup_at,omitempty"`
}

const CurrentWorkspaceVersion = "1.0.0"

// CreateDefaultStructure cria a árvore de diretórios e o arquivo de config
// inicial de um novo workspace (Blueprint 5.1 / README seção 7).
// O arquivo bitdash.db é criado separadamente pela camada database, pois
// depende da abertura da conexão SQLite.
func CreateDefaultStructure(path string) error {
	dirs := []string{
		path,
		BackupsPath(path),
		ExportsPath(path),
		LogsPath(path),
		TempPath(path),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	// Só cria o config.json se ainda não existir, para não sobrescrever
	// configurações de um workspace já inicializado.
	if _, err := os.Stat(ConfigPath(path)); os.IsNotExist(err) {
		cfg := WorkspaceConfig{
			WorkspaceVersion: CurrentWorkspaceVersion,
			DefaultCurrency:  "BRL",
			DefaultLocale:    "pt-BR",
			DefaultPort:      "8080",
			CreatedAt:        time.Now().UTC().Format(time.RFC3339),
		}
		return writeWorkspaceConfig(path, cfg)
	}

	return nil
}

func writeWorkspaceConfig(path string, cfg WorkspaceConfig) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigPath(path), data, 0o644)
}

// LoadWorkspaceConfig lê o bitdash.config.json existente no workspace.
func LoadWorkspaceConfig(path string) (WorkspaceConfig, error) {
	var cfg WorkspaceConfig
	data, err := os.ReadFile(ConfigPath(path))
	if err != nil {
		return cfg, err
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

