# notifications

notifications is a scaffold of the platform's push service, which is to read notification
events from Kafka and deliver them to phones and browsers through APNs, FCM and Web
Push. Today it is a Go service with a chi router that answers `/healthz` and Prometheus
metrics at `/metrics`, and nothing more: the Kafka consumer and the three dispatchers in
`internal/` are stubs that nothing calls. No app and no other service talks to it, the
platform's chart does not run it, and the manifests in `k8s/` run no replica.
How a change is proposed and what a pull request has to pass is in the
[guide to contributing](https://github.com/zaentrum/.github/blob/main/CONTRIBUTING.md);
what may be built is in [building from the specification](https://github.com/zaentrum/zaentrum/blob/main/AGENTS.md).

## Build

```sh
go build ./...
docker build -t notifications .
```

Go 1.23 (`go.mod`). The image builds a static binary in `golang:1.23-alpine` and runs
it on `gcr.io/distroless/static-debian12:nonroot` as a non-root user, on port 8080.
`--build-arg BASE=<registry>/library/` pulls the build image from a mirror; the runtime
image is the same whatever `BASE` says. `build.yml` builds the image on every pull
request and on `main`; no workflow publishes it.

## Test

As `build.yml` runs them on a pull request and on `main`:

```sh
go build ./...
go vet ./...
go test ./...
docker build -t notifications:ci .
```

The test needs Go alone: it starts the server on a free port of `127.0.0.1`.

## Run it

```sh
go run ./cmd/server
curl -fsS http://localhost:8080/healthz
```

It reads `LISTEN_ADDR` alone (default `:8080`). `internal/config` reads `KAFKA_BROKERS`,
`KAFKA_GROUP_ID`, `KAFKA_TOPICS`, `APNS_KEY_PATH`, `APNS_TEAM_ID`, `APNS_KEY_ID`,
`FCM_CREDS_PATH` and `WEBPUSH_SUB_KEY`, but `main` does not call it. Nothing on the
reference demo talks to it, and there are no recorded answers to run it against.

## What the tests leave unchecked

- `cmd/server/main_test.go` starts the server, checks that `/healthz` answers 200, and
  that it stops on SIGTERM. Nothing more has a test.
- `/metrics`, and every package in `internal/`.
- The image: `build.yml` builds it, and nothing starts it.
- The manifests in `k8s/`: nothing applies or checks them.
