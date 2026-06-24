
package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" // driver pure-Go, sem CGO (Blueprint seção 40)
)

// Open abre (ou cria, se não existir) o arquivo SQLite no caminho informado
// e configura os pragmas recomendados:
//   - foreign_keys=ON: garante integridade referencial (assets <-> transactions)
//   - journal_mode=WAL: melhor concorrência leitura/escrita para app local
func Open(dbPath string) (*sql.DB, error) {
	dsn := fmt.Sprintf("file:%s?_pragma=foreign_keys(1)", dbPath)

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if _, err := db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		return nil, err
	}
	if _, err := db.Exec("PRAGMA foreign_keys=ON;"); err != nil {
		return nil, err
	}

	return db, nil
}

