# Base-image registry prefix. Empty default = public Docker Hub; a private
# deploy mirror passes --build-arg BASE=registry.example/library/ .
ARG BASE=
FROM ${BASE}golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/notifications ./cmd/server

FROM ${BASE}distroless/static-debian12:nonroot
COPY --from=build /out/notifications /notifications
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/notifications"]
