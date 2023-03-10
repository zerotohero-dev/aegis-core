/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package audit

import (
	reqres "github.com/zerotohero-dev/aegis-core/entity/reqres/safe/v1"
	"github.com/zerotohero-dev/aegis-core/log"
)

type JournalEntry struct {
	CorrelationId string
	Entity        any
	Method        string
	Url           string
	Svid          string
	Event         AuditEvent
}

type AuditEvent string

const AuditEventEnter AuditEvent = "aegis-enter"
const AuditEventBadSvid AuditEvent = "aegis-bad-svid"
const AuditEventBrokenBody AuditEvent = "aegis-broken-body"
const AuditEventRequestTypeMismatch AuditEvent = "aegis-request-type-mismatch"
const AuditEventBadPeerSvid AuditEvent = "aegis-bad-peer-svid"
const AuditEventNoSecret AuditEvent = "aegis-no-secret"
const AuditEventOk AuditEvent = "aegis-ok"
const AuditEventNoWorkloadId AuditEvent = "aegis-no-workload-id"

func printAudit(correlationId, method, url, svid, message string) {
	log.InfoLn(
		correlationId,
		"method", method,
		"url", url,
		"svid", svid,
		"msg", message,
	)
}

func Log(e JournalEntry) {
	if e.Entity == nil {
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid, string(e.Event),
		)
	}

	switch v := e.Entity.(type) {
	case reqres.SecretFetchRequest:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"m"+string(e.Event),
		)
	case reqres.SecretFetchResponse:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"c:"+v.Created+"u:"+v.Updated+"m:"+string(e.Event),
		)
	case reqres.SecretUpsertRequest:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"m"+string(e.Event),
		)
	case reqres.SecretUpsertResponse:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"m"+string(e.Event),
		)
	case reqres.SecretListRequest:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"m"+string(e.Event),
		)
	case reqres.SecretListResponse:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"m"+string(e.Event),
		)
	default:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e: UNKNOWN ENTITY in AUDIT LOG",
		)
	}
}
