
package database

import (
	"database/sql"
	"embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

// migrationsTableDDL cria a tabela de controle de versão do schema, caso
// ainda não exista. Esta tabela é a "fonte da verdade" de quais migrations
// já foram aplicadas — independente da workspace_metadata (que guarda
// metadados gerais do workspace, conforme migration 004).
const migrationsTableDDL = `
CREATE TABLE IF NOT EXISTS schema_migrations (
	version TEXT PRIMARY KEY,
	applied_at TEXT NOT NULL
);
`

// ApplyMigrations aplica, em ordem, todas as migrations embutidas que ainda
// não constam em schema_migrations. Comportamento conforme README seção 14:
//   - se o banco não existir, ele é criado (responsabilidade do Open)
//   - migrations pendentes são aplicadas automaticamente
//   - a versão do schema é registrada
func ApplyMigrations(db *sql.DB) error {
	if _, err := db.Exec(migrationsTableDDL); err != nil {
		return fmt.Errorf("falha ao criar tabela schema_migrations: %w", err)
	}

	applied, err := appliedVersions(db)
	if err != nil {
		return err
	}

	files, err := migrationFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		version := versionFromFilename(file)
		if applied[version] {
			continue
		}

		sqlBytes, err := migrationsFS.ReadFile("migrations/" + file)
		if err != nil {
			return fmt.Errorf("falha ao ler migration %s: %w", file, err)
		}

		tx, err := db.Begin()
		if err != nil {
			return err
		}

		if _, err := tx.Exec(string(sqlBytes)); err != nil {
			tx.Rollback()
			return fmt.Errorf("falha ao aplicar migration %s: %w", file, err)
		}

		if _, err := tx.Exec(
			"INSERT INTO schema_migrations (version, applied_at) VALUES (?, datetime('now'))",
			version,
		); err != nil {
			tx.Rollback()
			return fmt.Errorf("falha ao registrar migration %s: %w", file, err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("falha ao commitar migration %s: %w", file, err)
		}
	}

	return nil
}

func appliedVersions(db *sql.DB) (map[string]bool, error) {
	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applied := make(map[string]bool)
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		applied[v] = true
	}
	return applied, rows.Err()
}

func migrationFiles() ([]string, error) {
	entries, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return nil, err
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			files = append(files, e.Name())
		}
	}
	sort.Strings(files) // 001_..., 002_..., garante ordem correta
	return files, nil
}

func versionFromFilename(filename string) string {
	// "001_create_assets.sql" -> "001"
	parts := strings.SplitN(filename, "_", 2)
	return parts[0]
}

