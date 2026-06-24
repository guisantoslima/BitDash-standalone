
"""
Stub do script de cálculo de dashboard.

Contrato definido no Blueprint seção 15:
  python calc_dashboard.py --workspace "<path>" --from "<YYYY-MM-DD>" --to "<YYYY-MM-DD>"

Neste Sprint 0, o script apenas valida os argumentos e devolve uma
estrutura JSON vazia, porém válida quanto ao schema esperado (Blueprint
seção 26 / 15). A lógica real de cálculo (summary, by_asset, séries
mensais) será implementada no Sprint 3.
"""

import argparse
import json
import sys


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Calcula dashboard do BitDash")
    parser.add_argument("--workspace", required=True, help="Caminho do workspace BitDashData")
    parser.add_argument("--from", dest="date_from", required=True, help="Data inicial (YYYY-MM-DD)")
    parser.add_argument("--to", dest="date_to", required=True, help="Data final (YYYY-MM-DD)")
    return parser.parse_args()


def empty_dashboard_payload() -> dict:
    """Estrutura mínima válida, conforme contrato Go<->Python do Blueprint."""
    return {
        "summary": {
            "total_entries": 0.0,
            "total_withdrawals": 0.0,
            "net_balance": 0.0,
            "transaction_count": 0,
        },
        "by_asset": [],
        "monthly_entries": [],
        "monthly_withdrawals": [],
    }


def main() -> None:
    args = parse_args()

    # Sprint 0: apenas confirma que workspace foi recebido; não lê o banco
    # ainda (isso é responsabilidade do Sprint 3, junto com a lógica real
    # em bitdash_analytics/dashboard.py e portfolio.py).
    payload = empty_dashboard_payload()

    print(json.dumps(payload))


if __name__ == "__main__":
    main()