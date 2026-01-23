# Frontend (React + TypeScript)

## Setup

- Requires Node.js 18+

## Run

```bash
npm install
npm run dev
```

The dev server proxies `/api` to `http://127.0.0.1:8080`.

## Notes
- API base is `/api` via Vite proxy.
- Keep UI minimal and aligned to MVP scope.
- Client stores a per-browser uid in localStorage under `todo_uid` and sends it as `X-User-Uid`.
