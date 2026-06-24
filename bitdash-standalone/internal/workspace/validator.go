
package workspace

import (
	"os"

	"bitdash-standalone/internal/domain"
)

// ValidationResult descreve o estado de uma pasta candidata a workspace,
// usado para decidir o fluxo de onboarding (Blueprint seção 8: Workspace
// novo / Workspace já BitDash / Pasta não compatível).
type ValidationResult struct {
	Exists        bool // a pasta existe no filesystem
	IsWritable    bool // há permissão de escrita
	HasDB         bool // bitdash.db já existe
	HasConfig     bool // bitdash.config.json já existe
	IsEmptyFolder bool // pasta existe mas está vazia (candidata a inicialização)
}

// IsValidWorkspace indica se a pasta já é um workspace BitDash válido e
// pronto para uso (banco + config presentes).
func (r ValidationResult) IsValidWorkspace() bool {
	return r.Exists && r.HasDB && r.HasConfig
}

// IsIncompatible indica se a pasta existe, tem conteúdo, mas não é uma
// estrutura BitDash reconhecível (Blueprint 8.C: "Pasta não compatível").
func (r ValidationResult) IsIncompatible() bool {
	return r.Exists && !r.IsEmptyFolder && !r.HasDB && !r.HasConfig
}

// ValidateWorkspace inspeciona o caminho informado e retorna um
// ValidationResult, sem efeitos colaterais (não cria nada).
func ValidateWorkspace(path string) (ValidationResult, error) {
	result := ValidationResult{}

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		result.Exists = false
		return result, nil
	}
	if err != nil {
		return result, err
	}
	if !info.IsDir() {
		return result, domain.ErrWorkspaceInvalid
	}

	result.Exists = true

	// Checagem de permissão de escrita: tenta criar e remover um arquivo
	// temporário dentro da pasta.
	testFile := DBPath(path) + ".writetest"
	if f, err := os.Create(testFile); err == nil {
		f.Close()
		os.Remove(testFile)
		result.IsWritable = true
	} else {
		result.IsWritable = false
	}

	if _, err := os.Stat(DBPath(path)); err == nil {
		result.HasDB = true
	}
	if _, err := os.Stat(ConfigPath(path)); err == nil {
		result.HasConfig = true
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return result, err
	}
	result.IsEmptyFolder = len(entries) == 0

	return result, nil
}

