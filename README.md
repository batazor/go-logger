# go-logger

Writes telemetry to TSDB [![GoDoc][doc-img]][doc] [![OpenTracing Badge](https://img.shields.io/badge/OpenTracing-enabled-blue.svg)](http://opentracing.io)

### Feature

+ Support JSON data
    + Convert nested JSON to flat JSON
+ Support transports
    + AMQP (RabbitMQ) (Input data)
+ API
  + gRPC
+ Grafana dashboard
+ Opentracing
+ Support
  + Kubernetes (Helm chart)
    + Healthcheck implementing Kubernetes liveness and readiness probe
  + Prometheus metrics
    + Grafana dashboard (https://grafana.com/dashboards/240)
  + GitLab CI

### Getting start

<details><summary>CLICK ME</summary>
<p>

```
# Generate gRPC code
go get -u github.com/golang/protobuf/proto
protoc --go_out=plugins=grpc:. pb/*.proto

# Run services
$ docker-compose up
```

</p>
</details>

### ENV

| Name                             | Default value                              |
|----------------------------------|--------------------------------------------|
| GRPC_ENABLE                      | true                                       |
| GRPC_PORT                        | "50051"                                    |
| AMQP_ENABLE                      | true                                       |
| AMQP_API                         | amqp://telemetry:telemetry@localhost:5672/ |
| AMQP_NAME_QUEUE                  | go-logger-packets                          |
| AMQP_EXCHANGE_LIST               | "demo1, demo2"                             |
| AMQP_EXCHANGE_TYPE               | "headers"                                  |
| DB_URL                           | "http://influxdb:8086"                     |
| DB_NAME                          | "telemetry"                                |
| DB_USERNAME                      | "telemetry"                                |
| DB_PASSWORD                      | "telemetry"                                |
| DB_ID                            | "_oid"                                     |
| PROMETHEUS_ENABLED               | "true"                                     |
| OPENTRACING_ENABLED              | "true"                                     |
| JAEGER_SERVICE_NAME              | go-logger                                  |
| JAEGER_AGENT_HOST_PORT           | "localhost:5778"                           |
| JAEGER_RPC_METRICS               | "true"                                     |

#### Grafana

<details><summary>CLICK ME</summary></details>

##### Grafana dashboard example

![Grafana dashboard example](./docs/grafana-example-dashboard.png)

**Support:**
- Group by ID object

</p>
</details>

#### Prometheus

<details><summary>CLICK ME</summary>

Prometheus metrics `localhost:9090/metrics`

Prometheus metrics:
- Basic metrics
</details>

#### GitLAb CI

This project support GitLab CI

<details><summary>CLICK ME</summary>
<p>

| Name                  | Description                                |
|-----------------------|--------------------------------------------|
| DOCKER_PASS           | --                                         |
| DOCKER_USER           | --                                         |
| GITHUB_PASSWORD       | --                                         |
| GITHUB_REPOSITORY_URL | --                                         |
| GITHUB_USER           | --                                         |
| HELM_CONTEXT          | --                                         |
| PROJECT_NAMESPACE     | --                                         |
| INFLUXDB_NODE_SELECT  | --                                         |

</p>
</details>


### Benchmark

<details><summary>CLICK ME</summary>
<p>

##### Run bot

Run `go run /tests/bot/bot.go`

##### Read from AMQP queue (1M message/1 instance)

![read_packets](./docs/read_packet.png)

</p>
</details>

[doc-img]: https://godoc.org/github.com/batazor/go-logger?status.svg
[doc]: https://godoc.org/github.com/batazor/go-logger
