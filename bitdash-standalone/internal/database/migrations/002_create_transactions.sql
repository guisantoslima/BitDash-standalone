
CREATE TABLE IF NOT EXISTS transactions (
  id TEXT PRIMARY KEY,
  asset_id TEXT NOT NULL,
  transaction_type TEXT NOT NULL CHECK (transaction_type IN ('ENTRY', 'WITHDRAWAL')),
  transaction_date TEXT NOT NULL,
  quantity REAL NOT NULL,
  unit_price REAL,
  total_value REAL,
  fee_amount REAL,
  fee_currency TEXT,
  notes TEXT,
  source TEXT,
  status TEXT NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE', 'CANCELED')),
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL,
  FOREIGN KEY (asset_id) REFERENCES assets(id)
);

CREATE INDEX IF NOT EXISTS idx_transactions_asset_id ON transactions(asset_id);
CREATE INDEX IF NOT EXISTS idx_transactions_date ON transactions(transaction_date);
CREATE INDEX IF NOT EXISTS idx_transactions_type ON transactions(transaction_type);
CREATE INDEX IF NOT EXISTS idx_transactions_status ON transactions(status);

