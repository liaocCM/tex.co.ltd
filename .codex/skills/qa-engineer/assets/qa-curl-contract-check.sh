#!/usr/bin/env bash
set -euo pipefail

BASE_URL="${BASE_URL:-http://localhost:8080}"
TOKEN="${TOKEN:-}"

headers=("-H" "Content-Type: application/json")
if [[ -n "$TOKEN" ]]; then
  headers+=("-H" "Authorization: Bearer ${TOKEN}")
fi

payload='{"limit":1}'

curl -sS -X POST "${headers[@]}" -d "$payload" "${BASE_URL}/reports/query" | jq .
