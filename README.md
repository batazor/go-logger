# go-logger

Writes telemetry to TSDB

### Getting start

```
# Run services
$ docker-compose up

# Create database on influx
$ docker exec -it influxdb influx
> create database telemetry
> exit
```

### ENV

| Name                  | Default value                              |
|-----------------------|--------------------------------------------|
| GRPC_ENABLE           | true                                       |
| GRPC_PORT             | "50051"                                    |
| AMQP_ENABLE           | false                                      |
| AMQP_API              | amqp://telemetry:telemetry@localhost:5672/ |
| AMQP_NAME_QUEUE       | go-logger-packets                          |
| AMQP_EXCHANGE_LIST    | "demo1, demo2"                             |
| DB_URL                | "http://localhost:8086"                    |
| DB_NAME               | "telemetry"                                |
| DB_USERNAME           | "telemetry"                                |
| DB_PASSWORD           | "telemetry"                                |
| DB_ID                 | "_oid"                                     |

### Feature

+ Support transports
    + AMQP (RabbitMQ)
+ Grafana dashboard

#### Grafana

<details><summary>CLICK ME</summary>
<p>

##### Grafana dashboard example

![Grafana dashboard example](./docs/grafana-example-dashboard.png)

**Support:**
- Group by ID object

</p>
</details>

### Benchmark

##### Run bot

Run `go run /tests/bot/bot.go`

##### Read from AMQP queue (1M message/1 instance)

![read_packets](./docs/read_packet.png)