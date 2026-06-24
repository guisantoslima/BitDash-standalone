
#!/usr/bin/env bash
set -euo pipefail

echo "==> Preparando ambiente Python (analytics/)"
cd analytics
if [ ! -d ".venv" ]; then
  python3 -m venv .venv
fi
source .venv/bin/activate
pip install -q -r requirements.txt
cd ..

echo "==> Subindo BitDash (Go)"
go run ./cmd/bitdash

