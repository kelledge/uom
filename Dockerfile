FROM golang:1.24.5-bookworm AS dev

WORKDIR /lib

# Pin tool versions (override with --build-arg as needed)
ARG GOPLS_VERSION=v0.20.0
ARG GOLANGCI_LINT_VERSION=v2.4.0
ARG GOTESTSUM_VERSION=v1.12.3

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOBIN=/usr/local/bin go install golang.org/x/tools/gopls@${GOPLS_VERSION}

# See warning: https://golangci-lint.run/docs/welcome/install/#install-from-sources
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOBIN=/usr/local/bin go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOBIN=/usr/local/bin go install gotest.tools/gotestsum@${GOTESTSUM_VERSION}

# Prime module cache for faster builds
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download