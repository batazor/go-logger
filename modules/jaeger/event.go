package jaeger

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func Add(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "say-hello")
	defer span.Finish()

	// Add tag
	span.SetTag("hello-to", Event{
		event: "HelloWorld",
	})

	span.LogKV("event", "println")
}
