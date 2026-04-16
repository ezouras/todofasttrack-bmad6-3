# Deferred Work

This file collects findings from code reviews that were deemed real but out of scope for the story under review. Pull from this list when planning future stories or hardening passes.

## Deferred from: code review of story 1-1-monorepo-scaffold-ci-pipeline (2026-04-16)

- **Gin server hardening** — `apps/api/cmd/server/main.go`: add `http.Server` with `ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`, `IdleTimeout`. Wire SIGINT/SIGTERM to `srv.Shutdown(ctx)` for graceful shutdown so in-flight requests aren't dropped on container stop. Suggested target: Story 1.2 (deployment) or Story 1.4 (first auth-protected routes).
- **`PORT` env validation** — `main.go`: parse with `strconv.Atoi`, reject non-numeric / out-of-range / `PORT=0` / strings with whitespace or shell metacharacters. Suggested target: Story 1.2.
- **Distinct EADDRINUSE error** — surface a clearer message ("port 8080 already in use; another `pnpm dev` running?") instead of generic listen error. Suggested target: Story 1.2.
- **Production Gin mode** — set `gin.SetMode(gin.ReleaseMode)` when `GIN_MODE` is unset to avoid debug logs leaking in deployed envs. Suggested target: Story 1.4.
- **Serve generated swagger docs** — `apps/api/docs/docs.go` is committed and tracked but never imported, so it's dead code in production builds. Either wire it up via `gin-swagger` at `/swagger/*` (the canonical pattern) or stop tracking. Suggested target: Story 1.4 when first real API endpoints exist.
- **`.gitignore` consistency** — root `.gitignore` blocks `.env*` broadly, `apps/mobile/.gitignore` blocks only `.env*.local`. Unify so all workspaces follow the same convention.
- **Windows-compatible `generate:types`** — current script uses POSIX `$(go env GOPATH)`. Either gate on platform, document Mac/Linux requirement, or rewrite as a Node script that handles both.
- **ESLint catalog/dedupe** — `apps/web` and `apps/mobile` both pin `eslint: ^9.39.4` independently. Move to pnpm `catalog:` so a bump in one workspace doesn't drift the other.
- **`go.work` toolchain mismatch warning** — `go.work` pins `go 1.26.2`. A contributor with older Go gets a cryptic "go: requires newer Go" error. Either pin to 1.22+ explicitly with `toolchain` directive, or add a clear "requires Go 1.26.2+" line in `README.md`.
- **`slog` log level configurable** — wire `LOG_LEVEL` or `DEBUG` env to `slog.NewJSONHandler`'s level option. Suggested target: Story 1.4 (observability scope).
- **Tame orphan `go run` processes** — `pnpm dev` Ctrl-C doesn't always reap the spawned `go run`-launched child. Common Turborepo + Go pattern. Either build a binary once and exec it, or use `air`/`reflex` with proper signal forwarding. Tooling improvement story.
- **Stronger smoke test** — `main_test.go` doesn't validate `Content-Type` before `json.Unmarshal`. If the handler accidentally returned HTML, the test would fail with a confusing parse error rather than a clear `Content-Type` assertion. Low priority.
- **`PORT=0` ephemeral port** — handled at OS layer; if surface this as a feature, the bound port needs to be logged so callers can find it. Currently silent.
- **Isolate `swag/v2` from production module graph (P15 from review)** — `apps/api/docs/docs.go` (auto-regenerated) imports `github.com/swaggo/swag/v2` unconditionally, which keeps swag + transitive deps (mongo-driver, quic-go, sv-tools/openapi) in the main `require` block of `go.mod`. Linker dead-code elimination strips them from the binary, but `go mod download` still fetches them in CI. Proper fix requires postprocessing `docs.go` after every `swag init` to prepend a build tag (e.g. `//go:build swagger`) so the file is excluded by default builds. Not a quick edit because the postprocess step lives in the generation script and the generated file would re-overwrite without it. Suggested target: a focused tooling hardening story.
