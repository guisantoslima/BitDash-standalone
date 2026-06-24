
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// RuntimeConfig é a configuração global do host, persistida FORA do
// workspace (ex: ~/.bitdash/runtime.json), guardando o último workspace
// utilizado — conforme Blueprint 7.1 Passo 1 ("localizar último workspace
// usado").
type RuntimeConfig struct {
	LastWorkspacePath string `json:"last_workspace_path"`
}

// runtimeConfigPath retorna o caminho do arquivo de config global do host,
// independente do workspace (ex: $HOME/.bitdash/runtime.json).
func runtimeConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".bitdash", "runtime.json"), nil
}

// LoadRuntimeConfig carrega a config global do host. Se não existir,
// retorna uma RuntimeConfig vazia (sem erro) — é um cenário esperado na
// primeira execução do app.
func LoadRuntimeConfig() (RuntimeConfig, error) {
	path, err := runtimeConfigPath()
	if err != nil {
		return RuntimeConfig{}, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return RuntimeConfig{}, nil
	}
	if err != nil {
		return RuntimeConfig{}, err
	}

	var cfg RuntimeConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return RuntimeConfig{}, err
	}
	return cfg, nil
}

// SaveRuntimeConfig persiste a config global do host, criando o diretório
// pai se necessário.
func SaveRuntimeConfig(cfg RuntimeConfig) error {
	path, err := runtimeConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

