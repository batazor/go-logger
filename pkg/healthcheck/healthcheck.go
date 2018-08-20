package healthcheck

import (
	"github.com/heptiolabs/healthcheck"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	log    = logrus.New()
	Health healthcheck.Handler
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)

	Health = healthcheck.NewHandler()
}

// Run prometheus exporter
func Listen() {
	log.Info("Run HealthCheck")
	time.Sleep(time.Second * 5)

	// Our app is not happy if we've got more than 100 goroutines running.
	Health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(300))

	// Our app is not ready if we can't resolve our upstream dependency in DNS.
	Health.AddReadinessCheck(
		"upstream-dep-dns",
		healthcheck.DNSResolveCheck("google.com", 50*time.Millisecond))

	// Our app is not ready if we can't connect to our database (`var db *sql.DB`) in <1s.
	//health.AddReadinessCheck("database", healthcheck.DatabasePingCheck(db, 1*time.Second))

	go http.ListenAndServe("0.0.0.0:8086", Health)
}
