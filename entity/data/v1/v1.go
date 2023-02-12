/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RubyDate))
	return []byte(stamp), nil
}

type Secret struct {
	Name    string   `json:"name"`
	Created JsonTime `json:"created"`
	Updated JsonTime `json:"updated"`
}

type BackingStore string

var Disk BackingStore = "disk"
var Memory BackingStore = "memory"
var Cluster BackingStore = "cluster"

type SecretMeta struct {
	// Overrides Env.SafeUseKubernetesSecrets()
	UseKubernetesSecret bool `json:"k8s"`
	// Overrides Env.SafeBackingStoreType()
	BackingStore BackingStore `json:"storage"`
}

type SecretStored struct {
	Name    string
	Value   string
	Meta    SecretMeta
	Created time.Time
	Updated time.Time
}
