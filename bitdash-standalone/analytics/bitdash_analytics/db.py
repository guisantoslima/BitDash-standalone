
"""
Conexão read-only com o SQLite do workspace.

Nesta fase (Sprint 0) apenas o stub de conexão está implementado. A lógica
de queries para dashboard/portfolio será adicionada no Sprint 3.
"""

import sqlite3
from pathlib import Path


def connect(workspace_path: str) -> sqlite3.Connection:
    """
    Abre uma conexão read-only com o bitdash.db dentro do workspace
    informado. Lança FileNotFoundError se o banco não existir.
    """
    db_path = Path(workspace_path) / "bitdash.db"
    if not db_path.exists():
        raise FileNotFoundError(f"bitdash.db não encontrado em {workspace_path}")

    # mode=ro: a engine Python NUNCA escreve no banco oficial (Go é o
    # único responsável por escritas, conforme ADR-003 do Blueprint).
    uri = f"file:{db_path}?mode=ro"
    return sqlite3.connect(uri, uri=True)

