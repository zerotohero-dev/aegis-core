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
	http.HandleFunc("/", ok)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error creating liveness probe: %s", err.Error())
		return
	}
}

func CreateReadiness() {
	http.HandleFunc("/", ok)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("error creating readiness probe: %s", err.Error())
		return
	}
}
