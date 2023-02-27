/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.ist>
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

var File BackingStore = "file"
var Memory BackingStore = "memory"
var Cluster BackingStore = "cluster"

type SecretFormat string

var Json SecretFormat = "json"
var Yaml SecretFormat = "yaml"
var None SecretFormat = "none"

type SecretMeta struct {
	// Overrides Env.SafeUseKubernetesSecrets()
	UseKubernetesSecret bool `json:"k8s"`
	// Overrides Env.SafeBackingStoreType()
	BackingStore BackingStore `json:"storage"`
	// Defaults to "aegis-system"
	Namespace string `json:"namespace"`
	// Go template used to transform the secret.
	// Sample secret:
	// '{"username":"admin","password":"AegisRocks"}'
	// Sample template:
	// '{"USER":"{{.username}}", "PASS":"{{.password}}"}"
	Template string `json:"template"`
	// Defaults to None
	Format SecretFormat
	// Defaults to false. If `true`, the secret needs to be encrypted using
	// Aegis Safe Age Public Key beforehand.
	Encrypted bool
}

type SecretStored struct {
	Name    string
	Value   string
	Meta    SecretMeta
	Created time.Time
	Updated time.Time
}
