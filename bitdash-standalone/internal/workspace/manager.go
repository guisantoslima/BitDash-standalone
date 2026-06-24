
package workspace

import (
	"bitdash-standalone/internal/domain"
)

// Manager centraliza o ciclo de vida do workspace ativo da aplicação,
// conforme responsabilidades definidas no Blueprint 5.1:
//   InitializeWorkspace(path), ValidateWorkspace(path), LoadWorkspace(),
//   SwitchWorkspace(path), CreateDefaultStructure(path)
type Manager struct {
	activePath string
	config     WorkspaceConfig
}

// NewManager cria um Manager ainda sem workspace carregado.
func NewManager() *Manager {
	return &Manager{}
}

// ActivePath retorna o caminho do workspace atualmente carregado.
func (m *Manager) ActivePath() string {
	return m.activePath
}

// Config retorna a configuração do workspace atualmente carregado.
func (m *Manager) Config() WorkspaceConfig {
	return m.config
}

// InitializeWorkspace cria a estrutura do workspace em `path` (se ainda não
// existir) e o carrega como workspace ativo. Trata os 3 cenários do
// Blueprint seção 8:
//   A. pasta vazia/inexistente → cria estrutura nova
//   B. pasta já é workspace BitDash → apenas valida e conecta
//   C. pasta existe com conteúdo incompatível → retorna erro explícito
func (m *Manager) InitializeWorkspace(path string) error {
	validation, err := ValidateWorkspace(path)
	if err != nil {
		return err
	}

	if validation.IsIncompatible() {
		return domain.ErrWorkspaceIncompatible
	}

	if validation.Exists && !validation.IsWritable {
		return domain.ErrWorkspaceNoWritable
	}

	// Cenário A (novo) ou B (já válido): CreateDefaultStructure é
	// idempotente — não sobrescreve config existente.
	if err := CreateDefaultStructure(path); err != nil {
		return err
	}

	return m.LoadWorkspace(path)
}

// ValidateWorkspace exposto no Manager para conveniência de chamadas
// externas que já têm uma instância de Manager em mãos.
func (m *Manager) ValidateWorkspace(path string) (ValidationResult, error) {
	return ValidateWorkspace(path)
}

// LoadWorkspace carrega um workspace já inicializado como ativo, lendo seu
// bitdash.config.json. Não cria nada — assume que a estrutura já existe.
func (m *Manager) LoadWorkspace(path string) error {
	validation, err := ValidateWorkspace(path)
	if err != nil {
		return err
	}
	if !validation.Exists {
		return domain.ErrWorkspaceNotFound
	}

	cfg, err := LoadWorkspaceConfig(path)
	if err != nil {
		return err
	}

	m.activePath = path
	m.config = cfg
	return nil
}

// SwitchWorkspace troca o workspace ativo para um novo caminho, validando-o
// e inicializando-o se necessário antes de trocar (Blueprint 5.1).
func (m *Manager) SwitchWorkspace(path string) error {
	return m.InitializeWorkspace(path)
}

