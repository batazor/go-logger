# go-logger

Writes telemetry to TSDB

### Getting start

<details><summary>CLICK ME</summary>
<p>

```
# Run services
$ docker-compose up
```

</p>
</details>

### ENV

| Name                  | Default value                              |
|-----------------------|--------------------------------------------|
| GRPC_ENABLE           | true                                       |
| GRPC_PORT             | "50051"                                    |
| AMQP_ENABLE           | true                                       |
| AMQP_API              | amqp://telemetry:telemetry@localhost:5672/ |
| AMQP_NAME_QUEUE       | go-logger-packets                          |
| AMQP_EXCHANGE_LIST    | "demo1, demo2"                             |
| AMQP_EXCHANGE_TYPE    | "headers"                                  |
| DB_URL                | "http://influxdb:8086"                     |
| DB_NAME               | "telemetry"                                |
| DB_USERNAME           | "telemetry"                                |
| DB_PASSWORD           | "telemetry"                                |
| DB_ID                 | "_oid"                                     |
| PROMETHEUS_ENABLED    | "true"                                     |

### Feature

+ Support JSON data
    + Convert nested JSON to flat JSON
+ Support transports
    + AMQP (RabbitMQ) (Input data)
+ Grafana dashboard
+ Support
  + Kubernetes (Helm chart)
  + Prometheus metrics
  + GitLab CI

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
