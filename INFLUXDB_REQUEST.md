# Get last state as time range

`SELECT last("*") FROM "telemetry"."autogen"."<ObjectId>" WHERE time > now() - 24h AND time < now() - 23h LIMIT 1`