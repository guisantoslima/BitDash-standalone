
package workspace

import (
	"os"
	"path/filepath"
	"runtime"
)

// Estrutura padrão do workspace BitDashData, conforme README seção 7 e
// Blueprint seção 6.2.
const (
	DBFileName     = "bitdash.db"
	ConfigFileName = "bitdash.config.json"
	BackupsDir     = "backups"
	ExportsDir     = "exports"
	LogsDir        = "logs"
	TempDir        = "temp"
)

// DefaultWorkspacePath resolve a pasta padrão do BitDash de acordo com o
// sistema operacional (README seção 13, Opção A):
//   - Windows: %USERPROFILE%/BitDashData
//   - Linux/macOS: ~/BitDashData
func DefaultWorkspacePath() (string, error) {
	if runtime.GOOS == "windows" {
		profile := os.Getenv("USERPROFILE")
		if profile == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			profile = home
		}
		return filepath.Join(profile, "BitDashData"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "BitDashData"), nil
}

// DBPath retorna o caminho completo do arquivo bitdash.db dentro do workspace.
func DBPath(workspacePath string) string {
	return filepath.Join(workspacePath, DBFileName)
}

// ConfigPath retorna o caminho completo do bitdash.config.json.
func ConfigPath(workspacePath string) string {
	return filepath.Join(workspacePath, ConfigFileName)
}

// BackupsPath retorna o caminho da pasta de backups.
func BackupsPath(workspacePath string) string {
	return filepath.Join(workspacePath, BackupsDir)
}

// ExportsPath retorna o caminho da pasta de exports.
func ExportsPath(workspacePath string) string {
	return filepath.Join(workspacePath, ExportsDir)
}

// LogsPath retorna o caminho da pasta de logs.
func LogsPath(workspacePath string) string {
	return filepath.Join(workspacePath, LogsDir)
}

// TempPath retorna o caminho da pasta temp.
func TempPath(workspacePath string) string {
	return filepath.Join(workspacePath, TempDir)
}

