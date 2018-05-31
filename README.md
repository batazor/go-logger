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

| Name                  | Default value                      |
|-----------------------|------------------------------------|
| AMQP_ENABLE           | false                              |
| AMQP_API              | amqp://guest:guest@localhost:5672/ |
| AMQP_NAME_QUEUE       | go-logger-packets                  |
| AMQP_EXCHANGE_LIST    | "demo1, demo2"                     |
| DB_URL                | "http://localhost:8086"            |
| DB_NAME               | "telemetry"                        |
| DB_USERNAME           | "telemetry"                        |
| DB_PASSWORD           | "telemetry"                        |
