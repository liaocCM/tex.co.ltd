# Runbook

## Operations
- Backend listens on `127.0.0.1:8080`.
- Frontend dev server proxies `/api` to backend.

## Rollback
- Stop services and revert to previous commit.

## Monitoring
- Watch backend logs for HTTP errors.

## Common Failures
- SQLite file permission issues in `data/`.
- Backend not running when frontend makes `/api` calls.
