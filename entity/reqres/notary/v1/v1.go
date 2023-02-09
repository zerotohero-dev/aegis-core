/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

type RegisterWorkloadRequest struct {
	WorkloadId string `json:"workloadId"`
	Err        string `json:"err,omitempty"`
}

type RegisterWorkloadResponse struct {
	Err string `json:"err,omitempty"`
}
