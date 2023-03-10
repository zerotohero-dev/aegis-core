/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package v1

import (
	data "github.com/zerotohero-dev/aegis-core/entity/data/v1"
)

type SecretUpsertRequest struct {
	WorkloadId    string            `json:"key"`
	BackingStore  data.BackingStore `json:"backingStore"`
	UseKubernetes bool              `json:"useKubernetes"`
	Namespace     string            `json:"namespace"`
	Value         string            `json:"value"`
	Template      string            `json:"template"`
	Format        data.SecretFormat `json:"format"`
	Encrypt       bool              `json:"bool"`
	Err           string            `json:"err,omitempty"`
}

type SecretUpsertResponse struct {
	Err string `json:"err,omitempty"`
}

type SecretFetchRequest struct {
	Err string `json:"err,omitempty"`
}

type SecretFetchResponse struct {
	Data    string `json:"data"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Err     string `json:"err,omitempty"`
}

type SecretListRequest struct {
	Err string `json:"err,omitempty"`
}

type SecretListResponse struct {
	Secrets []data.Secret `json:"secrets"`
	Err     string        `json:"err,omitempty"`
}

type GenericRequest struct {
	Err string `json:"err,omitempty"`
}

type GenericResponse struct {
	Err string `json:"err,omitempty"`
}
