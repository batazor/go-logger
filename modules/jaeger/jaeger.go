package jaeger

import (
	"github.com/batazor/go-logger/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"io"
)

var (
	log = logrus.New()

	// ENV
	JAEGER_SERVICE_NAME    = utils.Getenv("JAEGER_SERVICE_NAME", "go-logger")
	JAEGER_AGENT_HOST_PORT = utils.Getenv("JAEGER_AGENT_HOST_PORT", "localhost:6831")
	JAEGER_RPC_METRICS     = utils.Getenv("JAEGER_RPC_METRICS", "true")
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

// Listen returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Listen() (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: JAEGER_SERVICE_NAME,
		RPCMetrics:  JAEGER_RPC_METRICS == "true",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: JAEGER_AGENT_HOST_PORT,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaegerlog.StdLogger))
	if err != nil {
		log.Errorf("ERROR: cannot init Jaeger: ", err)
	}

	log.Info("Run OpenTracing")

	return tracer, closer
}
