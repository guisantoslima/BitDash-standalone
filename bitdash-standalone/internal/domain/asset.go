
package domain

import "time"

// Asset representa um criptoativo cadastrado no sistema (Blueprint 8.1 / 10.1).
type Asset struct {
	ID        string    `json:"id"`
	Symbol    string    `json:"symbol"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate verifica as invariantes mínimas de um Asset.
// Regras de negócio mais ricas (ex: unicidade de symbol) ficam no service,
// pois dependem de consulta ao repositório.
func (a *Asset) Validate() error {
	if a.Symbol == "" {
		return ErrInvalidAssetInput
	}
	if a.Name == "" {
		return ErrInvalidAssetInput
	}
	return nil
}

