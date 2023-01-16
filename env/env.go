/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package env

import (
	"os"
	"strconv"
	"time"
)

// SentinelSvidPrefix returns the prefix for the Safe
// SVID (Short-lived Verifiable Identity Document) used in the Aegis system.
// The prefix is obtained from the environment variable
// AEGIS_SENTINEL_SVID_PREFIX. If the variable is not set, the default prefix is
// used.
func SentinelSvidPrefix() string {
	p := os.Getenv("AEGIS_SENTINEL_SVID_PREFIX")
	if p == "" {
		p = "spiffe://aegis.z2h.dev/workload/aegis-sentinel/ns/aegis-system/sa/aegis-sentinel/n/"
	}
	return p
}

// SafeSvidPrefix returns the prefix for the Safe
// SVID (Short-lived Verifiable Identity Document) used in the Aegis system.
// The prefix is obtained from the environment variable
// AEGIS_SAFE_SVID_PREFIX. If the variable is not set, the default prefix is
// used.
func SafeSvidPrefix() string {
	p := os.Getenv("AEGIS_SAFE_SVID_PREFIX")
	if p == "" {
		p = "spiffe://aegis.z2h.dev/workload/aegis-safe/ns/aegis-system/sa/aegis-safe/n/"
	}
	return p
}

// WorkloadSvidPrefix returns the prefix for the Workload SVID
// (SPIFFE Verifiable Identity Document) used in the Aegis system.
// The prefix is obtained from the environment variable AEGIS_WORKLOAD_SVID_PREFIX.
// If the variable is not set, the default prefix is used.
func WorkloadSvidPrefix() string {
	p := os.Getenv("AEGIS_WORKLOAD_SVID_PREFIX")
	if p == "" {
		p = "spiffe://aegis.z2h.dev/workload/"
	}
	return p
}

// SpiffeSocketUrl returns the URL for the SPIFFE endpoint socket used in the
// Aegis system.
// The URL is obtained from the environment variable SPIFFE_ENDPOINT_SOCKET.
// If the variable is not set, the default URL is used.
func SpiffeSocketUrl() string {
	p := os.Getenv("SPIFFE_ENDPOINT_SOCKET")
	if p == "" {
		p = "unix:///spire-agent-socket/agent.sock"
	}
	return p
}

// SafeEndpointUrl returns the URL for the Safe endpoint used in the Aegis system.
// The URL is obtained from the environment variable AEGIS_SAFE_ENDPOINT_URL.
// If the variable is not set, the default URL is used.
func SafeEndpointUrl() string {
	u := os.Getenv("AEGIS_SAFE_ENDPOINT_URL")
	if u == "" {
		u = "https://aegis-safe.aegis-system.svc.cluster.local:8443/"
	}
	return u
}

// ProbeLivenessPort returns the port for liveness probe.
// It first checks the environment variable AEGIS_PROBE_LIVENESS_PORT.
// If the variable is not set, it returns the default value ":8081".
func ProbeLivenessPort() string {
	u := os.Getenv("AEGIS_PROBE_LIVENESS_PORT")
	if u == "" {
		u = ":8081"
	}
	return u
}

// ProbeReadinessPort returns the port for readiness probe.
// It first checks the environment variable AEGIS_PROBE_READINESS_PORT.
// If the variable is not set, it returns the default value ":8082".
func ProbeReadinessPort() string {
	u := os.Getenv("AEGIS_PROBE_READINESS_PORT")
	if u == "" {
		u = ":8082"
	}
	return u
}

// SidecarSecretsPath returns the path to the secrets file used by the sidecar.
// The path is determined by the AEGIS_SIDECAR_SECRETS_PATH environment variable,
// with a default value of "/opt/aegis/secrets.json" if the variable is not set.
func SidecarSecretsPath() string {
	p := os.Getenv("AEGIS_SIDECAR_SECRETS_PATH")
	if p == "" {
		p = "/opt/aegis/secrets.json"
	}
	return p
}

// SentryPollInterval returns the polling interval for sentry in time.Duration
// The interval is determined by the AEGIS_SIDECAR_POLL_INTERVAL environment
// variable, with a default value of 20 seconds if the variable is not set or
// if there is an error in parsing the value.
func SentryPollInterval() time.Duration {
	p := os.Getenv("AEGIS_SIDECAR_POLL_INTERVAL")
	if p == "" {
		p = "20"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 20 * time.Second
	}
	return time.Duration(i) * time.Second
}
