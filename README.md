# go-logger

Writes telemetry to TSDB

### Getting start

```
docker-compose up
```

### ENV

| Name                  | Default value                      |
|-----------------------|------------------------------------|
| AMQP_API              | amqp://guest:guest@localhost:5672/ |
| AMQP_NAME_QUEUE       | go-logger-packets                  |
| AMQP_EXCHANGE_LIST    | "demo1, demo2"                     |
