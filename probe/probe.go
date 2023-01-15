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
	mux := http.NewServeMux()
	mux.HandleFunc("/", ok)
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatalf("error creating liveness probe: %s", err.Error())
		return
	}
}

func CreateReadiness() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ok)
	err := http.ListenAndServe(":8082", mux)
	if err != nil {
		log.Fatalf("error creating readiness probe: %s", err.Error())
		return
	}
}
