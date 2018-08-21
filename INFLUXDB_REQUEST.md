# Get last state as time range

`SELECT last("*") FROM "telemetry"."autogen"."<ObjectId>" WHERE time > now() - 24h AND time < now() - 23h LIMIT 1`

# Get count points
`SELECT count("_oid") FROM "telemetry"."autogen"./^*/`
`SELECT count("_oid") from "telemetry"."autogen"."/^*/""`

# Get size each DB
```
SELECT last("diskBytes") FROM "_internal"."monitor"."shard" GROUP BY "database"
```