# BitDash-standalone
Aplicação local-first para gestão de lançamentos e retiradas de criptoativos, com persistência local na máquina do usuário ou em uma pasta estruturada reconhecida pelo app, a aplicação contém dashboards analíticos e operação 100% standalone.

## Objetivo
BitDash-standalone é uma aplicação local para controle e análise de movimentações de criptoativos, tendo como objetivo registrar lançamentos e retiradas, consolidar informações por ativo e visualizar indicadores e gráficos sem depender de backend externo ou armazenamento em nuvem.

Todos os dados são mantidos localmente, em banco SQLite e/ou em uma estrutura de pasta reconhecida pelo aplicativo, priorizando simplicidade operacional, criptografia, portabilidade e privacidade.

## Versão 1 - MVP

### Motor de persistência do standalone

### Modelo Escolhido:
Persistência principal em **SQLite**. Armazenado dentro de uma pasta C:\\BitDashData ou ~/BitDashData com possibilidade de backup/export/import de JSON/CSV.

```
Estrutura padrão

BitDashData/
├─ bitdash.db                 # banco SQLite principal
├─ bitdash.config.json        # configurações do app
├─ backups/
│  ├─ backup-YYYY-MM-DD.json
│  └─ backup-YYYY-MM-DD.db
├─ exports/
│  ├─ transactions-YYYY-MM-DD.csv
│  └─ dashboard-YYYY-MM-DD.json
├─ logs/
│  └─ bitdash.log
└─ temp/
```

Na primeira execução o app deve:

verificar se já existe configuração local e roda-la, se não existir:

Oferecer “Usar pasta padrão do BitDash” ou “Escolher pasta manualmente”
```
criar a estrutura da pasta
criar o bitdash.db
criar o bitdash.config.json
```

Quando o usuário aponta uma pasta manual o app deve:

Cenário A — pasta vazia
```
criar a estrutura padrão BitDash nela
```

Cenário B — pasta já contém estrutura BitDash válida
```
Abrir em contexto com o local escolhido
```
Cenário C — pasta existe mas não está no padrão

Informar que o arquivo está fora do padrão e Oferecer: 
```
inicializar BitDash ali ou cancelar
```

### Planejamento Arquitetural

Go como host da aplicação + Python como motor analítico

**O papel do Go**

O Go vira o “casco” do BitDash-standalone:
```
- sobe o servidor local
- expõe as páginas no browser
- lê/escreve o workspace local
- faz CRUD de ativos / lançamentos / retiradas
- expõe endpoints para dashboard
- chama o Python (Cálculos pesados ou geração de relatórios)
```
**O papel do Python**

O Python vira o “motor de inteligência analítica”:
```
- cálculo de preço médio
- agregações mensais
- consolidação por ativo
- relatórios CSV/Excel/PDF
- preparação de datasets de gráficos
- eventualmente simulações e indicadores mais sofisticados
```
**Arquitetura “dual runtime”**

Adotar Go como runtime principal da aplicação standalone e Python como engine analítica para relatórios e dashboards, usando SQLite como armazenamento local oficial e um workspace local BitDashData como diretório estruturado do produto.

Camada 1 — BitDash Web Host (Go)

Responsável por:

- subir servidor local (localhost:8080)
- renderizar UI no browser
- CRUD de ativos e transações
- acessar SQLite
- gerenciar workspace
- disparar geração de relatórios Python
- servir assets estáticos

Camada 2 — BitDash Analytics Engine (Python)

Responsável por:

- calcular métricas
- consolidar séries temporais
- gerar datasets do dashboard
- exportar relatórios
- recalcular snapshots
- futuras análises mais sofisticadas

Camada 3 — Storage

- SQLite como fonte oficial
- `BitDashData/` como workspace
- `exports/`, `backups/`, `logs/`

### Papéis e Responsábilidades

Go faz:

```
GET /dashboard
GET /transactions
POST /transactions
POST /withdrawals
GET /assets
GET /settings
POST /workspace/select
POST /backup/export
```
Python faz:
```
calculate_dashboard_summary(workspace, period)
calculate_monthly_entries(workspace, period)
calculate_monthly_withdrawals(workspace, period)
calculate_average_price_by_asset(workspace)
export_transactions_csv(workspace, filters)
```

### Stack-alvo

Go → aplicação principal, servidor local, CRUD, UI via browser, gestão do workspace

Python → analytics, consolidação, cálculo de indicadores, relatórios/exportações

SQLite → persistência local oficial

Browser local → interface do usuário

Workspace local BitDashData → diretório padrão do app


## Funcionalidades:

```
Funcionalidades Esperadas para V1-MVP

[✓] Rode o app localmente
[✓] Cadastre lançamentos e retiradas
[✓] Salve tudo na própria máquina
[✓] Pasta estruturada padrão do BitDash
[✓] Visualize um dashboard com indicadores e gráficos
[✓] Faça backup/restauração dos dados de forma simples
```





