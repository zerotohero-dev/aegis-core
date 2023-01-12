/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package validation

import (
	"github.com/zerotohero-dev/aegis-core/env"
	"strings"
)

// IsSentinel returns true if the given SVID (SPIFFE Verifiable Identity Document)
// is a Sentinel SVID.
// It does this by checking if the SVID has the SentinelSvidPrefix as its prefix.
func IsSentinel(svid string) bool {
	return strings.HasPrefix(svid, env.SentinelSvidPrefix())
}

// IsSafe returns true if the given SVID (SPIFFE Verifiable Identity Document)
// is a Safe SVID.
// It does this by checking if the SVID has the SafeSvidPrefix as its prefix.
func IsSafe(svid string) bool {
	return strings.HasPrefix(svid, env.SafeSvidPrefix())
}

// IsWorkload returns true if the given SVID (SPIFFE Verifiable Identity Document)
// is a Workload SVID.
// It does this by checking if the SVID has the WorkloadSvidPrefix as its prefix.
func IsWorkload(svid string) bool {
	return strings.HasPrefix(svid, env.WorkloadSvidPrefix())
}
