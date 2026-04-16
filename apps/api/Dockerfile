# syntax=docker/dockerfile:1.7

# --- Build stage -------------------------------------------------------------
FROM golang:1.26-alpine AS builder

WORKDIR /src

# Cache module downloads independently of source changes
COPY go.mod go.sum ./
RUN go mod download

# Build the static binary. CGO_ENABLED=0 is required because the runtime image
# is distroless/static and has no glibc.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/tend-api \
    ./cmd/server

# --- Runtime stage -----------------------------------------------------------
FROM gcr.io/distroless/static-debian12:nonroot

# distroless/static is ~2 MB and ships only the cert bundle + nonroot user.
# No shell, no package manager, no glibc — minimal attack surface.

COPY --from=builder /out/tend-api /tend-api

EXPOSE 8080

# nonroot user is UID 65532 — already set by the distroless base image.
ENTRYPOINT ["/tend-api"]
