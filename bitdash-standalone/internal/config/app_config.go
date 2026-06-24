
package config

import (
	"os"
	"strconv"
)

// AppConfig representa a configuração do host Go (Blueprint 28.1).
// É carregada a partir de variáveis de ambiente (ou .env via os.Getenv,
// sem dependência externa nesta fase).
type AppConfig struct {
	Port         string
	Debug        bool
	PythonPath   string
	Locale       string
	WorkspacePath string // se vazio, usa a pasta padrão do OS
}

// LoadAppConfig lê a configuração do host a partir do ambiente, aplicando
// defaults sensatos quando uma variável não está definida.
func LoadAppConfig() AppConfig {
	return AppConfig{
		Port:          getEnvOrDefault("BITDASH_PORT", "8080"),
		Debug:         getBoolEnvOrDefault("BITDASH_DEBUG", false),
		PythonPath:    getEnvOrDefault("BITDASH_PYTHON_PATH", "python3"),
		Locale:        getEnvOrDefault("BITDASH_LOCALE", "pt-BR"),
		WorkspacePath: os.Getenv("BITDASH_WORKSPACE_PATH"),
	}
}

func getEnvOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getBoolEnvOrDefault(key string, def bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	parsed, err := strconv.ParseBool(v)
	if err != nil {
		return def
	}
	return parsed
}

