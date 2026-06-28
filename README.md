# notifications

Push-notification fan-out service for the zaentrum platform. Subscribes to
Kafka topics under `stube.notify.user.*` and dispatches deliveries to APNs
(iOS), FCM (Android), and Web Push (browser PWAs).

## Status

**Scaffold; build deferred until the mobile clients ship.**

The Kafka consumer, APNs / FCM / Web Push dispatchers, and credential wiring
are all stubs. The HTTP server compiles and serves `/healthz` + `/metrics` so
the service can be deployed as a 0-replica placeholder.

## Layout

```
cmd/server/main.go                # chi router, /healthz, /metrics (no Kafka loop yet)
internal/config/config.go         # env-driven config
internal/apns/client.go           # APNs stub (needs .p8 key + Team ID + Key ID)
internal/fcm/client.go            # FCM stub (needs service-account creds)
internal/webpush/client.go        # web push stub (needs VAPID key pair)
internal/kafka/consumer.go        # sarama consumer-group stub
k8s/                              # Deployment (replicas: 0), Service, ServiceMonitor, GrafanaDashboard
Dockerfile
```

## Local development

```bash
go run ./cmd/server
curl -fsS http://localhost:8080/healthz
```

## Build the container

```bash
docker build -t zaentrum/notifications .
```

The `k8s/` manifests ship with `replicas: 0`; scale up once notification
delivery is wired. Build and push the image to your own registry and update
the image reference for your environment.

## License

[MPL-2.0](LICENSE).
