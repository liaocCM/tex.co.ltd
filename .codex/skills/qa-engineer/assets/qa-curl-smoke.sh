#!/usr/bin/env bash
set -euo pipefail

BASE_URL="${BASE_URL:-http://localhost:8080}"
TOKEN="${TOKEN:-}"

headers=()
if [[ -n "$TOKEN" ]]; then
  headers+=("-H" "Authorization: Bearer ${TOKEN}")
fi

curl -sS "${headers[@]}" "${BASE_URL}/health" | jq .
