# syntax = docker/dockerfile:1.4
FROM --platform=$BUILDPLATFORM golang:alpine3.17 as builder
WORKDIR /src
RUN apk --update --no-cache add git
ENV GO111MODULE=on
COPY . .
RUN ls -la /src

RUN go mod download && GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 go build -o semver-cli -ldflags '-w -extldflags "-static"' .

FROM scratch
COPY --from=builder /src/semver-cli /
ENTRYPOINT ["/semver-cli"]