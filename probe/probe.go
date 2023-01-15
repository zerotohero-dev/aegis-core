/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package probe

import (
	"fmt"
	"github.com/zerotohero-dev/aegis-core/env"
	"log"
	"net/http"
)

func ok(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "OK")
	if err != nil {
		log.Printf("probe response failure: %s", err.Error())
		return
	}
}

func CreateLiveness() {
	http.HandleFunc("/healthz", ok)
}

func CreateReadiness() {
	http.HandleFunc("/readyz", ok)
}

func Listen() {
	err := http.ListenAndServe(env.ProbeServerPort(), nil)
	if err != nil {
		log.Fatalf("error creating readiness probe: %s", err.Error())
		return
	}
}
