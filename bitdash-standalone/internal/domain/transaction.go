
package domain

import "time"

// TransactionType — domínio simplificado da V1 (Blueprint 10.2).
type TransactionType string

const (
	TransactionTypeEntry      TransactionType = "ENTRY"
	TransactionTypeWithdrawal TransactionType = "WITHDRAWAL"
)

// TransactionStatus — soft-delete via status, nunca hard delete (ADR de
// segurança / regra 11.2.6 do Blueprint).
type TransactionStatus string

const (
	TransactionStatusActive   TransactionStatus = "ACTIVE"
	TransactionStatusCanceled TransactionStatus = "CANCELED"
)

// Transaction representa uma movimentação financeira do usuário.
type Transaction struct {
	ID              string            `json:"id"`
	AssetID         string            `json:"asset_id"`
	TransactionType TransactionType   `json:"transaction_type"`
	TransactionDate time.Time         `json:"transaction_date"`
	Quantity        float64           `json:"quantity"`
	UnitPrice       *float64          `json:"unit_price,omitempty"`
	TotalValue      *float64          `json:"total_value,omitempty"`
	FeeAmount       *float64          `json:"fee_amount,omitempty"`
	FeeCurrency     *string           `json:"fee_currency,omitempty"`
	Notes           *string           `json:"notes,omitempty"`
	Source          *string           `json:"source,omitempty"`
	Status          TransactionStatus `json:"status"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}

// Validate verifica as invariantes mínimas (Blueprint 11.2: quantity > 0,
// tipo deve ser ENTRY/WITHDRAWAL).
func (t *Transaction) Validate() error {
	if t.TransactionType != TransactionTypeEntry && t.TransactionType != TransactionTypeWithdrawal {
		return ErrInvalidTransactionType
	}
	if t.Quantity <= 0 {
		return ErrInvalidQuantity
	}
	if t.AssetID == "" {
		return ErrInvalidTransactionInput
	}
	return nil
}

// CalculateTotalValue aplica a regra: total_value = quantity * unit_price,
// quando total_value não foi informado explicitamente (Blueprint 11.2.4).
func (t *Transaction) CalculateTotalValue() {
	if t.TotalValue == nil && t.UnitPrice != nil {
		v := t.Quantity * (*t.UnitPrice)
		t.TotalValue = &v
	}
}

