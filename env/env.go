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

// SidecarPollInterval returns the polling interval for sentry in time.Duration
// The interval is determined by the AEGIS_SIDECAR_POLL_INTERVAL environment
// variable, with a default value of 20000 milliseconds if the variable is not
// set or if there is an error in parsing the value.
func SidecarPollInterval() time.Duration {
	p := os.Getenv("AEGIS_SIDECAR_POLL_INTERVAL")
	if p == "" {
		p = "20000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 20000 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}

// SafeSvidRetrievalTimeout returns the allowed time for Aegis Safe to wait
// before killing the pod to retrieve an SVID, in time.Duration.
// The interval is determined by the AEGIS_SAFE_SVID_RETRIEVAL_TIMEOUT environment
// variable, with a default value of 30000 milliseconds if the variable is not
// set or if there is an error in parsing the value.
func SafeSvidRetrievalTimeout() time.Duration {
	p := os.Getenv("AEGIS_SAFE_SVID_RETRIEVAL_TIMEOUT")
	if p == "" {
		p = "30000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 30000 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}

// SafeTlsPort returns the secure port for Aegis Safe to listen on.
// It checks the AEGIS_SAFE_TLS_PORT environment variable. If the variable
// is not set, it defaults to ":8443".
func SafeTlsPort() string {
	p := os.Getenv("AEGIS_SAFE_TLS_PORT")
	if p == "" {
		p = ":8443"
	}
	return p
}

// SafeDataPath returns the path to the safe data directory.
// The path is determined by the AEGIS_SAFE_DATA_PATH environment variable.
// If the environment variable is not set, the default path "/data" is returned.
func SafeDataPath() string {
	p := os.Getenv("AEGIS_SAFE_DATA_PATH")
	if p == "" {
		p = "/data"
	}
	return p
}

// SafeAgeKeyPath returns the path to the safe age key directory.
// The path is determined by the AEGIS_SAFE_AGE_KEY_PATH environment variable.
// If the environment variable is not set, the default path "/key/key.txt"
// is returned.
func SafeAgeKeyPath() string {
	p := os.Getenv("AEGIS_SAFE_AGE_KEY_PATH")
	if p == "" {
		p = "/key/key.txt"
	}
	return p
}

// LogLevel returns the value set by AEGIS_LOG_LEVEL environment
// variable, or a default level.
//
// AEGIS_LOG_LEVEL determines the verbosity of the logs.
// 1: logs are off, 6: highest verbosity.
// Off = 1, Error = 2, Warn = 3, Info = 4, Debug = 5, Trace = 6
func LogLevel() int {
	p := os.Getenv("AEGIS_LOG_LEVEL")
	if p == "" {
		return 3
	}
	l, _ := strconv.Atoi(p)
	if l == 0 {
		return 3
	}
	if l < 0 || l > 6 {
		return 3
	}
	return l
}

// SafeSecretBufferSize returns the buffer size for the Aegis Safe secret queue.
//
// The buffer size is determined by the environment variable
// AEGIS_SAFE_SECRET_BUFFER_SIZE.
//
// If the environment variable is not set, the default buffer size is 10.
// If the environment variable is set and can be parsed as an integer,
// it will be used as the buffer size.
// If the environment variable is set but cannot be parsed as an integer,
// the default buffer size is used.
func SafeSecretBufferSize() int {
	p := os.Getenv("AEGIS_SAFE_SECRET_BUFFER_SIZE")
	if p == "" {
		return 10
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 10
	}
	return l
}

// SafeBackingStoreType returns the storage type for the data,
// as specified in the AEGIS_SAFE_BACKING_STORE_TYPE environment variable.
// If the environment variable is not set, it defaults to "persistent".
// Any value that is not "persistent" will mean Aegis Safe will store
// its state in-memory
func SafeBackingStoreType() string {
	s := os.Getenv("AEGIS_SAFE_BACKING_STORE")
	if s == "" {
		return "persistent"
	}
	return s
}

// SafeSecretBackupCount retrieves the number of backups to keep for Aegis Safe
// secrets. If the environment variable AEGIS_SAFE_SECRET_BACKUP_COUNT is not
// set or is not a valid integer, the default value of 3 will be returned.
func SafeSecretBackupCount() int {
	p := os.Getenv("AEGIS_SAFE_SECRET_BACKUP_COUNT")
	if p == "" {
		return 3
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 3
	}
	return l
}

// SidecarMaxPollInterval returns the maximum interval for polling by the
// sidecar process. The value is read from the environment variable
// `AEGIS_SIDECAR_MAX_POLL_INTERVAL` or returns 300000 milliseconds as default.
func SidecarMaxPollInterval() time.Duration {
	p := os.Getenv("AEGIS_SIDECAR_MAX_POLL_INTERVAL")
	if p == "" {
		p = "300000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 300000 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}

// SidecarExponentialBackoffMultiplier returns the multiplier for exponential
// backoff by the sidecar process.
// The value is read from the environment variable
// `AEGIS_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER` or returns 2 as default.
func SidecarExponentialBackoffMultiplier() int64 {
	p := os.Getenv("AEGIS_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER")
	if p == "" {
		p = "2"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 2
	}
	return i
}

// SidecarSuccessThreshold returns the number of consecutive successful
// polls before reducing the interval. The value is read from the environment
// variable `AEGIS_SIDECAR_SUCCESS_THRESHOLD` or returns 3 as default.
func SidecarSuccessThreshold() int64 {
	p := os.Getenv("AEGIS_SIDECAR_SUCCESS_THRESHOLD")
	if p == "" {
		p = "3"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 3
	}
	return i
}

// SidecarErrorThreshold returns the number of consecutive failed polls
// before increasing the interval. The value is read from the environment
// variable `AEGIS_SIDECAR_ERROR_THRESHOLD` or returns 2 as default.
func SidecarErrorThreshold() int64 {
	p := os.Getenv("AEGIS_SIDECAR_ERROR_THRESHOLD")
	if p == "" {
		p = "2"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 2
	}
	return i
}
