# axhub-282-hello-compose

Minimal Docker Compose project for AX Hub spec 282 e2e validation.

- `web` — Go HTTP server, depends on Redis
- `redis` — `redis:7-alpine` with named volume

The `apphub.yaml` declares `deploy_method=compose`, `resource_tier=S`.
