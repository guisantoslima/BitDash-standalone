
package domain

import "errors"

// Erros de domínio tipados, conforme Blueprint seção 30 (Erros de domínio).
// Handlers devem mapear estes erros para respostas HTTP amigáveis (Sprint 5
// fará o mapeamento centralizado; por enquanto os serviços apenas retornam
// estes valores).
var (
	ErrAssetNotFound        = errors.New("ativo não encontrado")
	ErrAssetSymbolTaken     = errors.New("já existe um ativo com este símbolo")
	ErrAssetHasTransactions = errors.New("ativo possui movimentações associadas e não pode ser excluído")
	ErrInvalidAssetInput    = errors.New("dados do ativo são inválidos")

	ErrTransactionNotFound     = errors.New("movimentação não encontrada")
	ErrInvalidTransactionInput = errors.New("dados da movimentação são inválidos")
	ErrInvalidTransactionType  = errors.New("tipo de movimentação inválido (esperado ENTRY ou WITHDRAWAL)")
	ErrInvalidQuantity         = errors.New("quantidade deve ser maior que zero")
	ErrTransactionCanceled     = errors.New("movimentação já está cancelada")

	ErrWorkspaceInvalid     = errors.New("workspace inválido ou inacessível")
	ErrWorkspaceNotFound    = errors.New("workspace não encontrado no caminho informado")
	ErrWorkspaceNoWritable  = errors.New("workspace sem permissão de escrita")
	ErrWorkspaceIncompatible = errors.New("pasta existe mas não é compatível com a estrutura do BitDash")

	ErrPythonNotFound  = errors.New("interpretador Python não encontrado")
	ErrAnalyticsFailed = errors.New("falha ao executar engine analítica")

	ErrMigrationFailed = errors.New("falha ao aplicar migrations")
)

