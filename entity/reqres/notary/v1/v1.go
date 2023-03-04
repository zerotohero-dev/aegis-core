/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package v1

// RegisterWorkloadRequest is experimental.
// It is not used anywhere, for anything at the moment.
type RegisterWorkloadRequest struct {
	WorkloadId string `json:"workloadId"`
	Err        string `json:"err,omitempty"`
}

// RegisterWorkloadResponse is experimental.
// It is not used anywhere, for anything at the moment.
type RegisterWorkloadResponse struct {
	Err string `json:"err,omitempty"`
}
