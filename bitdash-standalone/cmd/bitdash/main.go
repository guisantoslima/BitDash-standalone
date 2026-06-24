
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"bitdash-standalone/internal/config"
	"bitdash-standalone/internal/database"
	"bitdash-standalone/internal/logging"
	"bitdash-standalone/internal/web"
	"bitdash-standalone/internal/workspace"
)

func main() {
	migrateOnly := flag.Bool("migrate-only", false, "aplica migrations e sai, sem subir o servidor")
	flag.Parse()

	cfg := config.LoadAppConfig()

	// Passo 1 (Blueprint 7.1) — carregar configuração local do host e
	// localizar último workspace usado.
	runtimeCfg, err := config.LoadRuntimeConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "falha ao carregar runtime config: %v\n", err)
		os.Exit(1)
	}

	workspacePath := resolveWorkspacePath(cfg, runtimeCfg)

	// Passo 2/3/4 (Blueprint 7.1) — validar workspace; se inválido, este
	// Sprint 0 já cria a estrutura padrão automaticamente (o fluxo de
	// onboarding interativo via UI será refinado no Sprint 5).
	mgr := workspace.NewManager()
	if err := mgr.InitializeWorkspace(workspacePath); err != nil {
		fmt.Fprintf(os.Stderr, "falha ao inicializar workspace em %q: %v\n", workspacePath, err)
		os.Exit(1)
	}

	// Persiste o workspace usado para a próxima execução.
	runtimeCfg.LastWorkspacePath = workspacePath
	if err := config.SaveRuntimeConfig(runtimeCfg); err != nil {
		fmt.Fprintf(os.Stderr, "aviso: falha ao salvar runtime config: %v\n", err)
	}

	// Logger estruturado, escrevendo dentro do workspace já garantido.
	logger, err := logging.New(workspace.LogsPath(workspacePath), cfg.Debug)
	if err != nil {
		fmt.Fprintf(os.Stderr, "falha ao inicializar logger: %v\n", err)
		os.Exit(1)
	}

	logger.Info("workspace_loaded", "path", workspacePath)

	// Passo 4 (Blueprint 7.1) — inicializar banco e aplicar migrations.
	db, err := database.Open(workspace.DBPath(workspacePath))
	if err != nil {
		logger.Error("database_open_failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := database.ApplyMigrations(db); err != nil {
		logger.Error("migrations_failed", "error", err)
		os.Exit(1)
	}
	logger.Info("migrations_applied")

	if *migrateOnly {
		logger.Info("migrate_only_flag_set_exiting")
		return
	}

	// Passo 5 (Blueprint 7.1) — subir servidor local, bind apenas em
	// 127.0.0.1 (ADR de segurança / README seção 19).
	renderer := web.NewRenderer("web/templates")
	router := web.NewRouter(logger, renderer)

	addr := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	logger.Info("server_starting", "addr", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Error("server_failed", "error", err)
		os.Exit(1)
	}
}

// resolveWorkspacePath decide qual pasta usar como workspace, na seguinte
// ordem de prioridade (Blueprint 7.1 Passo 1 + README 13):
//  1. variável de ambiente BITDASH_WORKSPACE_PATH (override explícito)
//  2. último workspace usado, persistido em runtime config
//  3. pasta padrão do sistema operacional (~/BitDashData)
func resolveWorkspacePath(cfg config.AppConfig, runtimeCfg config.RuntimeConfig) string {
	if cfg.WorkspacePath != "" {
		return cfg.WorkspacePath
	}
	if runtimeCfg.LastWorkspacePath != "" {
		return runtimeCfg.LastWorkspacePath
	}

	defaultPath, err := workspace.DefaultWorkspacePath()
	if err != nil {
		// Fallback extremo: pasta atual + BitDashData
		return "./BitDashData"
	}
	return defaultPath
}

