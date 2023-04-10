# syntax=docker/dockerfile:1

ARG GO_VERSION="1.18.1"
ARG RUNNER_IMAGE="gcr.io/distroless/static-debian11"

# --------------------------------------------------------
# Builder
# --------------------------------------------------------

FROM golang:${GO_VERSION}-alpine as builder

ARG GIT_VERSION
ARG GIT_COMMIT

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers

# Download go dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download


# Copy the remaining files
COPY . .

# Build fanx binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    GOWORK=off go build \
        -mod=readonly \
        -tags "netgo,ledger,muslc" \
        -ldflags \
            "-X github.com/cosmos/cosmos-sdk/version.Name="fanx" \
            -X github.com/cosmos/cosmos-sdk/version.AppName="fanx" \
            -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
            -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
            -X github.com/cosmos/cosmos-sdk/version.BuildTags='netgo,ledger,muslc' \
            -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
        -trimpath \
        -o /app/build/fanx \
        /app/cmd/fanx

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM ${RUNNER_IMAGE}

COPY --from=builder /app/build/fanx /bin/fanx

ENV HOME /fanx
WORKDIR $HOME

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317

ENTRYPOINT ["fanx"]
