#!/usr/bin/env bash
# Regenerate the OpenAPI spec from Go annotations and the TypeScript types from the spec.
# Run from the repo root via `pnpm run generate:types`.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

SWAG="$(go env GOPATH)/bin/swag"
if [ ! -x "$SWAG" ]; then
  echo "❌ swag CLI not found at $SWAG"
  echo "   Install with: go install github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc5"
  exit 1
fi

# Step 1: Go annotations → OpenAPI 3.1 spec (in a subshell so cwd is preserved)
(
  cd "$REPO_ROOT/apps/api"
  "$SWAG" init -g cmd/server/main.go -o ./docs --v3.1
)

# Step 2: Verify the spec is non-empty before consuming it
SPEC="$REPO_ROOT/apps/api/docs/swagger.json"
if [ ! -s "$SPEC" ]; then
  echo "❌ Generated $SPEC is missing or empty — swag failed silently."
  exit 1
fi

# Step 3: OpenAPI spec → TypeScript types (use pnpm exec for reliable bin resolution)
cd "$REPO_ROOT"
pnpm exec openapi-typescript "$SPEC" -o packages/types/generated/api.ts

echo "✅ Types regenerated: packages/types/generated/api.ts"
